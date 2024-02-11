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
	setLogLevel(config.Exporter.LogLevel)
}

func main() {
	log.Info().Msg("Prusa exporter starting")
	initProcedure() // initialize

	if config.Exporter.ReloadInterval != 0 { // do not run reloader if interval is set to zero
		go configReloader() // run reloader as goroutine
	}
	var syslogCollector *syslogCollector
	if config.Exporter.SyslogMetrics {
		log.Warn().Msg("Syslog metrics enabled!")
		log.Warn().Msg("Syslog server starting at port: " + strconv.Itoa(config.Exporter.SyslogPort))
		syslogCollector = newSyslogCollector()
		go startSyslog(config.Exporter.SyslogPort)
	}

	log.Info().Msg("Initialized")
	buddyCollector := newBuddyCollector()
	einsyCollector := newEinsyCollector()
	prometheus.MustRegister(buddyCollector, einsyCollector, syslogCollector)
	log.Info().Msg("Metrics registered")
	http.Handle("/metrics", promhttp.Handler())
	log.Info().Msg("Listening at port: " + strconv.Itoa(config.Exporter.MetricsPort))
	log.Fatal().Msg(http.ListenAndServe(":"+strconv.Itoa(config.Exporter.MetricsPort), nil).Error())
}
