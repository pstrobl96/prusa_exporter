package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var loadedConfig config
var envVars envVariables

var logLevel string
var configPath string
var metricsPort string
var scrapeTimeout float64

type config struct {
	Printers struct {
		Buddy []struct {
			Address   string `yaml:"address"`
			Name      string `yaml:"name"`
			Type      string `yaml:"type"`
			Apikey    string `yaml:"apikey,omitempty"`
			Username  string `yaml:"username,omitempty"`
			Pass      string `yaml:"pass,omitempty"`
			Reachable bool 
		} `yaml:"buddy"`
		Einsy []struct {
			Address   string `yaml:"address"`
			Apikey    string `yaml:"apikey"`
			Name      string `yaml:"name"`
			Type      string `yaml:"type"`
			Reachable bool 
		} `yaml:"einsy"`
	} `yaml:"printers"`
}

type envVariables struct {
	configPath string
	metricsPort string
	scrapeTimeout float64
	reloadInteval int
}

func loadEnvVars() {
	logLevel := setLogLevel("")
	configPath = getCfgFile()
	metricsPort = getMetricsPort()
	scrapeTimeout = getScrapeTimeout()
	loadedConfig = probeCfgFile(parseCfg(configPath))

	log.Info().Msg("Log level - " + logLevel)
	log.Info().Msg("Using config - " + configPath)


}

func loadEnvVarsToStruct() envVariables {
	log.Info().Msg(os.Getenv("BUDDY_LOG_LEVEL"))
	var result envVariables
	result.configPath = getCfgFile()
	result.metricsPort = getMetricsPort()
	result.scrapeTimeout = getScrapeTimeout()

	
	configPath = getCfgFile()
	metricsPort = getMetricsPort()
	scrapeTimeout = getScrapeTimeout()
	loadedConfig = probeCfgFile(parseCfg(configPath))

	log.Info().Msg("Log level - " + logLevel)
	log.Info().Msg("Using config - " + configPath)

	return result

}

func setLogLevel(level string) string {
	if logLevel == "" {
		logLevel = "info"
	}

	logLevel = strings.ToLower(logLevel)
	var zeroLogLevel zerolog.Level

	switch logLevel {
	case "info":
		zeroLogLevel = zerolog.InfoLevel
	case "debug":
		zeroLogLevel = zerolog.DebugLevel
	case "trace":
		zeroLogLevel = zerolog.TraceLevel
	case "error":
		zeroLogLevel = zerolog.ErrorLevel		
	case "panic":
		zeroLogLevel = zerolog.PanicLevel
	case "fatal":
		zeroLogLevel = zerolog.FatalLevel
	default:
		zeroLogLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(zeroLogLevel)

	return logLevel
}

func getCfgFile() string {
	cfgFile := os.Getenv("BUDDY_EXPORTER_CONFIG")
	if cfgFile == "" {
		pwd, e := os.Getwd()
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		cfgFile = pwd + "/buddy.yaml"
	}

	return cfgFile
}

func probeCfgFile(parsedConfig config) config {
	for _, s := range parsedConfig.Printers.Buddy {
		if !head(s.Address) {
			s.Reachable = false
			log.Error().Msg(s.Address + " is not reachable")
		} else {
			s.Reachable = true
		}
	}
	return parsedConfig
}

func getMetricsPort() string {
	metricsPort := os.Getenv("BUDDY_EXPORTER_PORT")
	if metricsPort == "" {
		metricsPort = "10009"
	}

	log.Info().Msg("Using port - " + metricsPort)

	return metricsPort
}

func getScrapeTimeout() float64 {
	result := 1.0
	metricsPort := os.Getenv("BUDDY_EXPORTER_SCRAPE_TIMEOUT")
	if metricsPort != "" {
		parsed, e := strconv.ParseFloat(metricsPort, 64)
		if e != nil {
			result = 1.0
		} else {
			result = parsed
		}
	}

	log.Info().Msg("Scraping interval - " + strconv.FormatFloat(result, 'g', 5, 64) + " sec")

	return result
}

func parseCfg(path string) config {
	f, e := os.ReadFile(path)
	if e != nil {
		log.Error().Msg(e.Error())
	}
	var p config
	if e := yaml.Unmarshal(f, &p); e != nil {
		log.Error().Msg(e.Error())
	}
	return p
}

func testConnection(s string) (bool, error) {
	r, e := http.Head(s)
	return r.StatusCode == 200, e
}

func configReloader() {
	t := time.NewTicker(300 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C: // Activate periodically
			loadEnvVars()
			log.Debug().Msg("Config reloaded")
		}
	}
}
