package main

type buddyVersion struct {
	API      string `json:"api"`
	Server   string `json:"server"`
	Text     string `json:"text"`
	Hostname string `json:"hostname"`
}

type buddyFiles struct {
	Files []struct {
		Name     string `json:"name"`
		Path     string `json:"path"`
		Display  string `json:"display"`
		Type     string `json:"type"`
		Origin   string `json:"origin"`
		Children []struct {
			Name    string `json:"name"`
			Display string `json:"display"`
			Path    string `json:"path"`
			Origin  string `json:"origin"`
			Refs    struct {
				Resource       string `json:"resource"`
				ThumbnailSmall string `json:"thumbnailSmall"`
				ThumbnailBig   string `json:"thumbnailBig"`
				Download       string `json:"download"`
			} `json:"refs"`
		} `json:"children"`
	} `json:"files"`
}

type buddyPrinter struct {
	Telemetry struct {
		TempBed    float64 `json:"temp-bed"`
		TempNozzle float64 `json:"temp-nozzle"`
		PrintSpeed int     `json:"print-speed"`
		ZHeight    float64 `json:"z-height"`
		Material   string  `json:"material"`
	} `json:"telemetry"`
	Temperature struct {
		Tool0 struct {
			Actual  float64 `json:"actual"`
			Target  float64 `json:"target"`
			Display float64 `json:"display"`
			Offset  int     `json:"offset"`
		} `json:"tool0"`
		Bed struct {
			Actual float64 `json:"actual"`
			Target float64 `json:"target"`
			Offset int     `json:"offset"`
		} `json:"bed"`
	} `json:"temperature"`
	State struct {
		Text  string `json:"text"`
		Flags struct {
			Operational   bool `json:"operational"`
			Paused        bool `json:"paused"`
			Printing      bool `json:"printing"`
			Cancelling    bool `json:"cancelling"`
			Pausing       bool `json:"pausing"`
			SdReady       bool `json:"sdReady"`
			Error         bool `json:"error"`
			ClosedOnError bool `json:"closedOnError"`
			Ready         bool `json:"ready"`
			Busy          bool `json:"busy"`
		} `json:"flags"`
	} `json:"state"`
}

type buddyJob struct {
	State string `json:"state"`
	Job   struct {
		EstimatedPrintTime int `json:"estimatedPrintTime"`
		File               struct {
			Name    string `json:"name"`
			Path    string `json:"path"`
			Display string `json:"display"`
		} `json:"file"`
	} `json:"job"`
	Progress struct {
		PrintTimeLeft int     `json:"printTimeLeft"`
		Completion    float64 `json:"completion"`
		PrintTime     int     `json:"printTime"`
	} `json:"progress"`
}

