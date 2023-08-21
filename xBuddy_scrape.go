package main

import (
	"encoding/json"
)

func getBuddyVersion(address string, apiKey string, username string, password string) buddyVersion {
	resp := accessBuddyApi("version", address, apiKey, username, password)

	var result buddyVersion

	if e := json.Unmarshal(resp, &result); e != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func getBuddyFiles(address string, apiKey string, username string, password string) buddyFiles {
	resp := accessBuddyApi("files", address, apiKey, username, password)

	var result buddyFiles

	if e := json.Unmarshal(resp, &result); e != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func getBuddyJob(address string, apiKey string, username string, password string) buddyJob {
	resp := accessBuddyApi("job", address, apiKey, username, password)

	var result buddyJob

	if e := json.Unmarshal(resp, &result); e != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func getBuddyPrinter(address string, apiKey string, username string, password string) buddyPrinter {
	resp := accessBuddyApi("printer", address, apiKey, username, password)

	var result buddyPrinter

	if e := json.Unmarshal(resp, &result); e != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}
