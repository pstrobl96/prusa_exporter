package main

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/icholy/digest"
	"gopkg.in/yaml.v3"
)

type config struct {
	Printers struct {
		APIKey []struct {
			Address string `yaml:"address"`
			Apikey  string `yaml:"apikey"`
			Name    string `yaml:"name"`
			Type    string `yaml:"type"`
		} `yaml:"apiKey"`
		Password []struct {
			Address  string `yaml:"address"`
			Username string `yaml:"username"`
			Pass     string `yaml:"pass"`
			Name     string `yaml:"name"`
			Type     string `yaml:"type"`
		} `yaml:"password"`
	} `yaml:"printers"`
}

type scrapeItem struct {
	Address  string
	ApiKey   string
	Username string
	Password string
	Name     string
	Type     string
}

type scrape struct {
	Printers []scrapeItem
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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func generateHeader(apiKey string, username string, pass string) (string, string, error) {
	if apiKey == "" && (username == "" || pass == "") {
		return "", "", errors.New("no auth provided")
	}
	return "X-Api-Key", apiKey, nil
}

func accessApi(path string, address string, apiKey string, username string, password string) *http.Response {
	url := string("http://" + address + "/api/" + path)
	var res *http.Response
	var err error
	if apiKey == "" {
		client := &http.Client{
			Transport: &digest.Transport{
				Username: username,
				Password: password,
			},
		}
		res, err = client.Get(url)
		if err != nil {
			panic(err)
		}
	} else {
		req, _ := http.NewRequest("GET", url, nil)
		client := &http.Client{}
		req.Header.Add("X-Api-Key", apiKey)
		res, err = client.Do(req)
		if err != nil {
			panic(err)
		}
	}
	return res
}
