package main

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func getEinsyJob(address string, apiKey string) einsyJob {
	resp, _ := accessEinsyApi("job", address, apiKey)

	var result einsyJob

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")
	}

	return result
}

func getEinsyCameras(address string, apiKey string) einsyCameras {
	resp, _ := accessEinsyApi("v1/cameras", address, apiKey)

	var result einsyCameras

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}

func getEinsyPrinter(address string, apiKey string) einsyPrinter {
	resp, _ := accessEinsyApi("printer", address, apiKey)

	var result einsyPrinter

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}

//func getEinsyStorage(address string, apiKey string) einsyStorage { // currently unused
//	resp, _ := accessEinsyApi("v1/storage", address, apiKey)
//
//	var result einsyStorage
//
//	if e := json.Unmarshal(resp, &result); e != nil {
//		log.Error().Msg("Can not unmarshal JSON")
//
//	}
//
//	return result
//}

func getEinsySettings(address string, apiKey string) einsySettings {
	resp, _ := accessEinsyApi("settings", address, apiKey)

	var result einsySettings

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}

// func getEinsyConnection(address string, apiKey string) einsyConection { // currently unused
// 	resp, _ := accessEinsyApi("connection", address, apiKey)
// 
// 	var result einsyConection
// 
// 	if e := json.Unmarshal(resp, &result); e != nil {
// 		log.Error().Msg("Can not unmarshal JSON")
// 
// 	}
// 
// 	return result
// }

func getEinsyFiles(address string, apiKey string) einsyFiles {
	resp, _ := accessEinsyApi("files", address, apiKey)

	var result einsyFiles

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}

func getEinsyLogs(address string, apiKey string) einsyLogs {
	resp, _ := accessEinsyApi("logs", address, apiKey)

	var result einsyLogs

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}

func getEinsyInfo(address string, apiKey string) einsyInfo {
	resp, _ := accessEinsyApi("v1/info", address, apiKey)

	var result einsyInfo

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}

// func getEinsyStatus(address string, apiKey string) einsyStatus { // currently unused
// 	resp, _ := accessEinsyApi("v1/status", address, apiKey)
// 
// 	var result einsyStatus
// 
// 	if e := json.Unmarshal(resp, &result); e != nil {
// 		log.Error().Msg("Can not unmarshal JSON")
// 
// 	}
// 
// 	return result
// }

func getEinsyVersion(address string, apiKey string) einsyVersion {
	resp, _ := accessEinsyApi("version", address, apiKey)

	var result einsyVersion

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}

func getEinsyPorts(address string, apiKey string) einsyPorts {
	resp, _ := accessEinsyApi("ports", address, apiKey)

	var result einsyPorts

	if e := json.Unmarshal(resp, &result); e != nil {
		log.Error().Msg("Can not unmarshal JSON")

	}

	return result
}
