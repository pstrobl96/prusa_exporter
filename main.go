package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/exp/slog"
)

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stdout))
	logger.Info("Buddy Link Prometheus exporter starting")
	loadEnvVars()
	buddyCollector := newBuddyCollector()
	legacyCollector := newLegacyCollector()
	einsyCollector := newEinsyCollector()
	prometheus.MustRegister(buddyCollector, legacyCollector, einsyCollector)

	logger.Info("Metrics registered")
	http.Handle("/metrics", promhttp.Handler())
	logger.Error(http.ListenAndServe(":"+metricsPort, nil).Error())
}
