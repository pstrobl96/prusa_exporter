package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func getEinsyJob(address string, apiKey string) einsyJob {
	resp := accessEinsyApi("job", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyJob

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyCameras(address string, apiKey string) einsyCameras {
	resp := accessEinsyApi("v1/cameras", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyCameras

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyPrinter(address string, apiKey string) einsyPrinter {
	resp := accessEinsyApi("printer", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyPrinter

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyStorage(address string, apiKey string) einsyStorage {
	resp := accessEinsyApi("v1/storage", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyStorage

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsySettings(address string, apiKey string) einsySettings {
	resp := accessEinsyApi("settings", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsySettings

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyConnection(address string, apiKey string) einsyConection {
	resp := accessEinsyApi("connection", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyConection

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyFiles(address string, apiKey string) einsyFiles {
	resp := accessEinsyApi("files", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyFiles

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyLogs(address string, apiKey string) einsyLogs {
	resp := accessEinsyApi("logs", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyLogs

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyInfo(address string, apiKey string) einsyInfo {
	resp := accessEinsyApi("v1/info", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyInfo

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyStatus(address string, apiKey string) einsyStatus {
	resp := accessEinsyApi("v1/status", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyStatus

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyVersion(address string, apiKey string) einsyVersion {
	resp := accessEinsyApi("version	", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyVersion

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getEinsyPorts(address string, apiKey string) einsyPorts {
	resp := accessEinsyApi("ports", address, apiKey)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result einsyPorts

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}
