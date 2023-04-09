package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var configPath string
var metricsPort string
var scrapeTimeout int

func loadEnvVars() {
	configPath = getCfgFile()
	metricsPort = getMetricsPort()
	scrapeTimeout = getScrapeTimeout()
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
