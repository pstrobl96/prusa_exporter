package syslog

import (
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
	"gopkg.in/natefinch/lumberjack.v2"
)

// HandleLogs is a function to handle logs from syslog and send them to Loki or Promtail - promtail does not work because printers send logs in a different format than it should and Promtails throws EOF error
func HandleLogs(listenUDP string, directory string, filename string, maxSize int, maxBackups int, maxAge int) {
	channel, server := startSyslogServer(listenUDP)
	log.Debug().Msg("Syslog server for logs started at: " + listenUDP)

	if err := os.MkdirAll(directory, 0744); err != nil {
		log.Error().Err(err).Str("path", directory).Msg("Can't create log directory")
		return
	}

	writers := []io.Writer{&lumberjack.Logger{
		Filename:   path.Join(directory, filename),
		MaxBackups: maxBackups, // maximum number of backups
		MaxSize:    maxSize,    // in MB
		MaxAge:     maxAge,     // in Days
	}}

	mw := io.MultiWriter(writers...)

	syslogLogger := zerolog.New(mw).With().Timestamp().Logger()

	log.Debug().Msg("Syslog logs are being written to: " + path.Join(directory, filename))

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {

			currentTime := time.Now()

			syslogLogger.Info().
				Time("time", currentTime).
				Str("app_name", logParts["app_name"].(string)).
				Str("client", strings.Split(logParts["client"].(string), ":")[0]).
				Str("hostname", logParts["hostname"].(string)).
				Msg(logParts["message"].(string))

		}
	}(channel)

	server.Wait()
}
