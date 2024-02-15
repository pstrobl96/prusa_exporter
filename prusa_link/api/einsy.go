package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

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

func getFlagEinsy(printer einsyPrinter) float64 {
	if printer.State.Flags.Operational {
		return 1
	} else if printer.State.Flags.Prepared {
		return 2
	} else if printer.State.Flags.Paused {
		return 3
	} else if printer.State.Flags.Printing {
		return 4
	} else if printer.State.Flags.Cancelling {
		return 5
	} else if printer.State.Flags.Pausing {
		return 6
	} else if printer.State.Flags.Error {
		return 7
	} else if printer.State.Flags.SdReady {
		return 8
	} else if printer.State.Flags.ClosedOrError {
		return 9
	} else if printer.State.Flags.Ready {
		return 10
		/*} else if printer.State.Flags.Busy {
		return 11*/ // no busy state for Einsy
	} else if printer.State.Flags.Finished {
		return 12
	} else {
		return 0
	}
}

func (collector *einsyCollector) Collect(ch chan<- prometheus.Metric) {
	cfg := &config
	for _, s := range cfg.Printers.Einsy {
		log.Debug().Msg("Einsy scraping at " + s.Address)
		if !s.Reachable {
			printerUp := prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
				0, s.Address, s.Type, s.Name)

			ch <- printerUp

			log.Debug().Msg(s.Address + " is unreachable while scraping")
		} else {
			version, files, job, printer, cameras, info, settings, ports, e := getEinsyResponse(s)

			if e != nil {
				printerUp := prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
					0, s.Address, s.Type, s.Name)

				ch <- printerUp

				log.Debug().Msg(s.Address + " is unreachable while scraping")
				log.Error().Msg(e.Error())
				break
			}
			nozzleTemp := prometheus.MustNewConstMetric(
				collector.printerNozzleTemp, prometheus.GaugeValue,
				printer.Temperature.Tool0.Actual,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			bedTemp := prometheus.MustNewConstMetric(
				collector.printerBedTemp, prometheus.GaugeValue, // collector
				printer.Telemetry.TempBed,                                       // value
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path) // labels

			printSpeed := prometheus.MustNewConstMetric(
				collector.printerPrintSpeedRatio, prometheus.GaugeValue,
				float64(printer.Telemetry.PrintSpeed)/100,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			targetTempBed := prometheus.MustNewConstMetric(
				collector.printerTargetTempBed, prometheus.GaugeValue,
				printer.Temperature.Bed.Target,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			targetTempNozzle := prometheus.MustNewConstMetric(
				collector.printerTargetTempNozzle, prometheus.GaugeValue,
				printer.Temperature.Tool0.Target,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			for _, v := range files.Files {
				printerFiles := prometheus.MustNewConstMetric(
					collector.printerFiles, prometheus.GaugeValue,
					float64(len(v.Children)),
					s.Address, s.Type, s.Name, v.Display)
				ch <- printerFiles
			}

			printTime := prometheus.MustNewConstMetric(
				collector.printerPrintTime, prometheus.GaugeValue,
				float64(job.Progress.PrintTime),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printTimeRemaining := prometheus.MustNewConstMetric(
				collector.printerPrintTimeRemaining, prometheus.GaugeValue,
				float64(job.Progress.PrintTimeLeft),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printProgress := prometheus.MustNewConstMetric(
				collector.printerPrintProgress, prometheus.GaugeValue,
				job.Progress.Completion,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			filamentLoaded := 0 // does not work with MMU
			if printer.Telemetry.Material != " - " {
				filamentLoaded = 1
			}

			material := prometheus.MustNewConstMetric(
				collector.printerMaterial, prometheus.GaugeValue,
				float64(filamentLoaded),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.Telemetry.Material)

			printerInfo := prometheus.MustNewConstMetric(
				collector.printerInfo, prometheus.GaugeValue,
				1,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, version.API, version.Server, version.Text, info.Name, info.Location, info.Serial, info.Hostname, ports.Ports[0].Description)
			ch <- printerInfo

			farmMode := 0
			if settings.Printer.FarmMode {
				farmMode = 1
			}
			printerFarmMode := prometheus.MustNewConstMetric(
				collector.printerFarmMode, prometheus.GaugeValue,
				float64(farmMode),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			for _, v := range cameras.CameraList {
				cameraConnected := 0
				if v.Connected {
					cameraConnected = 1
				}
				printerCamera := prometheus.MustNewConstMetric(
					collector.printerCameras, prometheus.GaugeValue,
					float64(cameraConnected),
					s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, v.CameraID, v.Config.Name, v.Config.Resolution)
				ch <- printerCamera
			}

			printerAxisX := prometheus.MustNewConstMetric(
				collector.printerAxisX, prometheus.GaugeValue,
				printer.Telemetry.AxisX,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printerAxisY := prometheus.MustNewConstMetric(
				collector.printerAxisY, prometheus.GaugeValue,
				printer.Telemetry.AxisY,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printerAxisZ := prometheus.MustNewConstMetric(
				collector.printerAxisZ, prometheus.GaugeValue,
				printer.Telemetry.AxisZ,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printerStatus := prometheus.MustNewConstMetric(
				collector.printerStatus, prometheus.GaugeValue,
				getFlagEinsy(printer),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.State.Text)

			printerNozzleSize := prometheus.MustNewConstMetric(
				collector.printerNozzleSize, prometheus.GaugeValue,
				info.NozzleDiameter,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printerUp := prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
				1, s.Address, s.Type, s.Name)

			ch <- printerUp
			ch <- printerStatus
			ch <- printerNozzleSize
			ch <- printerAxisX
			ch <- printerAxisY
			ch <- printerAxisZ
			ch <- bedTemp
			ch <- nozzleTemp
			ch <- printProgress
			ch <- printSpeed
			ch <- printTimeRemaining
			ch <- printTime
			ch <- targetTempBed
			ch <- targetTempNozzle
			ch <- material
			ch <- printerFarmMode
		}
	}
}