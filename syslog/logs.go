package syslog

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

// HandleLogs is a function to handle logs from syslog and send them to Loki or Promtail - promtail does not work because printers send logs in a different format than it should and Promtails throws EOF error
func HandleLogs(listenUDP string) {
	channel, server := startSyslogServer(listenUDP)
	log.Debug().Msg("Syslog server for logs started at: " + listenUDP)
	syslogLogger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.InfoLevel).Output(zerolog.ConsoleWriter{Out: os.Stdout})
	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {

			currentTime := time.Now()
			syslogLogger.Info().
				Time("time", currentTime).
				Str("app_name", logParts["app_name"].(string)).
				Str("client", logParts["client"].(string)).
				Str("hostname", logParts["hostname"].(string)).
				Str("message", logParts["message"].(string))

		}
	}(channel)

	server.Wait()
}
