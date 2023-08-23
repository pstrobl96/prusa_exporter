package main

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func getBuddyResponse(config buddy) (buddyVersion, buddyFiles, buddyJob, buddyPrinter, error) {
	version := accessBuddyApi("version", config.Address, config.Apikey, config.Username, config.Pass)
	files := accessBuddyApi("files", config.Address, config.Apikey, config.Username, config.Pass)
	job := accessBuddyApi("job", config.Address, config.Apikey, config.Username, config.Pass)
	printer := accessBuddyApi("printer", config.Address, config.Apikey, config.Username, config.Pass)
	var resultVersion buddyVersion
	var resultFiles buddyFiles
	var resultJob buddyJob
	var resultPrinter buddyPrinter

	var e error

	if e = json.Unmarshal(version, &resultVersion); e != nil {
		log.Error().Msg("Can not unmarshal version JSON")
	}

	if e = json.Unmarshal(files, &resultFiles); e != nil {
		log.Error().Msg("Can not unmarshal files JSON")
	}

	if e = json.Unmarshal(job, &resultJob); e != nil {
		log.Error().Msg("Can not unmarshal job JSON")
	}

	if e = json.Unmarshal(printer, &resultPrinter); e != nil {
		log.Error().Msg("Can not unmarshal printer JSON")
	}

	return resultVersion, resultFiles, resultJob, resultPrinter, e

}
