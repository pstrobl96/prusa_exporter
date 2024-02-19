package prusalink

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/rs/zerolog/log"
)

func getBuddyResponse(printer config.Printers) (Version, Files, Job, Printer, Status, Info, StorageV1, error) {
	var (
		version     Version
		files       Files
		job         Job
		printerData Printer
		status      Status
		info        Info
		storage     StorageV1
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

	status, err = GetStatus(printer)

	if err != nil {
		log.Error().Msg("Error getting status" + err.Error())
	}

	info, err = GetInfo(printer)

	if err != nil {
		log.Error().Msg("Error getting info" + err.Error())
	}

	storage, err = GetStorageV1(printer)

	if err != nil {
		log.Error().Msg("Error getting storage" + err.Error())
	}

	return version, files, job, printerData, status, info, storage, err
}

// BuddyCollector collects buddy metrics
type BuddyCollector struct {
	printerNozzleTemp         *prometheus.Desc
	printerBedTemp            *prometheus.Desc
	printerVersion            *prometheus.Desc
	printerZHeight            *prometheus.Desc
	printerPrintSpeed         *prometheus.Desc
	printerTargetTempNozzle   *prometheus.Desc
	printerTargetTempBed      *prometheus.Desc
	printerFiles              *prometheus.Desc
	printerPrintTime          *prometheus.Desc
	printerPrintTimeRemaining *prometheus.Desc
	printerPrintProgress      *prometheus.Desc
	printerPrinting           *prometheus.Desc
	printerMaterial           *prometheus.Desc
	printerUp                 *prometheus.Desc
	printerNozzleSize         *prometheus.Desc
	printerStatus             *prometheus.Desc
	printerAxisX              *prometheus.Desc
	printerAxisY              *prometheus.Desc
	printerAxisZ              *prometheus.Desc
	printerFlow               *prometheus.Desc
	printerInfo               *prometheus.Desc
	printerMMU                *prometheus.Desc
	printerFanHotend          *prometheus.Desc
	printerFanPrint           *prometheus.Desc
	printerCover              *prometheus.Desc
	printerFanBlower          *prometheus.Desc
	printerFanRear            *prometheus.Desc
	printerFanUV              *prometheus.Desc
	printerAmbientTemp        *prometheus.Desc
	printerCPUTemp            *prometheus.Desc
	pritnerUVTemp             *prometheus.Desc
	printerBedTempTarget      *prometheus.Desc
	printerBedTempOffset      *prometheus.Desc
	printerChamberTemp        *prometheus.Desc
	printerChamberTempTarget  *prometheus.Desc
	printerChamberTempOffset  *prometheus.Desc
	printerToolTemp           *prometheus.Desc
	printerToolTempTarget     *prometheus.Desc
	printerToolTempOffset     *prometheus.Desc
	printerHeatedChamber      *prometheus.Desc
	printerPrintSpeedRatio    *prometheus.Desc
	printerLogs               *prometheus.Desc
	printerLogsDate           *prometheus.Desc
	printerFarmMode           *prometheus.Desc
	printerCameras            *prometheus.Desc
}

// NewBuddyCollector returns a new buddyCollector
func NewBuddyCollector() *BuddyCollector {
	defaultLabels := []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}
	return &BuddyCollector{
		printerNozzleTemp:         prometheus.NewDesc("prusa_nozzle_temperature", "Current temperature of printer nozzle in Celsius", defaultLabels, nil),
		printerBedTemp:            prometheus.NewDesc("prusa_bed_temperature", "Current temperature of printer bed in Celsius", defaultLabels, nil),
		printerVersion:            prometheus.NewDesc("prusa_version", "Return information about printer. This metric contains information mostly about Prusa Link", append(defaultLabels, "printer_api", "printer_server", "printer_text"), nil),
		printerZHeight:            prometheus.NewDesc("prusa_z_height", "Current height of Z", defaultLabels, nil),
		printerPrintSpeed:         prometheus.NewDesc("prusa_print_speed_ratio", "Current setting of printer speed in ratio (0.0-1.0)", defaultLabels, nil),
		printerTargetTempNozzle:   prometheus.NewDesc("prusa_nozzle_target_temperature", "Target temperature of printer nozzle in Celsius", defaultLabels, nil),
		printerTargetTempBed:      prometheus.NewDesc("prusa_bed_target_temperature", "Target temperature of printer bed in Celsius", defaultLabels, nil),
		printerFiles:              prometheus.NewDesc("prusa_files", "Number of files in storage", append(defaultLabels, "printer_storage"), nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_printing_time_remaining", "Returns time that remains for completion of current print", defaultLabels, nil),
		printerPrintProgress:      prometheus.NewDesc("prusa_printing_progress", "Returns information about completion of current print in percents", defaultLabels, nil),
		printerPrinting:           prometheus.NewDesc("prusa_printing", "Return information about printing", defaultLabels, nil),
		printerMaterial:           prometheus.NewDesc("prusa_material", "Returns information about loaded filament. Returns 0 if there is no loaded filament", append(defaultLabels, "printer_filament"), nil),
		printerPrintTime:          prometheus.NewDesc("prusa_print_time", "Returns information about current print time.", defaultLabels, nil),
		printerUp:                 prometheus.NewDesc("prusa_up", "Return information about online printers. If printer is registered as offline then returned value is 0.", []string{"printer_address", "printer_model", "printer_name"}, nil),
		printerNozzleSize:         prometheus.NewDesc("prusa_nozzle_size", "Returns information about selected nozzle size.", defaultLabels, nil),
		printerStatus:             prometheus.NewDesc("prusa_status", "Returns information status of printer.", append(defaultLabels, "printer_state"), nil),
		printerAxisX:              prometheus.NewDesc("prusa_axis_x", "Returns information about position of axis X.", defaultLabels, nil),
		printerAxisY:              prometheus.NewDesc("prusa_axis_y", "Returns information about position of axis Y.", defaultLabels, nil),
		printerAxisZ:              prometheus.NewDesc("prusa_axis_z", "Returns information about position of axis Z.", defaultLabels, nil),
		printerFlow:               prometheus.NewDesc("prusa_print_flow_ratio", "Returns information about of filament flow in ratio (0.0 - 1.0).", defaultLabels, nil),
		printerInfo:               prometheus.NewDesc("prusa_info", "Returns information about printer.", append(defaultLabels, "printer_serial", "printer_hostname"), nil),
		printerMMU:                prometheus.NewDesc("prusa_mmu", "Returns information if MMU is enabled.", defaultLabels, nil),
		printerFanHotend:          prometheus.NewDesc("prusa_fan_hotend", "Returns information about speed of hotend fan in rpm.", defaultLabels, nil),
		printerFanPrint:           prometheus.NewDesc("prusa_fan_print", "Returns information about speed of print fan in rpm.", defaultLabels, nil),
		printerPrintSpeedRatio:    prometheus.NewDesc("prusa_print_speed_ratio", "Current setting of printer speed in values from 0.0 - 1.0", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerLogs:               prometheus.NewDesc("prusa_logs", "Return size of logs in Prusa Link", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "log_name"}, nil),
		printerLogsDate:           prometheus.NewDesc("prusa_logs_date", "Return date of logs in Prusa Link", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "log_name"}, nil),
		printerFarmMode:           prometheus.NewDesc("prusa_farm_mode", "Return if printer is set to farm mode", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerCameras:            prometheus.NewDesc("prusa_cameras", "Return information about cameras", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "camera_id", "camera_name", "camera_resolution"}, nil),
		printerCover:              prometheus.NewDesc("prusa_cover", "Status of the printer - 0 = open, 1 = closed", defaultLabels, nil),
		printerFanBlower:          prometheus.NewDesc("prusa_fan_blower", "Status of the printer blower fan", defaultLabels, nil),
		printerFanRear:            prometheus.NewDesc("prusa_fan_rear", "Status of the printer fan rear", defaultLabels, nil),
		printerFanUV:              prometheus.NewDesc("prusa_fan_uv", "Status of the printer fan uv", defaultLabels, nil),
		printerAmbientTemp:        prometheus.NewDesc("prusa_ambient_temperature", "Status of the printer ambient temperature", defaultLabels, nil),
		printerCPUTemp:            prometheus.NewDesc("prusa_cpu_temperature", "Status of the printer cpu temperature", defaultLabels, nil),
		pritnerUVTemp:             prometheus.NewDesc("prusa_uv_temperature", "Status of the printer uv temperature", defaultLabels, nil),
		printerBedTempTarget:      prometheus.NewDesc("prusa_bed_target_temperature", "Target bed temperature", defaultLabels, nil),
		printerBedTempOffset:      prometheus.NewDesc("prusa_bed_offset_temperature", "Offset bed temperature", defaultLabels, nil),
		printerChamberTemp:        prometheus.NewDesc("prusa_chamber_temperature", "Status of the printer chamber temperature", defaultLabels, nil),
		printerChamberTempTarget:  prometheus.NewDesc("prusa_chamber_target_temperature", "Traget chamber temperature", defaultLabels, nil),
		printerChamberTempOffset:  prometheus.NewDesc("prusa_chamber_offset_temperature", "Offset chamber temperature", defaultLabels, nil),
		printerToolTemp:           prometheus.NewDesc("prusa_tool_temperature", "Status of the printer tool temperature", defaultLabels, nil),
		printerToolTempTarget:     prometheus.NewDesc("prusa_tool_target_temperature", "Target tool temperature", defaultLabels, nil),
		printerToolTempOffset:     prometheus.NewDesc("prusa_tool_offset_temperature", "Offset tool temperature", defaultLabels, nil),
		printerHeatedChamber:      prometheus.NewDesc("prusa_heated_chamber", "Status of the printer heated chamber", defaultLabels, nil),
	}
}

// Describe implements prometheus.Collector
func (collector *BuddyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printerNozzleTemp
	ch <- collector.printerBedTemp
	ch <- collector.printerVersion
	ch <- collector.printerZHeight
	ch <- collector.printerPrintSpeed
	ch <- collector.printerTargetTempNozzle
	ch <- collector.printerTargetTempBed
	ch <- collector.printerFiles
	ch <- collector.printerPrintTime
	ch <- collector.printerPrintTimeRemaining
	ch <- collector.printerPrintProgress
	ch <- collector.printerPrinting
	ch <- collector.printerMaterial
	ch <- collector.printerUp
	ch <- collector.printerNozzleSize
	ch <- collector.printerStatus
	ch <- collector.printerAxisX
	ch <- collector.printerAxisY
	ch <- collector.printerAxisZ
	ch <- collector.printerFlow
	ch <- collector.printerInfo
	ch <- collector.printerMMU
	ch <- collector.printerFanHotend
	ch <- collector.printerFanPrint
}

// Collect implements prometheus.Collector
func (collector *BuddyCollector) Collect(ch chan<- prometheus.Metric) {
	for _, s := range configuration.Printers {
		log.Debug().Msg("Buddy scraping at " + s.Address)
		if !s.Reachable && (s.Type == "MK4" || s.Type == "MINI" || s.Type == "XL") {
			printerUp := prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
				0, s.Address, s.Type, s.Name)

			ch <- printerUp

			log.Debug().Msg(s.Address + " is unreachable while scraping")
		} else if s.Type == "MK4" || s.Type == "MINI" || s.Type == "XL" {
			version, files, job, printer, status, info, _, err := getBuddyResponse(s)

			if err != nil {
				log.Error().Msg(err.Error())
			} else {
				bedTemp := prometheus.MustNewConstMetric(collector.printerBedTemp, prometheus.GaugeValue,
					printer.Temperature.Bed.Actual, GetLabels(s, job)...)

				nozzleTemp := prometheus.MustNewConstMetric(collector.printerNozzleTemp, prometheus.GaugeValue,
					float64(status.Printer.TempNozzle), GetLabels(s, job)...)

				printProgress := prometheus.MustNewConstMetric(collector.printerPrintProgress, prometheus.GaugeValue,
					float64(job.Progress.Completion), GetLabels(s, job)...)

				printSpeed := prometheus.MustNewConstMetric(collector.printerPrintSpeed, prometheus.GaugeValue,
					float64(printer.Telemetry.PrintSpeed)/100, GetLabels(s, job)...)

				printTimeRemaining := prometheus.MustNewConstMetric(collector.printerPrintTimeRemaining, prometheus.GaugeValue,
					float64(job.Progress.PrintTimeLeft), GetLabels(s, job)...)

				printTime := prometheus.MustNewConstMetric(collector.printerPrintTime, prometheus.GaugeValue,
					float64(job.Progress.PrintTime), GetLabels(s, job)...)

				targetTempBed := prometheus.MustNewConstMetric(collector.printerTargetTempBed, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Target), GetLabels(s, job)...)

				targetTempNozzle := prometheus.MustNewConstMetric(collector.printerTargetTempNozzle, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Target), GetLabels(s, job)...)

				zHeight := prometheus.MustNewConstMetric(collector.printerZHeight, prometheus.GaugeValue,
					printer.Telemetry.ZHeight, GetLabels(s, job)...)

				printing := prometheus.MustNewConstMetric(collector.printerPrinting, prometheus.GaugeValue,
					BoolToFloat(strings.Contains(job.State, "Printing")), GetLabels(s, job)...)

				printerVersion := prometheus.MustNewConstMetric(collector.printerVersion, prometheus.GaugeValue,
					1, GetLabels(s, job, version.API, version.Server, version.Text)...)

				material := prometheus.MustNewConstMetric(collector.printerMaterial, prometheus.GaugeValue,
					BoolToFloat(!strings.Contains(printer.Telemetry.Material, "---")),
					GetLabels(s, job, printer.Telemetry.Material)...)

				if len(files.Files) > 0 {
					printerFiles := prometheus.MustNewConstMetric(collector.printerFiles, prometheus.GaugeValue,
						float64(len(files.Files[0].Children)), GetLabels(s, job, files.Files[0].Display)...)
					ch <- printerFiles
				}

				printerUp := prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
					1, s.Address, s.Type, s.Name)

				printerNozzleSize := prometheus.MustNewConstMetric(collector.printerNozzleSize, prometheus.GaugeValue,
					info.NozzleDiameter, GetLabels(s, job)...)

				printerStatus := prometheus.MustNewConstMetric(collector.printerStatus, prometheus.GaugeValue,
					getStateFlag(printer), GetLabels(s, job, status.Printer.State)...)

				printerAxisX := prometheus.MustNewConstMetric(collector.printerAxisX, prometheus.GaugeValue,
					float64(status.Printer.AxisX), GetLabels(s, job)...)

				printerAxisY := prometheus.MustNewConstMetric(collector.printerAxisY, prometheus.GaugeValue,
					float64(status.Printer.AxisY), GetLabels(s, job)...)

				printerAxisZ := prometheus.MustNewConstMetric(collector.printerAxisZ, prometheus.GaugeValue,
					float64(status.Printer.AxisZ), GetLabels(s, job)...)

				printerFlow := prometheus.MustNewConstMetric(collector.printerFlow, prometheus.GaugeValue,
					float64(status.Printer.Flow)/100, GetLabels(s, job)...)

				printerInfo := prometheus.MustNewConstMetric(collector.printerInfo, prometheus.GaugeValue,
					1, GetLabels(s, job, info.Serial, info.Hostname)...)

				printerMMU := prometheus.MustNewConstMetric(collector.printerMMU, prometheus.GaugeValue,
					BoolToFloat(info.Mmu), GetLabels(s, job)...)

				printerFanHotend := prometheus.MustNewConstMetric(collector.printerFanHotend, prometheus.GaugeValue,
					float64(status.Printer.FanHotend), GetLabels(s, job)...)

				printerFanPrint := prometheus.MustNewConstMetric(collector.printerFanPrint, prometheus.GaugeValue,
					float64(status.Printer.FanPrint), GetLabels(s, job)...)

				ch <- printerStatus
				ch <- bedTemp
				ch <- nozzleTemp
				ch <- printProgress
				ch <- printSpeed
				ch <- printTimeRemaining
				ch <- printing
				ch <- printTime
				ch <- targetTempBed
				ch <- targetTempNozzle
				ch <- material
				ch <- printerVersion
				ch <- zHeight
				ch <- printerUp
				ch <- printerNozzleSize
				ch <- printerAxisX
				ch <- printerAxisY
				ch <- printerAxisZ
				ch <- printerFlow
				ch <- printerInfo
				ch <- printerMMU
				ch <- printerFanHotend
				ch <- printerFanPrint
			}
		}
	}
}
