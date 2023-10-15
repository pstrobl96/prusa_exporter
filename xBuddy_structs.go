package main

/*
type buddyVersion struct {
	API      string `json:"api"`
	Server   string `json:"server"`
	Text     string `json:"text"`
	Hostname string `json:"hostname"`
}
*/

type buddyVersion struct {
	API            string  `json:"api"`
	Server         string  `json:"server"`
	NozzleDiameter float64 `json:"nozzle_diameter"`
	Text           string  `json:"text"`
	Hostname       string  `json:"hostname"`
	Capabilities   struct {
		UploadByPut bool `json:"upload-by-put"`
	} `json:"capabilities"`
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
			LinkState     string `json:"link_state"`
			Operational   bool   `json:"operational"`
			Paused        bool   `json:"paused"`
			Printing      bool   `json:"printing"`
			Cancelling    bool   `json:"cancelling"`
			Pausing       bool   `json:"pausing"`
			Error         bool   `json:"error"`
			SdReady       bool   `json:"sdReady"`
			ClosedOnError bool   `json:"closedOnError"`
			Ready         bool   `json:"ready"`
			Busy          bool   `json:"busy"`
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

type buddyInfo struct {
	NozzleDiameter   float64 `json:"nozzle_diameter"`
	Mmu              bool    `json:"mmu"`
	Serial           string  `json:"serial"`
	Hostname         string  `json:"hostname"`
	MinExtrusionTemp int     `json:"min_extrusion_temp"`
}

type buddyStorage struct {
	StorageList []struct {
		Path      string `json:"path"`
		Name      string `json:"name"`
		Type      string `json:"type"`
		ReadOnly  bool   `json:"read_only"`
		Available bool   `json:"available"`
	} `json:"storage_list"`
}

type buddyStatus struct {
	Job struct {
		ID            int     `json:"id"`
		Progress      float64 `json:"progress"`
		TimeRemaining int     `json:"time_remaining"`
		TimePrinting  int     `json:"time_printing"`
	} `json:"job"`
	Storage struct {
		Path     string `json:"path"`
		Name     string `json:"name"`
		ReadOnly bool   `json:"read_only"`
	} `json:"storage"`
	Printer struct {
		State        string  `json:"state"`
		TempBed      float64 `json:"temp_bed"`
		TargetBed    float64 `json:"target_bed"`
		TempNozzle   float64 `json:"temp_nozzle"`
		TargetNozzle float64 `json:"target_nozzle"`
		AxisX        float64 `json:"axis_x"`
		AxisY        float64 `json:"axis_y"`
		AxisZ        float64 `json:"axis_z"`
		Flow         int     `json:"flow"`
		Speed        int     `json:"speed"`
		FanHotend    int     `json:"fan_hotend"`
		FanPrint     int     `json:"fan_print"`
	} `json:"printer"`
}
