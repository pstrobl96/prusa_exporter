package prusalink

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/icholy/digest"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/rs/zerolog/log"
)

var (
	// printerEndpoints is a map of printer names to their respective API endpoints
	printerEndpoints = map[string][]string{
		"MINI":    printerEndpointsList["buddy"],
		"MK4":     printerEndpointsList["buddy"],
		"XL":      printerEndpointsList["buddy"],
		"I3MK3S":  printerEndpointsList["einsy"],
		"I3MK3":   printerEndpointsList["einsy"],
		"I3MK25S": printerEndpointsList["einsy"],
		"I3MK25":  printerEndpointsList["einsy"],
		"SL1":     printerEndpointsList["sl"],
	}

	// printerEndpointsList is a map of printer names to their respective API endpoints
	printerEndpointsList = map[string][]string{
		"buddy": {"version", "files", "job", "printer", "v1/status", "v1/info", "v1/storage"},
		"einsy": {"version", "files", "job", "printer", "settings", "ports", "v1/cameras", "v1/status", "v1/info", "v1/storage"},
		"sl":    {"files?recursive=true", "job", "printer", "printerprofiles", "version"},
	}

	printerTypes = map[string]string{
		"PrusaMINI":         "MINI",
		"PrusaMK4":          "MK4",
		"PrusaXL":           "XL",
		"PrusaLink I3MK3S":  "I3MK3S",
		"PrusaLink I3MK3":   "I3MK3",
		"PrusaLink I3MK25S": "I3MK25S",
		"PrusaLink I3MK25":  "I3MK25",
		"prusa-sl1":         "SL1",
	}

	configuration *config.Config
)

// SetConfig is used to set the configuration
func SetConfig(config *config.Config) {
	configuration = config
}

// GetLabels is used to get the labels for the given printer and job
func GetLabels(printer config.Printers, job Job, labelValues ...string) []string {
	if job == (Job{}) {
		return append([]string{printer.Address, printer.Type, printer.Name, "", ""}, labelValues...)
	}
	return append([]string{printer.Address, printer.Type, printer.Name, job.Job.File.Name, job.Job.File.Path}, labelValues...)
}

// BoolToFloat is used for basic parsing boolean to float64
// 0.0 for false, 1.0 for true
func BoolToFloat(boolean bool) float64 {
	if !boolean {
		return 0.0
	}

	return 1.0
}

// getStateFlag returns the state flag for the given printer.
// The state flag is a float64 value representing the current state of the printer.
// It is used for tracking the printer's status and progress.
func getStateFlag(printer Printer) float64 {
	if printer.State.Flags.Operational {
		return 1
	} else if printer.State.Flags.Prepared {
		return 2
	} else if printer.State.Flags.Paused {
		return 3
	} else if printer.State.Flags.Printing {
		return 4
	} else if printer.State.Flags.Cancelling {
		return 5
	} else if printer.State.Flags.Pausing {
		return 6
	} else if printer.State.Flags.Error {
		return 7
	} else if printer.State.Flags.SdReady {
		return 8
	} else if printer.State.Flags.ClosedOrError || printer.State.Flags.ClosedOnError {
		return 9
	} else if printer.State.Flags.Ready {
		return 10
	} else if printer.State.Flags.Busy {
		return 11
	} else if printer.State.Flags.Finished {
		return 12
	} else {
		return 0
	}
}

// accessPrinterEndpoint is used to access the printer's API endpoint
func accessPrinterEndpoint(path string, printer config.Printers) ([]byte, error) {
	url := string("http://" + printer.Address + "/api/" + path)
	var (
		res    *http.Response
		result []byte
		err    error
	)

	if printer.Apikey == "" {
		client := &http.Client{
			Transport: &digest.Transport{
				Username: printer.Username,
				Password: printer.Password,
			},
		}
		res, err = client.Get(url)

		if err != nil {
			return result, err
		}
	} else {
		req, err := http.NewRequest("GET", url, nil)
		client := &http.Client{}

		if err != nil {
			return result, err
		}

		req.Header.Add("X-Api-Key", printer.Apikey)
		res, err = client.Do(req)
		if err != nil {
			return result, err
		}
	}
	result, err = io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Error().Msg(err.Error())
	}

	return result, nil
}

// GetVersion is used to get the printer's version API endpoint
func GetVersion(printer config.Printers) (Version, error) {
	var version Version
	response, err := accessPrinterEndpoint("version", printer)

	if err != nil {
		return version, err
	}

	err = json.Unmarshal(response, &version)

	return version, err
}

// GetJob is used to get the printer's job API endpoint
func GetJob(printer config.Printers) (Job, error) {
	var job Job
	response, err := accessPrinterEndpoint("job", printer)

	if err != nil {
		return job, err
	}

	err = json.Unmarshal(response, &job)

	return job, err
}

// GetPrinter is used to get the printer's printer API endpoint
func GetPrinter(printer config.Printers) (Printer, error) {
	var printerData Printer
	response, err := accessPrinterEndpoint("printer", printer)

	if err != nil {
		return printerData, err
	}

	err = json.Unmarshal(response, &printerData)

	return printerData, err
}

// GetFiles is used to get the printer's files API endpoint
func GetFiles(printer config.Printers) (Files, error) {
	var files Files
	response, err := accessPrinterEndpoint("files?recursive=true", printer)

	if err != nil {
		return files, err
	}

	err = json.Unmarshal(response, &files)

	return files, err
}

// GetJobV1 is used to get the printer's job v1 API endpoint
func GetJobV1(printer config.Printers) (JobV1, error) {
	var job JobV1
	response, err := accessPrinterEndpoint("v1/job", printer)

	if err != nil {
		return job, err
	}

	err = json.Unmarshal(response, &job)

	return job, err
}

// GetStatus is used to get Buddy status endpoint
func GetStatus(printer config.Printers) (Status, error) {
	var status Status
	response, err := accessPrinterEndpoint("v1/status", printer)

	if err != nil {
		return status, err
	}

	err = json.Unmarshal(response, &status)

	return status, err
}

// GetStatusV1 is used to get the printer's status v1 API endpoint
func GetStatusV1(printer config.Printers) (StatusV1, error) {
	var status StatusV1
	response, err := accessPrinterEndpoint("v1/status", printer)

	var objmap []map[string]interface{}
	if err := json.Unmarshal(response, &objmap); err != nil {
		fmt.Println(err)
	}
	fmt.Println(objmap) // to parse out your value

	if err != nil {
		return status, err
	}

	err = json.Unmarshal(response, &status)

	return status, err
}

// GetStorageV1 is used to get the printer's storage v1 API endpoint
func GetStorageV1(printer config.Printers) (StorageV1, error) {
	var storage StorageV1
	response, err := accessPrinterEndpoint("v1/storage", printer)

	if err != nil {
		return storage, err
	}

	err = json.Unmarshal(response, &storage)

	return storage, err
}

// GetInfo is used to get the printer's info API endpoint
func GetInfo(printer config.Printers) (Info, error) {
	var info Info
	response, err := accessPrinterEndpoint("v1/info", printer)

	if err != nil {
		return info, err
	}

	err = json.Unmarshal(response, &info)

	return info, err
}

// GetSettings is used to get the printer's settings API endpoint
func GetSettings(printer config.Printers) (Settings, error) {
	var settings Settings
	response, err := accessPrinterEndpoint("settings", printer)

	if err != nil {
		return settings, err
	}

	err = json.Unmarshal(response, &settings)

	return settings, err
}

// GetCameras is used to get the printer's cameras API endpoint
func GetCameras(printer config.Printers) (Cameras, error) {
	var cameras Cameras
	response, err := accessPrinterEndpoint("v1/cameras", printer)

	if err != nil {
		return cameras, err
	}

	err = json.Unmarshal(response, &cameras)

	return cameras, err
}

// GetPrinterProfiles is used to get the printer's printerprofiles API endpoint
func GetPrinterProfiles(printer config.Printers) (PrinterProfiles, error) {
	var profiles PrinterProfiles
	response, err := accessPrinterEndpoint("v1/printerprofiles", printer)

	if err != nil {
		return profiles, err
	}

	err = json.Unmarshal(response, &profiles)

	return profiles, err
}

// GetPrinterType returns the printer type of the given printer - e.g. "MINI", "MK4", "XL", "I3MK3S", "I3MK3", "I3MK25S",
func GetPrinterType(printer config.Printers) (string, error) {
	version, err := GetVersion(printer)
	if err != nil {
		return "", err
	}

	if printerTypes[version.Hostname] == "" {
		// If the hostname is not found in the map, try to find the original variable
		return printerTypes[version.Original], nil
	}

	return printerTypes[version.Hostname], nil
}

// ProbePrinter is used to probe the printer - just testing the connection
func ProbePrinter(printer config.Printers) (bool, error) {
	req, _ := http.NewRequest("GET", "http://"+printer.Address+"/", nil)
	client := &http.Client{Timeout: time.Duration(1) * time.Second}
	r, e := client.Do(req)

	if e != nil {
		return false, e
	}

	if r.StatusCode == 401 {
		log.Warn().Msg("401 Unauthorized, trying to access with API key")
		req.Header.Add("X-Api-Key", printer.Apikey)
		r, e = client.Do(req)
		if e != nil {
			return false, e
		}
	}

	return r.StatusCode == 200, nil
}
