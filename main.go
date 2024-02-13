package main

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func initProcedure() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	loadConfigFile()
}

func main() {
	log.Info().Msg("Prusa exporter starting")
	initProcedure()                          // initialize
	if config.Exporter.ReloadInterval != 0 { // do not run reloader if interval is set to zero
		go configReloader() // run reloader as goroutine
	}

	if config.Exporter.Syslog.Metrics.Enabled {
		log.Warn().Msg("Syslog metrics enabled!")
		log.Warn().Msg("Syslog metrics server starting at: " + config.Exporter.Syslog.Metrics.ListenUDP)
		go handleMetrics(config.Exporter.Syslog.Metrics.ListenUDP)
	}

	if config.Exporter.Syslog.Logs.Enabled {
		log.Warn().Msg("Syslog logs enabled!")
		log.Warn().Msg("Syslog logs server starting at: " + config.Exporter.Syslog.Logs.ListenUDP)
		go handleLogs(config.Exporter.Syslog.Logs.ListenUDP, config.Exporter.Syslog.Logs.LokiEndpoint)
	}

	log.Info().Msg("Initialized")
	buddyCollector := newBuddyCollector()
	einsyCollector := newEinsyCollector()

	if config.Exporter.Syslog.Metrics.Enabled {
		syslogCollector := newSyslogCollector()
		prometheus.MustRegister(buddyCollector, einsyCollector, syslogCollector)
	} else {
		prometheus.MustRegister(buddyCollector, einsyCollector)
	}

	log.Info().Msg("Metrics registered")
	http.Handle("/metrics", promhttp.Handler())
	log.Info().Msg("Listening at port: " + strconv.Itoa(config.Exporter.MetricsPort))
	log.Fatal().Msg(http.ListenAndServe(":"+strconv.Itoa(config.Exporter.MetricsPort), nil).Error())
}
