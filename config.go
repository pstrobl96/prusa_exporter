package main

import (
	"fmt"
	"os"
	"strconv"

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
			Address  string `yaml:"address"`
			Name     string `yaml:"name"`
			Type     string `yaml:"type"`
			Apikey   string `yaml:"apikey,omitempty"`
			Username string `yaml:"username,omitempty"`
			Pass     string `yaml:"pass,omitempty"`
		} `yaml:"buddy"`
		Einsy []struct {
			Address string `yaml:"address"`
			Apikey  string `yaml:"apikey"`
			Name    string `yaml:"name"`
			Type    string `yaml:"type"`
		} `yaml:"einsy"`
		Legacy []struct {
			Address string `yaml:"address"`
			Name    string `yaml:"name"`
			Type    string `yaml:"type"`
		} `yaml:"legacy"`
	} `yaml:"printers"`
}

func loadEnvVars() {
	configPath = getCfgFile()
	metricsPort = getMetricsPort()
	scrapeTimeout = getScrapeTimeout()
	loadedConfig = parseCfg(configPath)
}

func getCfgFile() string {
	cfgFile := os.Getenv("BUDDY_EXPORTER_CONFIG")
	if cfgFile == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pwd)
		cfgFile = pwd + "/buddy.yaml"
	}

	logger.Info("Using config - " + cfgFile)

	return cfgFile
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
		parsed, err := strconv.ParseFloat(metricsPort, 64)
		if err != nil {
			result = 1.0
		} else {
			result = parsed
		}
	}

	logger.Info("Scraping interval - " + strconv.FormatFloat(result, 'E', -1, 32) + " sec")

	return result
}

func parseCfg(path string) config {
	f, err := os.ReadFile(path)
	if err != nil {
		logger.Error(err.Error())
	}
	var p config
	if err := yaml.Unmarshal(f, &p); err != nil {
		logger.Error(err.Error())
	}
	return p
}
