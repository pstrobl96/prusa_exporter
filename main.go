package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("Buddy Prusa Link Prometheus exporter starting")

	collector := newBuddyCollector()
	prometheus.MustRegister(collector)

	log.Println("Metrics registered")

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":10009", nil))
}
