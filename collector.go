package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

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

func loadCfg(path string) config {
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
