package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/pstrobl96/prusa_exporter/prusalink"
	"github.com/pstrobl96/prusa_exporter/syslog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	configFile      = kingpin.Flag("config.file", "Configuration file for prusa_exporter.").Default("./prusa.yml").ExistingFile()
	configReload    = kingpin.Flag("config.reload", "Interval how often should be config reloaded - 0 for no reload.").Default("300").Int()
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

	config, err = probeConfigFile(config)

	if err != nil {
		log.Error().Msg("Error probing configuration file " + err.Error())
		os.Exit(1)
	}

	logLevel, err := zerolog.ParseLevel(config.Exporter.LogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel // default log level
	}

	zerolog.SetGlobalLevel(logLevel)

	if config.Exporter.ReloadInterval != 0 { // do not run reloader if interval is set to zero
		go configReloader(&config, *configReload) // run reloader as goroutine
	}

	if config.Exporter.Syslog.Enabled {
		log.Warn().Msg("Syslog metrics enabled!")
		log.Warn().Msg("Syslog metrics server starting at: " + config.Exporter.Syslog.ListenAddress)
		go syslog.HandleMetrics(config.Exporter.Syslog.ListenAddress)
	}

}

func probeConfigFile(config config.Config) (config.Config, error) {
	for _, printer := range config.Printers {
		status, err := prusalink.ProbePrinter(printer)
		if err != nil {
			log.Error().Msg(err.Error())
			printer.Reachable = false
		} else {
			printerType, err := prusalink.GetPrinterType(printer)
			if err != nil {
				log.Error().Msg(err.Error())
				printer.Type = "unknown"
			} else if status {
				printer.Type = printerType
			} else {
				printer.Type = "unknown"
			}
			printer.Reachable = status
		}
	}
	return config, nil
}

func configReloader(configuration *config.Config, reloadInterval int) {
	ticker := time.NewTicker(time.Duration(reloadInterval) * time.Second)

	for t := range ticker.C {
		log.Info().Msg(fmt.Sprintf("Config reloaded at: %v\n", t.UTC()))
		config, err := config.LoadConfig(*configFile)
		if err != nil {
			log.Error().Msg("Error loading configuration file " + err.Error())
		}
		config, err = probeConfigFile(config)
		if err != nil {
			log.Error().Msg("Error probing configuration file " + err.Error())
		}

	}
}
