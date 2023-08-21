package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/slog"
	"gopkg.in/yaml.v3"
)

var configPath string
var metricsPort string
var scrapeTimeout float64
var loadedConfig config
var logger *slog.Logger

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

func loadEnvVars() {
	configPath = getCfgFile()
	metricsPort = getMetricsPort()
	scrapeTimeout = getScrapeTimeout()
	loadedConfig = probeCfgFile(parseCfg(configPath))
}

func getCfgFile() string {
	cfgFile := os.Getenv("BUDDY_EXPORTER_CONFIG")
	if cfgFile == "" {
		pwd, e := os.Getwd()
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		fmt.Println(pwd)
		cfgFile = pwd + "/buddy.yaml"
	}

	logger.Info("Using config - " + cfgFile)

	return cfgFile
}

func probeCfgFile(parsedConfig config) config {
	for _, s := range parsedConfig.Printers.Buddy {
		if !head(s.Address) {
			s.Reachable = false
			logger.Info("NOT REACHABLE - " + s.Address)
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

	logger.Info("Using port - " + metricsPort)

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

	logger.Info("Scraping interval - " + strconv.FormatFloat(result, 'g', 5, 64) + " sec")

	return result
}

func parseCfg(path string) config {
	f, e := os.ReadFile(path)
	if e != nil {
		logger.Error(e.Error())
	}
	var p config
	if e := yaml.Unmarshal(f, &p); e != nil {
		logger.Error(e.Error())
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
			fmt.Println("tick")
		}
	}
}