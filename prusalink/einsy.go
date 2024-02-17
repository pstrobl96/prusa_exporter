package prusalink

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/rs/zerolog/log"
)

func getEinsyResponse(printer config.Printers) (Version, Files, Job, Printer, Info, error) {
	var (
		version     Version
		files       Files
		job         Job
		printerData Printer
		info        Info
		err         error
	)

	version, err = GetVersion(printer)

	if err != nil {
		log.Error().Msg("Error getting version" + err.Error())
	}

	files, err = GetFiles(printer)

	if err != nil {
		log.Error().Msg("Error getting files" + err.Error())
	}

	job, err = GetJob(printer)

	if err != nil {
		log.Error().Msg("Error getting job" + err.Error())
	}

	printerData, err = GetPrinter(printer)

	if err != nil {
		log.Error().Msg("Error getting printer" + err.Error())
	}

	info, err = GetInfo(printer)

	if err != nil {
		log.Error().Msg("Error getting info" + err.Error())
	}

	return version, files, job, printerData, info, err
}

type einsyCollector struct {
	printerNozzleTemp         *prometheus.Desc
	printerBedTemp            *prometheus.Desc
	printerPrintSpeedRatio    *prometheus.Desc
	printerTargetTempNozzle   *prometheus.Desc
	printerTargetTempBed      *prometheus.Desc
	printerFiles              *prometheus.Desc
	printerPrintTime          *prometheus.Desc
	printerPrintTimeRemaining *prometheus.Desc
	printerPrintProgress      *prometheus.Desc
	printerMaterial           *prometheus.Desc
	printerLogs               *prometheus.Desc
	printerLogsDate           *prometheus.Desc
	printerInfo               *prometheus.Desc
	printerFarmMode           *prometheus.Desc
	printerCameras            *prometheus.Desc
	printerAxisX              *prometheus.Desc
	printerAxisY              *prometheus.Desc
	printerAxisZ              *prometheus.Desc
	printerStatus             *prometheus.Desc
	printerNozzleSize         *prometheus.Desc
	printerUp                 *prometheus.Desc
}

func newEinsyCollector() *einsyCollector {
	return &einsyCollector{
		printerNozzleTemp:         prometheus.NewDesc("prusa_einsy_nozzle_temperature", "Current temperature of printer nozzle in Celsius", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerBedTemp:            prometheus.NewDesc("prusa_einsy_bed_temperature", "Current temperature of printer bed in Celsius", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerPrintSpeedRatio:    prometheus.NewDesc("prusa_einsy_print_speed_ratio", "Current setting of printer speed in values from 0.0 - 1.0", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerTargetTempNozzle:   prometheus.NewDesc("prusa_einsy_nozzle_target_temperature", "Target temperature of printer nozzle in Celsius", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerTargetTempBed:      prometheus.NewDesc("prusa_einsy_bed_target_temperature", "Target temperature of printer bed in Celsius", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerFiles:              prometheus.NewDesc("prusa_einsy_files", "Number of files in storage", []string{"printer_address", "printer_model", "printer_name", "printer_storage"}, nil),
		printerPrintTime:          prometheus.NewDesc("prusa_einsy_print_time", "Returns actual printing time of current print", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_einsy_printing_time_remaining", "Returns time that remains for completion of current print", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerPrintProgress:      prometheus.NewDesc("prusa_einsy_printing_progress", "Returns information about completion of current print in percents", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerMaterial:           prometheus.NewDesc("prusa_einsy_material", "Returns information about loaded filament. Returns 0 if there is no loaded filament", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "printer_filament"}, nil),
		printerLogs:               prometheus.NewDesc("prusa_einsy_logs", "Return size of logs in Prusa Link", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "log_name"}, nil),
		printerLogsDate:           prometheus.NewDesc("prusa_einsy_logs_date", "Return date of logs in Prusa Link", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "log_name"}, nil),
		printerInfo:               prometheus.NewDesc("prusa_einsy_info", "Return info about printer", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "printer_api", "printer_server", "printer_text", "printer_link_name", "printer_location", "printer_sn", "printer_hostname", "printer_type"}, nil),
		printerFarmMode:           prometheus.NewDesc("prusa_einsy_farm_mode", "Return if printer is set to farm mode", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerCameras:            prometheus.NewDesc("prusa_einsy_cameras", "Return information about cameras", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "camera_id", "camera_name", "camera_resolution"}, nil),
		printerAxisX:              prometheus.NewDesc("prusa_einsy_axis_x", "Return coordinates - x axis of printer", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerAxisY:              prometheus.NewDesc("prusa_einsy_axis_y", "Return coordinates - y axis of printer", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerAxisZ:              prometheus.NewDesc("prusa_einsy_axis_z", "Return coordinates - z axis of printer", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerStatus:             prometheus.NewDesc("prusa_einsy_status", "Return state of printer", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "printer_state"}, nil),
		printerNozzleSize:         prometheus.NewDesc("prusa_einsy_nozzle_size", "Return size of nozzle", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerUp:                 prometheus.NewDesc("prusa_einsy_up", "Return if printer is up", []string{"printer_address", "printer_model", "printer_name"}, nil),
	}
}

func (collector *einsyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printerNozzleTemp
	ch <- collector.printerBedTemp
	ch <- collector.printerPrintSpeedRatio
	ch <- collector.printerTargetTempNozzle
	ch <- collector.printerTargetTempBed
	ch <- collector.printerFiles
	ch <- collector.printerPrintTime
	ch <- collector.printerPrintTimeRemaining
	ch <- collector.printerPrintProgress
	ch <- collector.printerMaterial
	ch <- collector.printerStatus
	ch <- collector.printerAxisX
	ch <- collector.printerAxisY
	ch <- collector.printerAxisZ
	ch <- collector.printerCameras
	ch <- collector.printerFarmMode
	ch <- collector.printerInfo
	ch <- collector.printerLogsDate
	ch <- collector.printerLogs
	ch <- collector.printerUp
}
