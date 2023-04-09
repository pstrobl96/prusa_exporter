package main

import (
	"encoding/json"
	"log"
)

func getLegacyTelemetry(address string) legacyTelemetry {
	resp := accessLegacyApi("telemetry", address)

	var result legacyTelemetry

	if err := json.Unmarshal(resp, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}
