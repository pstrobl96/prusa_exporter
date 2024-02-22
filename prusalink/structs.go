package prusalink

// Version is a struct that holds the version information of the printer - buddy, einsy and sl
type Version struct {
	API          string `json:"api"`
	Server       string `json:"server"`
	Original     string `json:"original"`
	Text         string `json:"text"`
	Firmware     string `json:"firmware"`
	Sdk          string `json:"sdk"`
	Capabilities struct {
		UploadByPut bool `json:"upload-by-put"`
	} `json:"capabilities"`
	Hostname string `json:"hostname"`
}

// Job is a struct that contains data about print job
type Job struct {
	State string `json:"state"`
	Job   struct {
		EstimatedPrintTime float64 `json:"estimatedPrintTime"`
		File               struct {
			Name    string  `json:"name"`
			Path    string  `json:"path"`
			Display string  `json:"display"`
			Size    float64 `json:"size"`
			Origin  string  `json:"origin"`
			Date    float64 `json:"date"`
		} `json:"file"`
		AveragePrintTime any    `json:"averagePrintTime"`
		LastPrintTime    any    `json:"lastPrintTime"`
		Filament         any    `json:"filament"`
		User             string `json:"user"`
	} `json:"job"`
	Progress struct {
		PrintTimeLeft       float64 `json:"printTimeLeft"`
		Completion          float64 `json:"completion"`
		PrintTime           float64 `json:"printTime"`
		Filepos             float64 `json:"filepos"`
		PrintTimeLeftOrigin string  `json:"printTimeLeftOrigin"`
		PosZMm              float64 `json:"pos_z_mm"`
		PrintSpeed          float64 `json:"printSpeed"`
		FlowFactor          float64 `json:"flow_factor"`
	} `json:"progress"`
}

// Printer is a struct that contains data about the printer - merged buddy and einsy
type Printer struct {
	Telemetry struct {
		TempBed     float64 `json:"temp-bed"`
		TempNozzle  float64 `json:"temp-nozzle"`
		PrintSpeed  float64 `json:"print-speed"`
		ZHeight     float64 `json:"z-height"`
		Material    string  `json:"material"`
		AxisX       float64 `json:"axis_x"`
		AxisY       float64 `json:"axis_y"`
		AxisZ       float64 `json:"axis_z"`
		CoverClosed bool    `json:"coverClosed"`
		FanBlower   float64 `json:"fanBlower"`
		FanRear     float64 `json:"fanRear"`
		FanUvLed    float64 `json:"fanUvLed"`
		TempAmbient float64 `json:"tempAmbient"`
		TempCPU     float64 `json:"tempCpu"`
		TempUvLed   float64 `json:"tempUvLed"`
	} `json:"telemetry"`
	Temperature struct {
		Tool0 struct {
			Actual  float64 `json:"actual"`
			Target  float64 `json:"target"`
			Display float64 `json:"display"`
			Offset  float64 `json:"offset"`
		} `json:"tool0"`
		Bed struct {
			Actual float64 `json:"actual"`
			Target float64 `json:"target"`
			Offset float64 `json:"offset"`
		} `json:"bed"`
		Chamber struct {
			Actual float64 `json:"actual"`
			Offset float64 `json:"offset"`
			Target float64 `json:"target"`
		} `json:"chamber"`
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
			ClosedOrError bool   `json:"closedOrError"`
			Finished      bool   `json:"finished"`
			Prepared      bool   `json:"prepared"`
		} `json:"flags"`
	} `json:"state"`
	Storage struct {
		Local struct {
			FreeSpace  float64 `json:"free_space"`
			TotalSpace float64 `json:"total_space"`
		} `json:"local"`
		SdCard any `json:"sd_card"`
	} `json:"storage"`
}

// Files is a struct that contains data about the files on the printer
type Files struct {
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
		Date     float64  `json:"date"`
		Size     float64  `json:"size"`
		TypePath []string `json:"typePath"`
		Refs     struct {
			Resource any `json:"resource"`
		} `json:"refs"`
		ReadOnly bool `json:"read_only,omitempty"`
	} `json:"files"`
}

// JobV1 is a struct that contains data about the print job from path /api/v1/job
type JobV1 struct {
	ID                  float64 `json:"id"`
	State               string  `json:"state"`
	Progress            float64 `json:"progress"`
	TimeRemaining       float64 `json:"time_remaining"`
	TimePrinting        float64 `json:"time_printing"`
	InaccurateEstimates bool    `json:"inaccurate_estimates"`
	File                struct {
		Refs struct {
			Icon      any `json:"icon"`
			Thumbnail any `json:"thumbnail"`
			Download  any `json:"download"`
		} `json:"refs"`
		Name        string  `json:"name"`
		DisplayName string  `json:"display_name"`
		Path        string  `json:"path"`
		Size        float64 `json:"size"`
		MTimestamp  float64 `json:"m_timestamp"`
		DisplayPath string  `json:"display_path"`
		Meta        struct {
			EstimatedPrintingTimeNormalMode string  `json:"estimated printing time (normal mode)"`
			PrinterModel                    string  `json:"printer_model"`
			LayerHeight                     float64 `json:"layer_height"`
			FilamentType                    string  `json:"filament_type"`
			EstimatedPrintTime              float64 `json:"estimated_print_time"`
		} `json:"meta"`
	} `json:"file"`
}

// Status is struct that returns /api/v1/status endpoint. Unfortunately, Buddy returns different schema, than Einsy and second struct is needed
type Status struct {
	Job struct {
		ID            float64 `json:"id"`
		Progress      float64 `json:"progress"`
		TimeRemaining float64 `json:"time_remaining"`
		TimePrinting  float64 `json:"time_printing"`
	} `json:"job"`
	Printer struct {
		State        string  `json:"state"`
		TempBed      float64 `json:"temp_bed"`
		TargetBed    float64 `json:"target_bed"`
		TempNozzle   float64 `json:"temp_nozzle"`
		TargetNozzle float64 `json:"target_nozzle"`
		AxisX        float64 `json:"axis_x"`
		AxisY        float64 `json:"axis_y"`
		AxisZ        float64 `json:"axis_z"`
		Flow         float64 `json:"flow"`
		Speed        float64 `json:"speed"`
		FanHotend    float64 `json:"fan_hotend"`
		FanPrint     float64 `json:"fan_print"`
	} `json:"printer"`
}

// StorageV1 is a struct that contains data about the storage from path /api/v1/storage
type StorageV1 struct {
	StorageList []struct {
		Path        string  `json:"path"`
		Name        string  `json:"name"`
		Type        string  `json:"type"`
		ReadOnly    bool    `json:"read_only"`
		Available   bool    `json:"available"`
		FreeSpace   float64 `json:"free_space,omitempty"`
		TotalSpace  float64 `json:"total_space,omitempty"`
		PrintFiles  float64 `json:"print_files"`
		SystemFiles float64 `json:"system_files"`
	} `json:"storage_list"`
}

// Info is a struct that contains data about the printer
type Info struct {
	Mmu               bool    `json:"mmu"`
	Name              string  `json:"name"`
	Location          string  `json:"location"`
	FarmMode          bool    `json:"farm_mode"`
	NetworkErrorChime bool    `json:"network_error_chime"`
	NozzleDiameter    float64 `json:"nozzle_diameter"`
	MinExtrusionTemp  float64 `json:"min_extrusion_temp"`
	Serial            string  `json:"serial"`
	Hostname          string  `json:"hostname"`
	Port              float64 `json:"port"`
}

// PrinterProfiles is a struct that contains data about the printer profiles
type PrinterProfiles struct {
	Profiles []struct {
		Color    string `json:"color"`
		Current  bool   `json:"current"`
		Default  bool   `json:"default"`
		Extruder struct {
			Count   float64 `json:"count"`
			Offsets []int   `json:"offsets"`
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

// Settings is a struct that contains data about the printer settings
type Settings struct {
	APIKey   string `json:"api-key"`
	Username string `json:"username"`
	Printer  struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		FarmMode bool   `json:"farm_mode"`
	} `json:"printer"`
}

// Cameras is a struct that contains data about the cameras connected to the printer
type Cameras struct {
	CameraList []struct {
		CameraID string `json:"camera_id"`
		Config   struct {
			IDString      string `json:"id_string"`
			Name          string `json:"name"`
			Driver        string `json:"driver"`
			Resolution    string `json:"resolution"`
			TriggerScheme string `json:"trigger_scheme"`
		} `json:"config"`
		Connected  bool `json:"connected"`
		Detected   bool `json:"detected"`
		Stored     bool `json:"stored"`
		Registered bool `json:"registered"`
	} `json:"camera_list"`
}
