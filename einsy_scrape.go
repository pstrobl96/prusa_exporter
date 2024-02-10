package main

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func getEinsyResponse(config einsy) (einsyVersion, einsyFiles, einsyJob, einsyPrinter, einsyCameras, einsyInfo, einsySettings, error) {

	log.Debug().Msg("Getting response from " + config.Address)

	version, e := accessEinsyAPI("version", config.Address, config.Apikey)
	var resultVersion einsyVersion
	if e = json.Unmarshal(version, &resultVersion); e != nil {
		log.Error().Msg("Can not unmarshal version JSON")
	}

	files, e := accessEinsyAPI("files", config.Address, config.Apikey)
	var resultFiles einsyFiles
	if e = json.Unmarshal(files, &resultFiles); e != nil {
		log.Error().Msg("Can not unmarshal files JSON")
	}

	job, e := accessEinsyAPI("job", config.Address, config.Apikey)
	var resultJob einsyJob
	if e = json.Unmarshal(job, &resultJob); e != nil {
		log.Error().Msg("Can not unmarshal job JSON")
	}

	printer, e := accessEinsyAPI("printer", config.Address, config.Apikey)
	var resultPrinter einsyPrinter
	if e = json.Unmarshal(printer, &resultPrinter); e != nil {
		log.Error().Msg("Can not unmarshal printer JSON")
	}

	cameras, e := accessEinsyAPI("v1/cameras", config.Address, config.Apikey)
	var resultCameras einsyCameras
	if e = json.Unmarshal(cameras, &resultCameras); e != nil {
		log.Error().Msg("Can not unmarshal printer JSON")
	}

	info, e := accessEinsyAPI("v1/info", config.Address, config.Apikey)
	var resultInfo einsyInfo
	if e = json.Unmarshal(info, &resultInfo); e != nil {
		log.Error().Msg("Can not unmarshal printer JSON")
	}

	settings, e := accessEinsyAPI("settings", config.Address, config.Apikey)
	var resultSettings einsySettings
	if e = json.Unmarshal(settings, &resultSettings); e != nil {
		log.Error().Msg("Can not unmarshal printer JSON")
	}

	return resultVersion, resultFiles, resultJob, resultPrinter, resultCameras, resultInfo, resultSettings, e
}
