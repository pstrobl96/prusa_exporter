package prusalink

import (
	"net/http"

	"github.com/icholy/digest"
	"github.com/pstrobl96/prusa_exporter/config"
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
	} else if printer.State.Flags.ClosedOrError {
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
