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
	initProcedure()                          // initialize
	if config.Exporter.ReloadInterval != 0 { // do not run reloader if interval is set to zero
		go configReloader() // run reloader as goroutine
	}
	if config.Exporter.SyslogMetrics {
		log.Warn().Msg("Syslog metrics enabled!")
		go startSyslog(10008)
	}

	log.Info().Msg("Initialized")
	buddyCollector := newBuddyCollector()
	//syslogCollector := newSyslogCollector()
	einsyCollector := newEinsyCollector()
	prometheus.MustRegister(buddyCollector, einsyCollector)
	log.Info().Msg("Metrics registered")
	http.Handle("/metrics", promhttp.Handler())
	log.Info().Msg("Listening at port: " + strconv.Itoa(config.Exporter.MetricsPort))
	log.Fatal().Msg(http.ListenAndServe(":"+strconv.Itoa(config.Exporter.MetricsPort), nil).Error())
}
