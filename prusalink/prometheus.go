package prusalink

import (
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/rs/zerolog/log"
)

// Collector is a struct of all printer metrics
type Collector struct {
	printerBedTemp            *prometheus.Desc
	printerPrintSpeed         *prometheus.Desc
	printerFiles              *prometheus.Desc
	printerPrintTime          *prometheus.Desc
	printerPrintTimeRemaining *prometheus.Desc
	printerPrintProgress      *prometheus.Desc
	printerMaterial           *prometheus.Desc
	printerUp                 *prometheus.Desc
	printerNozzleSize         *prometheus.Desc
	printerStatus             *prometheus.Desc
	printerAxis               *prometheus.Desc
	printerZHeight            *prometheus.Desc
	printerFlow               *prometheus.Desc
	printerInfo               *prometheus.Desc
	printerMMU                *prometheus.Desc
	printerCover              *prometheus.Desc
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
	printerFanSpeed           *prometheus.Desc
}

// NewCollector returns a new Collector for printer metrics
func NewCollector(config config.Config) *Collector {
	configuration = config
	defaultLabels := []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}
	return &Collector{
		printerBedTemp:            prometheus.NewDesc("prusa_bed_temp", "Current temp of printer bed in Celsius", defaultLabels, nil),
		printerPrintSpeed:         prometheus.NewDesc("prusa_print_speed_ratio", "Current setting of printer speed in ratio (0.0-1.0)", defaultLabels, nil),
		printerFiles:              prometheus.NewDesc("prusa_files", "Number of files in storage", append(defaultLabels, "printer_storage"), nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_printing_time_remaining", "Returns time that remains for completion of current print", defaultLabels, nil),
		printerPrintProgress:      prometheus.NewDesc("prusa_printing_progress", "Returns information about completion of current print in percents", defaultLabels, nil),
		printerMaterial:           prometheus.NewDesc("prusa_material", "Returns information about loaded filament. Returns 0 if there is no loaded filament", append(defaultLabels, "printer_filament"), nil),
		printerPrintTime:          prometheus.NewDesc("prusa_print_time", "Returns information about current print time.", defaultLabels, nil),
		printerUp:                 prometheus.NewDesc("prusa_up", "Return information about online printers. If printer is registered as offline then returned value is 0.", []string{"printer_address", "printer_model", "printer_name"}, nil),
		printerNozzleSize:         prometheus.NewDesc("prusa_nozzle_size", "Returns information about selected nozzle size.", defaultLabels, nil),
		printerStatus:             prometheus.NewDesc("prusa_status", "Returns information status of printer.", append(defaultLabels, "printer_state"), nil),
		printerAxis:               prometheus.NewDesc("prusa_axis", "Returns information about position of axis.", append(defaultLabels, "printer_axis"), nil),
		printerZHeight:            prometheus.NewDesc("prusa_z_height_meters", "Returns information about current height of Z axis.", defaultLabels, nil),
		printerFlow:               prometheus.NewDesc("prusa_print_flow_ratio", "Returns information about of filament flow in ratio (0.0 - 1.0).", defaultLabels, nil),
		printerInfo:               prometheus.NewDesc("prusa_info", "Returns information about printer.", append(defaultLabels, "api_version", "server_version", "version_text", "prusalink_name", "printer_location", "serial_number", "printer_hostname"), nil),
		printerMMU:                prometheus.NewDesc("prusa_mmu", "Returns information if MMU is enabled.", defaultLabels, nil),
		printerFanSpeed:           prometheus.NewDesc("prusa_fan_speed", "Returns information about speed of hotend fan in rpm.", append(defaultLabels, "fan"), nil),
		printerPrintSpeedRatio:    prometheus.NewDesc("prusa_print_speed_ratio", "Current setting of printer speed in values from 0.0 - 1.0", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerLogs:               prometheus.NewDesc("prusa_logs", "Return size of logs in Prusa Link", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "log_name"}, nil),
		printerLogsDate:           prometheus.NewDesc("prusa_logs_date", "Return date of logs in Prusa Link", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "log_name"}, nil),
		printerFarmMode:           prometheus.NewDesc("prusa_farm_mode", "Return if printer is set to farm mode", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil),
		printerCameras:            prometheus.NewDesc("prusa_cameras", "Return information about cameras", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "camera_id", "camera_name", "camera_resolution"}, nil),
		printerCover:              prometheus.NewDesc("prusa_cover", "Status of the printer - 0 = open, 1 = closed", defaultLabels, nil),
		printerAmbientTemp:        prometheus.NewDesc("prusa_ambient_temp", "Status of the printer ambient temp", defaultLabels, nil),
		printerCPUTemp:            prometheus.NewDesc("prusa_cpu_temp", "Status of the printer cpu temp", defaultLabels, nil),
		pritnerUVTemp:             prometheus.NewDesc("prusa_uv_temp", "Status of the printer uv temp", defaultLabels, nil),
		printerBedTempTarget:      prometheus.NewDesc("prusa_bed_temp_target", "Target bed temp", defaultLabels, nil),
		printerBedTempOffset:      prometheus.NewDesc("prusa_bed_temp_offset", "Offset bed temp", defaultLabels, nil),
		printerChamberTemp:        prometheus.NewDesc("prusa_chamber_temp", "Status of the printer chamber temp", defaultLabels, nil),
		printerChamberTempTarget:  prometheus.NewDesc("prusa_chamber_temp_target", "Traget chamber temp", defaultLabels, nil),
		printerChamberTempOffset:  prometheus.NewDesc("prusa_chamber_temp_offset", "Offset chamber temp", defaultLabels, nil),
		printerToolTemp:           prometheus.NewDesc("prusa_tool_temp", "Status of the printer tool temp", append(defaultLabels, "tool"), nil),
		printerToolTempTarget:     prometheus.NewDesc("prusa_tool_temp_target", "Target tool temp", append(defaultLabels, "tool"), nil),
		printerToolTempOffset:     prometheus.NewDesc("prusa_tool_temp_offset", "Offset tool temp", append(defaultLabels, "tool"), nil),
		printerHeatedChamber:      prometheus.NewDesc("prusa_heated_chamber", "Status of the printer heated chamber", defaultLabels, nil),
	}
}

// Describe implements prometheus.Collector
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printerBedTemp
	ch <- collector.printerPrintSpeed
	ch <- collector.printerFiles
	ch <- collector.printerPrintTime
	ch <- collector.printerPrintTimeRemaining
	ch <- collector.printerPrintProgress
	ch <- collector.printerMaterial
	ch <- collector.printerUp
	ch <- collector.printerNozzleSize
	ch <- collector.printerStatus
	ch <- collector.printerAxis
	ch <- collector.printerZHeight
	ch <- collector.printerFlow
	ch <- collector.printerInfo
	ch <- collector.printerMMU
	ch <- collector.printerCover
	ch <- collector.printerAmbientTemp
	ch <- collector.printerCPUTemp
	ch <- collector.pritnerUVTemp
	ch <- collector.printerChamberTemp
	ch <- collector.printerToolTemp
	ch <- collector.printerHeatedChamber
	ch <- collector.printerBedTempTarget
	ch <- collector.printerBedTempOffset
	ch <- collector.printerChamberTempTarget
	ch <- collector.printerChamberTempOffset
	ch <- collector.printerToolTempTarget
	ch <- collector.printerToolTempOffset
	ch <- collector.printerCameras
	ch <- collector.printerFarmMode
	ch <- collector.printerLogsDate
	ch <- collector.printerLogs
	ch <- collector.printerFanSpeed
}

// Collect implements prometheus.Collector
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {

	var wg sync.WaitGroup
	for _, s := range configuration.Printers {
		wg.Add(1)
		go func(s config.Printers) {
			defer wg.Done()

			log.Debug().Msg("Printer scraping at " + s.Address)
			printerUp := prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
				0, s.Address, s.Type, s.Name)

			if s.Type == "" {
				printerType, err := GetPrinterType(s)
				if err != nil {
					log.Error().Msg("Error while probing printer at " + s.Address + " - " + err.Error())
					ch <- printerUp
					return
				}
				s.Type = printerType
			}

			job, err := GetJob(s)
			if err != nil {
				log.Error().Msg("Error while scraping job endpoint at " + s.Address + " - " + err.Error())
				ch <- printerUp
				return
			}

			printer, err := GetPrinter(s)
			if err != nil {
				log.Error().Msg("Error while scraping printer endpoint at " + s.Address + " - " + err.Error())
				ch <- printerUp
				return
			}

			files, err := GetFiles(s)
			if err != nil {
				log.Error().Msg("Error while scraping files endpoint at " + s.Address + " - " + err.Error())
				ch <- printerUp
				return
			}

			version, err := GetVersion(s)
			if err != nil {
				log.Error().Msg("Error while scraping version endpoint at " + s.Address + " - " + err.Error())
				ch <- printerUp
				return
			}

			// metrics specific for both buddy and einsy
			if printerBoards[s.Type] == "buddy" || printerBoards[s.Type] == "einsy" {

				status, err := GetStatus(s)

				if err != nil {
					log.Error().Msg("Error while scraping status endpoint at " + s.Address + " - " + err.Error())
				}

				info, err := GetInfo(s)

				if err != nil {
					log.Error().Msg("Error while scraping info endpoint at " + s.Address + " - " + err.Error())
				}

				// only einsy related metrics
				if printerBoards[s.Type] == "einsy" {
					settings, err := GetSettings(s)

					if err != nil {
						log.Error().Msg("Error while scraping settings endpoint at " + s.Address + " - " + err.Error())
					} else {

						printerFarmMode := prometheus.MustNewConstMetric(
							collector.printerFarmMode, prometheus.GaugeValue,
							BoolToFloat(settings.Printer.FarmMode),
							GetLabels(s, job)...)

						ch <- printerFarmMode

					}

					cameras, err := GetCameras(s)

					if err != nil {
						log.Error().Msg("Error while scraping cameras endpoint at " + s.Address + " - " + err.Error())
					} else {

						for _, v := range cameras.CameraList {
							printerCamera := prometheus.MustNewConstMetric(
								collector.printerCameras, prometheus.GaugeValue,
								BoolToFloat(v.Connected),
								GetLabels(s, job, v.CameraID, v.Config.Name, v.Config.Resolution)...)
							ch <- printerCamera
						}
					}

					for _, v := range files.Files {
						printerFiles := prometheus.MustNewConstMetric(
							collector.printerFiles, prometheus.GaugeValue,
							float64(len(v.Children)),
							GetLabels(s, job, v.Display)...)
						ch <- printerFiles
					}

				}

				printerInfo := prometheus.MustNewConstMetric(
					collector.printerInfo, prometheus.GaugeValue,
					1,
					GetLabels(s, job, version.API, version.Server, version.Text, info.Name, info.Location, info.Serial, info.Hostname)...)

				ch <- printerInfo

				printerFanHotend := prometheus.MustNewConstMetric(collector.printerFanSpeed, prometheus.GaugeValue,
					status.Printer.FanHotend, GetLabels(s, job, "hotend")...)

				ch <- printerFanHotend

				printerFanPrint := prometheus.MustNewConstMetric(collector.printerFanSpeed, prometheus.GaugeValue,
					status.Printer.FanPrint, GetLabels(s, job, "print")...)

				ch <- printerFanPrint

				printerNozzleSize := prometheus.MustNewConstMetric(collector.printerNozzleSize, prometheus.GaugeValue,
					info.NozzleDiameter, GetLabels(s, job)...)

				ch <- printerNozzleSize

				printSpeed := prometheus.MustNewConstMetric(
					collector.printerPrintSpeedRatio, prometheus.GaugeValue,
					printer.Telemetry.PrintSpeed/100,
					s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

				ch <- printSpeed

				printTime := prometheus.MustNewConstMetric(
					collector.printerPrintTime, prometheus.GaugeValue,
					job.Progress.PrintTime,
					s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

				ch <- printTime

				printTimeRemaining := prometheus.MustNewConstMetric(
					collector.printerPrintTimeRemaining, prometheus.GaugeValue,
					job.Progress.PrintTimeLeft,
					s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

				ch <- printTimeRemaining

				printProgress := prometheus.MustNewConstMetric(
					collector.printerPrintProgress, prometheus.GaugeValue,
					job.Progress.Completion,
					s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

				ch <- printProgress

				material := prometheus.MustNewConstMetric(
					collector.printerMaterial, prometheus.GaugeValue,
					BoolToFloat(!(strings.Contains(printer.Telemetry.Material, "-"))),
					s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.Telemetry.Material)

				ch <- material

				printerAxisX := prometheus.MustNewConstMetric(
					collector.printerAxis, prometheus.GaugeValue,
					printer.Telemetry.AxisX,
					GetLabels(s, job, "x")...)

				ch <- printerAxisX

				printerAxisY := prometheus.MustNewConstMetric(
					collector.printerAxis, prometheus.GaugeValue,
					printer.Telemetry.AxisY,
					GetLabels(s, job, "y")...)

				ch <- printerAxisY

				printerAxisZ := prometheus.MustNewConstMetric(
					collector.printerAxis, prometheus.GaugeValue,
					printer.Telemetry.AxisZ,
					GetLabels(s, job, "z")...)

				ch <- printerAxisZ

				printerZHeight := prometheus.MustNewConstMetric(
					collector.printerZHeight, prometheus.GaugeValue,
					printer.Telemetry.ZHeight/1000,
					GetLabels(s, job)...)

				ch <- printerZHeight

				printerFlow := prometheus.MustNewConstMetric(collector.printerFlow, prometheus.GaugeValue,
					status.Printer.Flow/100, GetLabels(s, job)...)

				ch <- printerFlow

				if printerBoards[s.Type] == "buddy" {
					printerMMU := prometheus.MustNewConstMetric(collector.printerMMU, prometheus.GaugeValue,
						BoolToFloat(info.Mmu), GetLabels(s, job)...)
					ch <- printerMMU
				}
			}

			// only sl related metrics
			if printerBoards[s.Type] == "sl" {
				printerCover := prometheus.MustNewConstMetric(collector.printerCover, prometheus.GaugeValue,
					BoolToFloat(printer.Telemetry.CoverClosed), GetLabels(s, job)...)

				ch <- printerCover

				printerFanBlower := prometheus.MustNewConstMetric(collector.printerFanSpeed, prometheus.GaugeValue,
					printer.Telemetry.FanBlower, GetLabels(s, job, "blower")...)

				ch <- printerFanBlower

				printerFanRear := prometheus.MustNewConstMetric(collector.printerFanSpeed, prometheus.GaugeValue,
					printer.Telemetry.FanRear, GetLabels(s, job, "rear")...)

				ch <- printerFanRear

				printerFanUV := prometheus.MustNewConstMetric(collector.printerFanSpeed, prometheus.GaugeValue,
					printer.Telemetry.FanUvLed, GetLabels(s, job, "uv")...)

				ch <- printerFanUV

				printerAmbientTemp := prometheus.MustNewConstMetric(collector.printerAmbientTemp, prometheus.GaugeValue,
					printer.Telemetry.TempAmbient, GetLabels(s, job)...)

				ch <- printerAmbientTemp

				printerCPUTemp := prometheus.MustNewConstMetric(collector.printerCPUTemp, prometheus.GaugeValue,
					printer.Telemetry.TempCPU, GetLabels(s, job)...)

				ch <- printerCPUTemp

				pritnerUVTemp := prometheus.MustNewConstMetric(collector.pritnerUVTemp, prometheus.GaugeValue,
					printer.Telemetry.TempUvLed, GetLabels(s, job)...)

				ch <- pritnerUVTemp

				printerChamberTempTarget := prometheus.MustNewConstMetric(collector.printerChamberTempTarget, prometheus.GaugeValue,
					printer.Temperature.Chamber.Target, GetLabels(s, job)...)

				ch <- printerChamberTempTarget

				printerChamberTempOffset := prometheus.MustNewConstMetric(collector.printerChamberTempOffset, prometheus.GaugeValue,
					printer.Temperature.Chamber.Offset, GetLabels(s, job)...)

				ch <- printerChamberTempOffset

				printerChamberTemp := prometheus.MustNewConstMetric(collector.printerChamberTemp, prometheus.GaugeValue,
					printer.Temperature.Chamber.Actual, GetLabels(s, job)...)

				ch <- printerChamberTemp
			}

			printerBedTemp := prometheus.MustNewConstMetric(collector.printerBedTemp, prometheus.GaugeValue,
				printer.Temperature.Bed.Actual, GetLabels(s, job)...)

			ch <- printerBedTemp

			printerBedTempTarget := prometheus.MustNewConstMetric(collector.printerBedTempTarget, prometheus.GaugeValue,
				printer.Temperature.Bed.Target, GetLabels(s, job)...)

			ch <- printerBedTempTarget

			printerBedTempOffset := prometheus.MustNewConstMetric(collector.printerBedTempOffset, prometheus.GaugeValue,
				printer.Temperature.Bed.Offset, GetLabels(s, job)...)

			ch <- printerBedTempOffset

			printerStatus := prometheus.MustNewConstMetric(
				collector.printerStatus, prometheus.GaugeValue,
				getStateFlag(printer),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.State.Text)

			ch <- printerStatus

			printerToolTempTarget := prometheus.MustNewConstMetric(collector.printerToolTempTarget, prometheus.GaugeValue,
				printer.Temperature.Tool0.Target, GetLabels(s, job, "0")...)

			ch <- printerToolTempTarget

			printerToolTempOffset := prometheus.MustNewConstMetric(collector.printerToolTempOffset, prometheus.GaugeValue,
				printer.Temperature.Tool0.Offset, GetLabels(s, job, "0")...)

			ch <- printerToolTempOffset

			printerToolTemp := prometheus.MustNewConstMetric(collector.printerToolTemp, prometheus.GaugeValue,
				printer.Temperature.Tool0.Actual, GetLabels(s, job, "0")...)

			ch <- printerToolTemp

			printerUp = prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
				1, s.Address, s.Type, s.Name)

			ch <- printerUp

			log.Debug().Msg("Scraping done at " + s.Address)
		}(s)
	}
	wg.Wait()
}
