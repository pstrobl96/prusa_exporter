package prusalink

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

func getSLResponse(config sl) (SLFiles, SLJob, SLPrinter, SLPrinterProfiles, SLVersion, error) {
	//files := accessSLAPI("files?recursive=true", config.Address, config.Username, config.Pass)
	job := accessSLAPI("job", config.Address, config.Username, config.Pass)
	printer := accessSLAPI("printer", config.Address, config.Username, config.Pass)
	//printerProfiles := accessSLAPI("printerprofiles", config.Address, config.Username, config.Pass)
	version := accessSLAPI("version", config.Address, config.Username, config.Pass)

	var resultFiles SLFiles
	var resultJob SLJob
	var resultPrinter SLPrinter
	var resultPrinterProfiles SLPrinterProfiles
	var resultVersion SLVersion

	var e error

	log.Debug().Msg("Getting response from " + config.Address)

	//if e = json.Unmarshal(files, &resultFiles); e != nil {
	//	log.Error().Msg("Can not unmarshal files JSON")
	//}

	if e = json.Unmarshal(job, &resultJob); e != nil {
		log.Error().Msg("Can not unmarshal job JSON")
	}

	if e = json.Unmarshal(printer, &resultPrinter); e != nil {
		log.Error().Msg("Can not unmarshal printer JSON")
	}

	//if e = json.Unmarshal(printerProfiles, &resultPrinterProfiles); e != nil {
	//	log.Error().Msg("Can not unmarshal printerProfiles JSON")
	//}

	if e = json.Unmarshal(version, &resultVersion); e != nil {
		log.Error().Msg("Can not unmarshal printer JSON")
	}

	return resultFiles, resultJob, resultPrinter, resultPrinterProfiles, resultVersion, e

}
