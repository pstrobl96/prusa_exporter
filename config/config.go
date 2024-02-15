package config

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var config configuration
var configPath string

type buddy struct {
	Address   string `yaml:"address"`
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Username  string `yaml:"username,omitempty"`
	Pass      string `yaml:"pass,omitempty"`
	Apikey    string `yaml:"apikey,omitempty"`
	Reachable bool
}

type sl struct {
	Address   string `yaml:"address"`
	Username  string `yaml:"username"`
	Pass      string `yaml:"pass"`
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Reachable bool
}

type einsy struct {
	Address   string `yaml:"address"`
	Apikey    string `yaml:"apikey"`
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Reachable bool
}

type configuration struct {
	Printers struct {
		Buddy []buddy `yaml:"buddy"`
		Einsy []einsy `yaml:"einsy"`
		Sl    []sl    `yaml:"sl"`
	} `yaml:"printers"`
	Exporter struct {
		MetricsPort    int    `yaml:"metrics_port"`
		ScrapeTimeout  int    `yaml:"scrape_timeout"`
		ReloadInterval int    `yaml:"reload_interval"`
		LogLevel       string `yaml:"log_level"`
		Syslog         struct {
			Metrics struct {
				Enabled   bool   `yaml:"enabled"`
				ListenUDP string `yaml:"listen_udp"`
			} `yaml:"metrics"`
			Logs struct {
				Enabled      bool   `yaml:"enabled"`
				ListenUDP    string `yaml:"listen_udp"`
				LokiEndpoint string `yaml:"loki_endpoint"`
			} `yaml:"logs"`
		} `yaml:"syslog"`
	} `yaml:"exporter"`
}

func setLogLevel(level string) string {
	if level == "" {
		level = "info"
	}

	level = strings.ToLower(level)
	var zeroLogLevel zerolog.Level

	switch level {
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
	return level
}

func loadConfigFile() {
	parsedConfig := parseConfig(getConfigPath())
	setLogLevel(config.Exporter.LogLevel)
	config = probeConfigFile(parsedConfig)
}

func getConfigPath() string {
	if configPath == "" {
		var cfgFile string
		flag.StringVar(&cfgFile, "config.file", "", "Path to prusa.yml config file") // later will use flag.Args
		flag.Parse()
		if cfgFile == "" {
			pwd, e := os.Getwd()
			if e != nil {
				log.Error().Msg(e.Error())
				os.Exit(1)
			}
			cfgFile = pwd + "/prusa.yml"
		}
		configPath = cfgFile
		return cfgFile
	}
	return configPath
}

func parseConfig(path string) configuration {
	f, e := os.ReadFile(path)
	if e != nil {
		log.Panic().Msg(e.Error())
	}
	var p configuration
	if e := yaml.Unmarshal(f, &p); e != nil {
		log.Error().Msg(e.Error())
	}
	return p
}

func probeConfigFilmain.go
	}

	for i, s := range parsedConfig.Printers.Sl {
		conn, status := testConnection(s.Address)
		if conn && status == 200 {
			parsedConfig.Printers.Sl[i].Reachable = true
			_, _, _, _, version, err := getSLResponse(s)
			if err == nil {
				if version.Hostname == "" {
					parsedConfig.Printers.Sl[i].Type = "unknown"
				} else {
					parsedConfig.Printers.Sl[i].Type = version.Hostname
				}
			}
		} else {
			parsedConfig.Printers.Sl[i].Reachable = false
			log.Error().Msg(s.Address + " is not reachable")
		}
	}

	for i, s := range parsedConfig.Printers.Einsy {
		_, status := testConnection(s.Address)

		if status == 401 { // yup it's weird, but it's how it works
			version, _, _, _, _, _, _, _, err := getEinsyResponse(s)
			if err == nil && version.Text != "" {

				parsedConfig.Printers.Einsy[i].Type = version.Original
				parsedConfig.Printers.Einsy[i].Reachable = true
			} else {
				parsedConfig.Printers.Einsy[i].Reachable = false
				log.Error().Msg(s.Address + " is not reachable") // i know, i repeated code will resolve later
			}
		} else {
			parsedConfig.Printers.Einsy[i].Reachable = false
			log.Error().Msg(s.Address + " is not reachable")
		}
	}
	return parsedConfig
}

func testConnection(s string) (bool, int) {
	req, _ := http.NewRequest("GET", "http://"+s+"/", nil)
	client := &http.Client{Timeout: time.Duration(config.Exporter.ScrapeTimeout) * time.Second}
	r, e := client.Do(req)
	return e == nil && r.StatusCode == 200, r.StatusCode
}

func configReloader() {
	ticker := time.NewTicker(time.Duration(config.Exporter.ReloadInterval) * time.Second)

	for t := range ticker.C {
		log.Info().Msg(fmt.Sprintf("Config reloaded at: %v\n", t.UTC()))
		loadConfigFile()
	}
}
