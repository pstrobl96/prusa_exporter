package prusalink

type einsyJob struct {
	Job struct {
		EstimatedPrintTime int `json:"estimatedPrintTime"`
		AveragePrintTime   any `json:"averagePrintTime"`
		LastPrintTime      any `json:"lastPrintTime"`
		Filament           any `json:"filament"`
		File               struct {
			Name    string `json:"name"`
			Path    string `json:"path"`
			Size    int    `json:"size"`
			Origin  string `json:"origin"`
			Date    int    `json:"date"`
			Display string `json:"display"`
		} `json:"file"`
		User string `json:"user"`
	} `json:"job"`
	Progress struct {
		Completion          float64 `json:"completion"`
		Filepos             int     `json:"filepos"`
		PrintTime           int     `json:"printTime"`
		PrintTimeLeft       int     `json:"printTimeLeft"`
		PrintTimeLeftOrigin string  `json:"printTimeLeftOrigin"`
		PosZMm              float64 `json:"pos_z_mm"`
		PrintSpeed          int     `json:"printSpeed"`
		FlowFactor          int     `json:"flow_factor"`
	} `json:"progress"`
	State string `json:"state"`
}

type einsyCameras struct {
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

type einsyPrinter struct {
	Temperature struct {
		Tool0 struct {
			Actual float64 `json:"actual"`
			Target float64 `json:"target"`
		} `json:"tool0"`
		Bed struct {
			Actual float64 `json:"actual"`
			Target float64 `json:"target"`
		} `json:"bed"`
	} `json:"temperature"`
	Sd struct {
		Ready bool `json:"ready"`
	} `json:"sd"`
	State struct {
		Text  string `json:"text"`
		Flags struct {
			Operational   bool   `json:"operational"`
			Paused        bool   `json:"paused"`
			Printing      bool   `json:"printing"`
			Cancelling    bool   `json:"cancelling"`
			Pausing       bool   `json:"pausing"`
			SdReady       bool   `json:"sdReady"`
			Error         bool   `json:"error"`
			Ready         bool   `json:"ready"`
			ClosedOrError bool   `json:"closedOrError"`
			Finished      bool   `json:"finished"`
			Prepared      bool   `json:"prepared"`
			LinkState     string `json:"link_state"`
		} `json:"flags"`
	} `json:"state"`
	Telemetry struct {
		TempBed    float64 `json:"temp-bed"`
		TempNozzle float64 `json:"temp-nozzle"`
		Material   string  `json:"material"`
		ZHeight    float64 `json:"z-height"`
		PrintSpeed int     `json:"print-speed"`
		AxisX      float64 `json:"axis_x"`
		AxisY      float64 `json:"axis_y"`
		AxisZ      float64 `json:"axis_z"`
	} `json:"telemetry"`
	Storage struct {
		Local struct {
			FreeSpace  int64 `json:"free_space"`
			TotalSpace int64 `json:"total_space"`
		} `json:"local"`
		SdCard any `json:"sd_card"`
	} `json:"storage"`
}

// type einsyStorage struct { // currently unused
// 	StorageList []struct {
// 		Type        string `json:"type"`
// 		Path        string `json:"path"`
// 		Available   bool   `json:"available"`
// 		FreeSpace   int64  `json:"free_space,omitempty"`
// 		TotalSpace  int64  `json:"total_space,omitempty"`
// 		ReadOnly    bool   `json:"read_only"`
// 		Name        string `json:"name"`
// 		PrintFiles  int    `json:"print_files"`
// 		SystemFiles int    `json:"system_files"`
// 	} `json:"storage_list"`
// }

type einsySettings struct {
	APIKey   string `json:"api-key"`
	Username string `json:"username"`
	Printer  struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		FarmMode bool   `json:"farm_mode"`
	} `json:"printer"`
}

// type einsyConection struct { // currently unused
// 	Current struct {
// 		Baudrate       int    `json:"baudrate"`
// 		Port           string `json:"port"`
// 		PrinterProfile string `json:"printerProfile"`
// 		State          string `json:"state"`
// 	} `json:"current"`
// 	Options struct {
// 		Ports           []string `json:"ports"`
// 		Baudrates       []int    `json:"baudrates"`
// 		PrinterProfiles []struct {
// 			ID   string `json:"id"`
// 			Name string `json:"name"`
// 		} `json:"printerProfiles"`
// 		Autoconnect bool `json:"autoconnect"`
// 	} `json:"options"`
// 	Connect struct {
// 		Hostname     string `json:"hostname"`
// 		Port         int    `json:"port"`
// 		TLS          bool   `json:"tls"`
// 		Registration string `json:"registration"`
// 		Code         any    `json:"code"`
// 	} `json:"connect"`
// 	States struct {
// 		Printer struct {
// 			Ok      bool   `json:"ok"`
// 			Message string `json:"message"`
// 		} `json:"printer"`
// 		Connect struct {
// 			Ok      bool   `json:"ok"`
// 			Message string `json:"message"`
// 		} `json:"connect"`
// 	} `json:"states"`
// }

type einsyFiles struct {
	Files []struct {
		Name     string   `json:"name"`
		Path     string   `json:"path"`
		Display  string   `json:"display"`
		Date     int      `json:"date"`
		Size     int      `json:"size"`
		Type     string   `json:"type"`
		TypePath []string `json:"typePath"`
		Origin   string   `json:"origin"`
		Refs     struct {
			Resource any `json:"resource"`
		} `json:"refs"`
		Children []struct {
			Name     string   `json:"name"`
			Path     string   `json:"path"`
			Display  string   `json:"display"`
			Date     int      `json:"date"`
			Size     int      `json:"size"`
			Type     string   `json:"type"`
			TypePath []string `json:"typePath"`
			Origin   string   `json:"origin"`
			Refs     struct {
				Resource any `json:"resource"`
			} `json:"refs,omitempty"`
			Children []struct {
				Name     string   `json:"name"`
				Path     string   `json:"path"`
				Display  string   `json:"display"`
				Date     int      `json:"date"`
				Size     int      `json:"size"`
				Origin   string   `json:"origin"`
				Type     string   `json:"type"`
				TypePath []string `json:"typePath"`
				Hash     any      `json:"hash"`
				Refs     struct {
					Download  string `json:"download"`
					Icon      any    `json:"icon"`
					Thumbnail string `json:"thumbnail"`
				} `json:"refs"`
				GcodeAnalysis struct {
					EstimatedPrintTime int     `json:"estimatedPrintTime"`
					Material           string  `json:"material"`
					LayerHeight        float64 `json:"layerHeight"`
				} `json:"gcodeAnalysis"`
			} `json:"children,omitempty"`
			Hash          any `json:"hash,omitempty"`
			GcodeAnalysis struct {
				EstimatedPrintTime int     `json:"estimatedPrintTime"`
				Material           string  `json:"material"`
				LayerHeight        float64 `json:"layerHeight"`
			} `json:"gcodeAnalysis,omitempty"`
		} `json:"children"`
		Ro bool `json:"ro,omitempty"`
	} `json:"files"`
	Free  string `json:"free"`
	Total string `json:"total"`
}

type einsyLogs struct {
	Files []struct {
		Name string `json:"name"`
		Size int    `json:"size"`
		Date int    `json:"date"`
	} `json:"files"`
}

type einsyInfo struct {
	Name             string  `json:"name"`
	Location         string  `json:"location"`
	FarmMode         bool    `json:"farm_mode"`
	NozzleDiameter   float64 `json:"nozzle_diameter"`
	MinExtrusionTemp int     `json:"min_extrusion_temp"`
	Serial           string  `json:"serial"`
	Hostname         string  `json:"hostname"`
	Port             int     `json:"port"`
}

// type einsyStatus struct { // currently unused
// 	Storage []struct {
// 		Path      string `json:"path"`
// 		ReadOnly  bool   `json:"read_only"`
// 		FreeSpace int64  `json:"free_space,omitempty"`
// 		Name      string `json:"name"`
// 	} `json:"storage"`
// 	Printer struct {
// 		State         string  `json:"state"`
// 		TempNozzle    float64 `json:"temp_nozzle"`
// 		TempBed       float64 `json:"temp_bed"`
// 		AxisZ         float64 `json:"axis_z"`
// 		Flow          int     `json:"flow"`
// 		Speed         int     `json:"speed"`
// 		FanHotend     int     `json:"fan_hotend"`
// 		FanPrint      int     `json:"fan_print"`
// 		StatusConnect struct {
// 			Ok      bool   `json:"ok"`
// 			Message string `json:"message"`
// 		} `json:"status_connect"`
// 		StatusPrinter struct {
// 			Ok      bool   `json:"ok"`
// 			Message string `json:"message"`
// 		} `json:"status_printer"`
// 		TargetNozzle float64 `json:"target_nozzle"`
// 		TargetBed    float64 `json:"target_bed"`
// 	} `json:"printer"`
// 	Camera struct {
// 		ID string `json:"id"`
// 	} `json:"camera"`
// 	Job struct {
// 		ID            int     `json:"id"`
// 		Progress      float64 `json:"progress"`
// 		TimeRemaining int     `json:"time_remaining"`
// 		TimePrinting  int     `json:"time_printing"`
// 	} `json:"job"`
// }

type einsyVersion struct {
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

type einsyPorts struct {
	Ports []struct {
		Path        string `json:"path"`
		IsRpiPort   bool   `json:"is_rpi_port"`
		Checked     bool   `json:"checked"`
		Usable      bool   `json:"usable"`
		Selected    bool   `json:"selected"`
		Description string `json:"description"`
		Baudrate    int    `json:"baudrate"`
		Timeout     int    `json:"timeout"`
		Sn          any    `json:"sn"`
	} `json:"ports"`
}
