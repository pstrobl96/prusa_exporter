package main

import (
	"encoding/json"
	"log"
)

func getBuddyVersion(address string, apiKey string, username string, password string) buddyVersion {
	resp := accessBuddyApi("version", address, apiKey, username, password)

	var result buddyVersion

	if err := json.Unmarshal(resp, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getBuddyFiles(address string, apiKey string, username string, password string) buddyFiles {
	resp := accessBuddyApi("files", address, apiKey, username, password)

	var result buddyFiles

	if err := json.Unmarshal(resp, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getBuddyJob(address string, apiKey string, username string, password string) buddyJob {
	resp := accessBuddyApi("job", address, apiKey, username, password)

	var result buddyJob

	if err := json.Unmarshal(resp, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getBuddyPrinter(address string, apiKey string, username string, password string) buddyPrinter {
	resp := accessBuddyApi("printer", address, apiKey, username, password)

	var result buddyPrinter

	if err := json.Unmarshal(resp, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}
