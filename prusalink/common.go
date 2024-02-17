package prusalink

import (
	"net/http"

	"github.com/icholy/digest"
	"github.com/pstrobl96/prusa_exporter/config"
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
)

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
func accessPrinterEndpoint(path string, address string, printer config.Printers) (*http.Response, error) {
	url := string("http://" + address + "/api/" + path)
	var res *http.Response

	if printer.Apikey == "" {
		client := &http.Client{
			Transport: &digest.Transport{
				Username: printer.Username,
				Password: printer.Password,
			},
		}
		res, err := client.Get(url)

		if err != nil {
			return res, err
		}
	} else {
		req, err := http.NewRequest("GET", url, nil)
		client := &http.Client{}

		if err != nil {
			return res, err
		}

		req.Header.Add("X-Api-Key", printer.Apikey)
		res, err := client.Do(req)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}
