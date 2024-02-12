package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
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
		return []byte{}, err
	}

	return jsonMessage, nil
}

func sendLokiMessage(message []byte, loki string, timestamp time.Time) (*http.Response, error) {

	r, err := http.NewRequest("POST", loki+"/loki/api/v1/push", bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	return res, nil
}
