package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("Buddy Prusa Link Prometheus exporter starting")
	loadEnvVars()
	//snap, err := getSnap("192.168.20.146")
	buddyCollector := newBuddyCollector()
	legacyCollector := newLegacyCollector()
	einsyCollector := newEinsyCollector()
	prometheus.MustRegister(buddyCollector, legacyCollector, einsyCollector)

	log.Println("Metrics registered")

	http.Handle("/metrics", promhttp.Handler())
	//http.Handle("/snap")
	http.HandleFunc("/snap", getSnap("192.168.20.162"))
	log.Fatal(http.ListenAndServe(":"+metricsPort, nil))
}
