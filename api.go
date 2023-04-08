package main

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/icholy/digest"
)

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

func accessBuddyApi(path string, address string, apiKey string, username string, password string) []byte {
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	return body
}

func accessLegacyApi(path string, address string) []byte {
	url := string("http://" + address + "/api/" + path)
	var res *http.Response
	var err error
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func accessEinsyApi(path string, address string, apiKey string) []byte {
	url := string("http://" + address + "/api/" + path)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := &http.Client{}
	req.Header.Set("X-Api-Key", apiKey)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	return body
}
