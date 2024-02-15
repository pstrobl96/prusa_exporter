package prusalink

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/icholy/digest"
	"github.com/rs/zerolog/log"
)

func getURL(path string, address string) string {
	return string("http://" + address + "/api/" + path)
}

func accessBuddyAPI(path string, address string, apiKey string, username string, password string) []byte {
	url := getURL(path, address)
	var res *http.Response
	var err error
	var body []byte
	if apiKey == "" {
		client := &http.Client{
			Transport: &digest.Transport{
				Username: username,
				Password: password,
			},
		}
		res, err = client.Get(url)

		if err != nil {
			log.Error().Msg(err.Error())
		}
	} else {
		req, _ := http.NewRequest("GET", url, nil)
		client := &http.Client{}
		req.Header.Add("X-Api-Key", apiKey)
		res, err = client.Do(req)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}
	if err == nil {
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	} else {
		log.Error().Msg(err.Error())
	}
	return body
}

func accessEinsyAPI(path string, address string, apiKey string) ([]byte, error) {
	url := getURL(path, address)
	var res *http.Response
	var err error
	var body []byte

	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	req.Header.Add("X-Api-Key", apiKey)
	res, err = client.Do(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return []byte{}, err
	} else if res.StatusCode != 200 {
		message := fmt.Sprintf("Return status code is: %d", res.StatusCode)
		log.Error().Msg(message)
		return []byte{}, errors.New(message)
	} else if err == nil && res.StatusCode == 200 {
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Error().Msg(err.Error())
			return []byte{}, err
		}
	}
	return body, err
}

func accessSLAPI(path string, address string, username string, password string) []byte {
	url := getURL(path, address)
	var res *http.Response
	var err error
	var body []byte
	client := &http.Client{
		Transport: &digest.Transport{
			Username: username,
			Password: password,
		},
	}
	res, err = client.Get(url)

	if err == nil {
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	} else {
		log.Error().Msg(err.Error())
	}
	return body
}
