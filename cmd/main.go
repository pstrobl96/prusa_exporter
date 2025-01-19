package cmd

import (
	"net/http"
	"os"
	"strconv"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pstrobl96/prusa_exporter/config"
	prusalink "github.com/pstrobl96/prusa_exporter/prusalink/buddy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	configFile             = kingpin.Flag("config.file", "Configuration file for prusa_exporter.").Default("./prusa.yml").ExistingFile()
	metricsPath            = kingpin.Flag("exporter.metrics-path", "Path where to expose metrics.").Default("/metrics").String()
	metricsPort            = kingpin.Flag("exporter.metrics-port", "Port where to expose metrics.").Default("10009").Int()
	prusaLinkScrapeTimeout = kingpin.Flag("prusalink.scrape-timeout", "Timeout in seconds to scrape prusalink metrics.").Default("10").Int()
	logLevel               = kingpin.Flag("log.level", "Log level for zerolog.").Default("info").String()
)

// Run function to start the exporter
func Run() {
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
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano

	var collectors []prometheus.Collector

	config, err = probeConfigFile(config)

	if err != nil {
		log.Error().Msg("Error probing configuration file " + err.Error())
		os.Exit(1)
	}
	log.Info().Msg("PrusaLink metrics enabled!")
	collectors = append(collectors, prusalink.NewCollector(config))

	prometheus.MustRegister(collectors...)
	log.Info().Msg("Metrics registered")
	http.Handle(*metricsPath, promhttp.Handler())
	log.Info().Msg("Listening at port: " + strconv.Itoa(*metricsPort))
	log.Fatal().Msg(http.ListenAndServe(":"+strconv.Itoa(*metricsPort), nil).Error())

}

func probeConfigFile(config config.Config) (config.Config, error) {
	for i, printer := range config.Printers {
		if printer.Type == "" {
			status, err := prusalink.ProbePrinter(printer)
			if err != nil {
				log.Error().Msg(err.Error())
			} else if status {

				printerType, err := prusalink.GetPrinterType(printer)

				if err != nil {
					log.Error().Msg(err.Error())
				}

				config.Printers[i].Type = printerType
			}
		}
	}
	return config, nil
}
