package main

// SLJob is a struct to hold the response from the job API endpoint
type SLJob struct {
	State string `json:"state"`
}

// SLPrinter is a struct to hold the response from the printers API endpoint
type SLPrinter struct {
	Sd []struct {
		Ready bool `json:"ready"`
	} `json:"sd"`
	State struct {
		Flags struct {
			Cancelling    bool `json:"cancelling"`
			ClosedOrError bool `json:"closedOrError"`
			Error         bool `json:"error"`
			Operational   bool `json:"operational"`
			Paused        bool `json:"paused"`
			Pausing       bool `json:"pausing"`
			Printing      bool `json:"printing"`
			Ready         bool `json:"ready"`
			SdReady       bool `json:"sdReady"`
		} `json:"flags"`
		Text string `json:"text"`
	} `json:"state"`
	Telemetry struct {
		CoverClosed bool    `json:"coverClosed"`
		FanBlower   int     `json:"fanBlower"`
		FanRear     int     `json:"fanRear"`
		FanUvLed    int     `json:"fanUvLed"`
		TempAmbient float64 `json:"tempAmbient"`
		TempCPU     float64 `json:"tempCpu"`
		TempUvLed   float64 `json:"tempUvLed"`
	} `json:"telemetry"`
	Temperature struct {
		Bed struct {
			Actual float64 `json:"actual"`
			Offset int     `json:"offset"`
			Target int     `json:"target"`
		} `json:"bed"`
		Chamber struct {
			Actual float64 `json:"actual"`
			Offset int     `json:"offset"`
			Target int     `json:"target"`
		} `json:"chamber"`
		Tool0 struct {
			Actual float64 `json:"actual"`
			Offset int     `json:"offset"`
			Target int     `json:"target"`
		} `json:"tool0"`
	} `json:"temperature"`
}

// SLPrinterProfiles is a struct to hold the response from the printerprofiles API endpoint
type SLPrinterProfiles struct {
	Profiles []struct {
		Color    string `json:"color"`
		Current  bool   `json:"current"`
		Default  bool   `json:"default"`
		Extruder struct {
			Count   int   `json:"count"`
			Offsets []int `json:"offsets"`
		} `json:"extruder"`
		HeatedBed         bool     `json:"heatedBed"`
		HeatedChamber     bool     `json:"heatedChamber"`
		ID                string   `json:"id"`
		Model             string   `json:"model"`
		Name              string   `json:"name"`
		ProjectExtensions []string `json:"projectExtensions"`
		Resource          string   `json:"resource"`
	} `json:"profiles"`
}

// SLVersion is a struct to hold the response from the version API endpoint
type SLVersion struct {
	API      string `json:"api"`
	Hostname string `json:"hostname"`
	Server   string `json:"server"`
	Text     string `json:"text"`
}

// SLFiles is a struct to hold the response from the files?recursive=true API endpoint
type SLFiles struct {
	Files []struct {
		Path     string `json:"path"`
		Origin   string `json:"origin"`
		Type     string `json:"type"`
		Children []struct {
			Path     string `json:"path"`
			Origin   string `json:"origin"`
			Type     string `json:"type"`
			Children []struct {
				Path          string   `json:"path"`
				Origin        string   `json:"origin"`
				Type          string   `json:"type"`
				Size          int      `json:"size"`
				Name          string   `json:"name"`
				Display       string   `json:"display"`
				Date          float64  `json:"date"`
				TypePath      []string `json:"typePath"`
				GcodeAnalysis struct {
					EstimatedPrintTime int     `json:"estimatedPrintTime"`
					LayerHeight        float64 `json:"layerHeight"`
					Material           string  `json:"material"`
				} `json:"gcodeAnalysis"`
				Refs struct {
					Resource       string `json:"resource"`
					Download       string `json:"download"`
					ThumbnailSmall string `json:"thumbnailSmall"`
					ThumbnailBig   string `json:"thumbnailBig"`
				} `json:"refs"`
			} `json:"children,omitempty"`
			Name     string   `json:"name"`
			Display  string   `json:"display"`
			Date     float64  `json:"date"`
			TypePath []string `json:"typePath"`
			Size     int      `json:"size,omitempty"`
			Refs     struct {
				Resource       string `json:"resource"`
				Download       string `json:"download"`
				ThumbnailSmall string `json:"thumbnailSmall"`
				ThumbnailBig   string `json:"thumbnailBig"`
			} `json:"refs,omitempty"`
			GcodeAnalysis struct {
				EstimatedPrintTime int     `json:"estimatedPrintTime,omitempty"`
				LayerHeight        float64 `json:"layerHeight,omitempty"`
				Material           string  `json:"material,omitempty"`
			} `json:"gcodeAnalysis,omitempty"`
		} `json:"children"`
		Name     string   `json:"name"`
		Display  string   `json:"display"`
		Date     float64  `json:"date"`
		TypePath []string `json:"typePath"`
	} `json:"files"`
	Free  int `json:"free"`
	Total int `json:"total"`
}
