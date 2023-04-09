package main

type legacyTelemetry struct {
	TempNozzle    int     `json:"temp_nozzle"`
	TempBed       int     `json:"temp_bed"`
	Material      string  `json:"material"`
	PosZMm        float64 `json:"pos_z_mm"`
	PrintingSpeed int     `json:"printing_speed"`
	FlowFactor    int     `json:"flow_factor"`
	Progress      int     `json:"progress"`
	PrintDur      string  `json:"print_dur"`
	TimeEst       string  `json:"time_est"`
	ProjectName   string  `json:"project_name"`
}
