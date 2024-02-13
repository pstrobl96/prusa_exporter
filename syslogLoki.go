package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

// LokiLabels is a struct for the stream that will be sent to Loki
type LokiLabels struct {
	App string `json:"app"`
	IP  string `json:"ip"`
	Mac string `json:"mac"`
}

// LogData is a struct for the data that will be sent to Loki
type LogData struct {
	Stream LokiLabels `json:"stream"`
	Values [][]string `json:"values"`
}

// LokiMessage is a struct for the message that will be sent to Loki
type LokiMessage struct {
	Streams []LogData `json:"streams"` // Embed LogData fields
}

func startSyslogLoggingService(listenUDP string, loki string) { // yep i just copied it from startSyslog, I wanted to use Promtail but it returned EOF and I was not able to get it up and running. Maybe up. However this part could be used also later for log analysis, there are data that looks interesting tho.
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)
	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenUDP(listenUDP)
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {

			currentTime := time.Now()
			message, err := composeLokiMessage(logParts["app_name"].(string), strings.Split(logParts["client"].(string), ":")[0], logParts["hostname"].(string), logParts["message"].(string), currentTime)
			if err != nil {
				log.Error().Msg(err.Error())
				continue
			}

			_, err = sendLokiMessage(message, loki, currentTime)
			if err != nil {
				log.Error().Msg(err.Error())
				continue
			}
		}
	}(channel)

	server.Wait()
}

func composeLokiMessage(app string, ip string, mac string, message string, timestamp time.Time) ([]byte, error) {
	lokiLabels := LokiLabels{
		App: app,
		IP:  ip,
		Mac: mac,
	}

	lokiData := LogData{
		Stream: lokiLabels,
		Values: [][]string{{strconv.FormatInt(timestamp.UnixNano(), 10), message}},
	}

	lokiMessage := LokiMessage{
		Streams: []LogData{lokiData},
	}

	jsonMessage, err := json.Marshal(lokiMessage)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error marshalling JSON: %s", err))
		return nil, err
	}

	return jsonMessage, nil
}

func sendLokiMessage(message []byte, loki string, timestamp time.Time) (*http.Response, error) {

	r, err := http.NewRequest("POST", loki+"/loki/api/v1/push", bytes.NewBuffer(message))
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("Error creating Loki request: %s", err))
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Warn().Msg(fmt.Sprintf("Error sending Loki message: %s - skipping", err))
		return nil, err
	}

	defer res.Body.Close()

	return res, nil
}
