package cmd

import (
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	configFile      = kingpin.Flag("config.file", "Configuration file for prusa_exporter.").Default("./prusa.yml").ExistingFile()
	metricsPath     = kingpin.Flag("exporter.metrics-path", "Path where to expose metrics.").Default("/metrics").String()
	exporterMetrics = kingpin.Flag("exporter.metrics", "Decides if expose metrics about exporter itself.").Default("true").Bool()
)

// Run function to start the exporter
func Run() {
	kingpin.Parse()
	log.Info().Msg("Prusa exporter starting")
	log.Info().Msg("Loading configuration file: " + *configFile)

	config, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Error().Msg("Error loading configuration file " + err.Error())
		os.Exit(1)
	}

	logLevel, err := zerolog.ParseLevel(config.Exporter.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel // default log level
	}

	zerolog.SetGlobalLevel(logLevel)
}
