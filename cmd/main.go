package cmd

import (
	"net/http"
	"os"
	"strconv"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/pstrobl96/prusa_exporter/lineprotocol"
	prusalink "github.com/pstrobl96/prusa_exporter/prusalink/buddy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	configFile                = kingpin.Flag("config.file", "Configuration file for prusa_exporter.").Default("./prusa.yml").ExistingFile()
	metricsPath               = kingpin.Flag("exporter.metrics-path", "Path where to expose metrics.").Default("/metrics").String()
	metricsPort               = kingpin.Flag("exporter.metrics-port", "Port where to expose metrics.").Default("10009").Int()
	prusaLinkScrapeTimeout    = kingpin.Flag("prusalink.scrape-timeout", "Timeout in seconds to scrape prusalink metrics.").Default("10").Int()
	logLevel                  = kingpin.Flag("log.level", "Log level for zerolog.").Default("info").String()
	syslogListenAddress       = kingpin.Flag("listen-address", "Address where to expose port for gathering metrics. - format <address>:<port>").Default("0.0.0.0:8514").String()
	lineprotocolPrefix        = kingpin.Flag("prefix", "Prefix for lineprotocol metrics").Default("prusa_").String()
	lineprotocolExportAddress = kingpin.Flag("lineprotocol.export", "Export lineprotocol metrics to InfluxDB Proxy.").String()
)

// Run function to start the exporter
func Run() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano

	kingpin.Parse()
	log.Info().Msg("Prusa exporter starting")
	log.Info().Msg("Loading configuration file: " + *configFile)

	config, err := config.LoadConfig(*configFile, *prusaLinkScrapeTimeout)
	if err != nil {
		log.Error().Msg("Error loading configuration file " + err.Error())
		os.Exit(1)
	}

	logLevel, err := zerolog.ParseLevel(*logLevel)

	if err != nil {
		logLevel = zerolog.InfoLevel // default log level
	}
	zerolog.SetGlobalLevel(logLevel)

	var collectors []prometheus.Collector

	log.Info().Msg("PrusaLink metrics enabled!")
	collectors = append(collectors, prusalink.NewCollector(config))

	prometheus.MustRegister(collectors...)
	log.Info().Msg("Metrics registered")
	http.Handle(*metricsPath, promhttp.Handler())
	log.Info().Msg("Listening at port: " + strconv.Itoa(*metricsPort))

	log.Info().Msg("Syslog logs server starting at: " + *syslogListenAddress)
	lineprotocol.InitInfluxClient(*lineprotocolExportAddress)
	go lineprotocol.MetricsListener(*syslogListenAddress, *lineprotocolPrefix)
	log.Info().Msg("Syslog server started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
    <head><title>prusa_exporter</title></head>
    <body>
    <h1>prusa_exporter</h1>
	<p>Syslog server running at - <b>` + *syslogListenAddress + `</b></p>
    <p><a href="` + *metricsPath + `">Metrics</a></p>
	</body>
    </html>`))
	})

	log.Fatal().Msg(http.ListenAndServe(":"+strconv.Itoa(*metricsPort), nil).Error())

}
