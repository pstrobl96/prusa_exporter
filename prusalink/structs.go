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
		EstimatedPrintTime int `json:"estimatedPrintTime"`
		File               struct {
			Name    string `json:"name"`
			Path    string `json:"path"`
			Display string `json:"display"`
			Size    int    `json:"size"`
			Origin  string `json:"origin"`
			Date    int    `json:"date"`
		} `json:"file"`
		AveragePrintTime any    `json:"averagePrintTime"`
		LastPrintTime    any    `json:"lastPrintTime"`
		Filament         any    `json:"filament"`
		User             string `json:"user"`
	} `json:"job"`
	Progress struct {
		PrintTimeLeft       int     `json:"printTimeLeft"`
		Completion          int     `json:"completion"`
		PrintTime           int     `json:"printTime"`
		Filepos             int     `json:"filepos"`
		PrintTimeLeftOrigin string  `json:"printTimeLeftOrigin"`
		PosZMm              float64 `json:"pos_z_mm"`
		PrintSpeed          int     `json:"printSpeed"`
		FlowFactor          int     `json:"flow_factor"`
	} `json:"progress"`
}

// Printer is a struct that contains data about the printer - merged buddy and einsy
type Printer struct {
	Telemetry struct {
		TempBed    float64 `json:"temp-bed"`
		TempNozzle float64 `json:"temp-nozzle"`
		PrintSpeed int     `json:"print-speed"`
		ZHeight    float64 `json:"z-height"`
		Material   string  `json:"material"`
		AxisX      float64 `json:"axis_x"`
		AxisY      float64 `json:"axis_y"`
		AxisZ      float64 `json:"axis_z"`
	} `json:"telemetry"`
	Temperature struct {
		Tool0 struct {
			Actual  float64 `json:"actual"`
			Target  int     `json:"target"`
			Display int     `json:"display"`
			Offset  int     `json:"offset"`
		} `json:"tool0"`
		Bed struct {
			Actual float64 `json:"actual"`
			Target int     `json:"target"`
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
			ClosedOrError bool   `json:"closedOrError"`
			Finished      bool   `json:"finished"`
			Prepared      bool   `json:"prepared"`
		} `json:"flags"`
	} `json:"state"`
	Sd struct {
		Ready bool `json:"ready"`
	} `json:"sd"`
	Storage struct {
		Local struct {
			FreeSpace  int64 `json:"free_space"`
			TotalSpace int64 `json:"total_space"`
		} `json:"local"`
		SdCard any `json:"sd_card"`
	} `json:"storage"`
}

// Files is a struct that contains data about the files on the printer
type Files struct {
	Free  string `json:"free"`
	Total string `json:"total"`
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
		Date     int      `json:"date"`
		Size     int      `json:"size"`
		TypePath []string `json:"typePath"`
		Refs     struct {
			Resource any `json:"resource"`
		} `json:"refs"`
		ReadOnly bool `json:"read_only,omitempty"`
	} `json:"files"`
}

// JobV1 is a struct that contains data about the print job from path /api/v1/job
type JobV1 struct {
	ID                  int    `json:"id"`
	State               string `json:"state"`
	Progress            int    `json:"progress"`
	TimeRemaining       int    `json:"time_remaining"`
	TimePrinting        int    `json:"time_printing"`
	InaccurateEstimates bool   `json:"inaccurate_estimates"`
	File                struct {
		Refs struct {
			Icon      any `json:"icon"`
			Thumbnail any `json:"thumbnail"`
			Download  any `json:"download"`
		} `json:"refs"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Path        string `json:"path"`
		Size        int    `json:"size"`
		MTimestamp  int    `json:"m_timestamp"`
		DisplayPath string `json:"display_path"`
		Meta        struct {
			EstimatedPrintingTimeNormalMode string  `json:"estimated printing time (normal mode)"`
			PrinterModel                    string  `json:"printer_model"`
			LayerHeight                     float64 `json:"layer_height"`
			FilamentType                    string  `json:"filament_type"`
			EstimatedPrintTime              int     `json:"estimated_print_time"`
		} `json:"meta"`
	} `json:"file"`
}

// StatusV1 is a struct that contains data about the printer status from path /api/v1/status
type StatusV1 struct {
	Storage []struct {
		Path      string `json:"path"`
		ReadOnly  bool   `json:"read_only"`
		FreeSpace int64  `json:"free_space,omitempty"`
		Name      string `json:"name"`
	} `json:"storage"`
	Printer struct {
		State         string  `json:"state"`
		TempNozzle    int     `json:"temp_nozzle"`
		TempBed       float64 `json:"temp_bed"`
		AxisZ         int     `json:"axis_z"`
		AxisY         int     `json:"axis_y"`
		AxisX         int     `json:"axis_x"`
		Flow          int     `json:"flow"`
		Speed         int     `json:"speed"`
		FanHotend     int     `json:"fan_hotend"`
		FanPrint      int     `json:"fan_print"`
		StatusConnect struct {
			Ok      bool   `json:"ok"`
			Message string `json:"message"`
		} `json:"status_connect"`
		StatusPrinter struct {
			Ok      bool   `json:"ok"`
			Message string `json:"message"`
		} `json:"status_printer"`
		TargetNozzle int `json:"target_nozzle"`
		TargetBed    int `json:"target_bed"`
	} `json:"printer"`
	Job struct {
		ID            int `json:"id"`
		Progress      int `json:"progress"`
		TimeRemaining int `json:"time_remaining"`
		TimePrinting  int `json:"time_printing"`
	} `json:"job"`
}

// StorageV1 is a struct that contains data about the storage from path /api/v1/storage
type StorageV1 struct {
	StorageList []struct {
		Path        string `json:"path"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		ReadOnly    bool   `json:"read_only"`
		Available   bool   `json:"available"`
		FreeSpace   int64  `json:"free_space,omitempty"`
		TotalSpace  int64  `json:"total_space,omitempty"`
		PrintFiles  int    `json:"print_files"`
		SystemFiles int    `json:"system_files"`
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
	MinExtrusionTemp  int     `json:"min_extrusion_temp"`
	Serial            string  `json:"serial"`
	Hostname          string  `json:"hostname"`
	Port              int     `json:"port"`
}

// PrinterProfiles is a struct that contains data about the printer profiles
type PrinterProfiles struct {
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
