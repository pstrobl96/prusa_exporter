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

/*
	func createScrape(c config) scrape {
		var result scrape
		for _, v := range c.Printers.APIKey {
			header, headerVal, err := generateHeader(v.Apikey, "", "")
			if err != nil {
				panic(err)
			}
			result.Printers = append(result.Printers, scrapeItem{Address: v.Address, Header: header, HeaderVal: headerVal, Name: v.Name, Type: v.Type})
		}
		for _, v := range c.Printers.Password {
			header, headerVal, err := generateHeader("", v.Username, v.Pass)
			if err != nil {
				panic(err)
			}
			result.Printers = append(result.Printers, scrapeItem{Address: v.Address, Header: header, HeaderVal: headerVal, Name: v.Name, Type: v.Type})
		}
		return result
	}
*/
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

/*func accessApi(path string, address string, header string, headerVal string) *http.Response {
	url := string("http://" + address + "/api/" + path)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add(header, headerVal)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}*/

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

/*
 */
/*func digestApi(path string, address string, username string, pass string) *http.Response {
	// create a new digest authentication request
	dr := dac.NewRequest(username, pass, "GET", address, "")
	response1, _ := dr.Execute()

	fmt.Println(response1.Body)

	// check error, get response
}
*/
