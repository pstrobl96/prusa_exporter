package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

var configPath string
var metricsPort string
var scrapeTimeout int
var loadedConfig config

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

	log.Println("Using config - " + cfgFile)

	return cfgFile
}

func getMetricsPort() string {
	metricsPort := os.Getenv("BUDDY_EXPORTER_PORT")
	if metricsPort == "" {
		metricsPort = "10009"
	}

	log.Println("Using port - " + metricsPort)

	return metricsPort
}

func getScrapeTimeout() int {
	var result int
	metricsPort := os.Getenv("BUDDY_EXPORTER_SCRAPE_TIMEOUT")
	if metricsPort == "" {
		result = 1
	} else {
		parsed, err := strconv.Atoi(metricsPort)
		if err != nil {
			result = 1
		} else {
			result = parsed
		}
	}

	log.Println("Scraping interval - " + strconv.Itoa(result) + " sec")

	return result
}

func parseCfg(path string) config {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var p config
	if err := yaml.Unmarshal(f, &p); err != nil {
		log.Fatal(err)
	}
	return p
}
