package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Buddy Link Prometheus exporter starting")
	loadEnvVars()
	go configReloader()

	buddyCollector := newBuddyCollector()
	einsyCollector := newEinsyCollector()
	prometheus.MustRegister(buddyCollector, einsyCollector)
	log.Info().Msg("Metrics registered")
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal().Msg(http.ListenAndServe(":"+metricsPort, nil).Error())
}
