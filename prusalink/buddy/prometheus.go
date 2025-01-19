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
	printerTemp               *prometheus.Desc
	printerTempTarget         *prometheus.Desc
	printerPrintSpeed         *prometheus.Desc
	printerPrintTime          *prometheus.Desc
	printerPrintTimeRemaining *prometheus.Desc
	printerPrintProgressRatio *prometheus.Desc
	printerFiles              *prometheus.Desc
	printerMaterial           *prometheus.Desc
	printerUp                 *prometheus.Desc
	printerNozzleSize         *prometheus.Desc
	printerStatus             *prometheus.Desc
	printerAxis               *prometheus.Desc
	printerFlow               *prometheus.Desc
	printerInfo               *prometheus.Desc
	printerMMU                *prometheus.Desc
	printerFanSpeedRpm        *prometheus.Desc
	printerPrintSpeedRatio    *prometheus.Desc
}

// NewCollector returns a new Collector for printer metrics
func NewCollector(config config.Config) *Collector {
	configuration = config
	defaultLabels := []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}
	return &Collector{
		printerTemp:               prometheus.NewDesc("prusa_temperature_celsius", "Current temp of printer in Celsius", append(defaultLabels, "printer_heated_element"), nil),
		printerTempTarget:         prometheus.NewDesc("prusa_temperature_target_celsius", "Target temp of printer in Celsius", append(defaultLabels, "printer_heated_element"), nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_printing_time_remaining_seconds", "Returns time that remains for completion of current print", defaultLabels, nil),
		printerPrintProgressRatio: prometheus.NewDesc("prusa_printing_progress_ratio", "Returns information about completion of current print in ratio (0.0-1.0)", defaultLabels, nil),
		printerFiles:              prometheus.NewDesc("prusa_files_count", "Number of files in storage", append(defaultLabels, "printer_storage"), nil),
		printerMaterial:           prometheus.NewDesc("prusa_material_info", "Returns information about loaded filament. Returns 0 if there is no loaded filament", append(defaultLabels, "printer_filament"), nil),
		printerPrintTime:          prometheus.NewDesc("prusa_print_time_seconds", "Returns information about current print time.", defaultLabels, nil),
		printerUp:                 prometheus.NewDesc("prusa_up", "Return information about online printers. If printer is registered as offline then returned value is 0.", []string{"printer_address", "printer_model", "printer_name"}, nil),
		printerNozzleSize:         prometheus.NewDesc("prusa_nozzle_size_meters", "Returns information about selected nozzle size.", defaultLabels, nil),
		printerStatus:             prometheus.NewDesc("prusa_status_info", "Returns information status of printer.", append(defaultLabels, "printer_state"), nil),
		printerAxis:               prometheus.NewDesc("prusa_axis", "Returns information about position of axis.", append(defaultLabels, "printer_axis"), nil),
		printerFlow:               prometheus.NewDesc("prusa_print_flow_ratio", "Returns information about of filament flow in ratio (0.0 - 1.0).", defaultLabels, nil),
		printerInfo:               prometheus.NewDesc("prusa_info", "Returns information about printer.", append(defaultLabels, "api_version", "server_version", "version_text", "prusalink_name", "printer_location", "serial_number", "printer_hostname"), nil),
		printerMMU:                prometheus.NewDesc("prusa_mmu", "Returns information if MMU is enabled.", defaultLabels, nil),
		printerFanSpeedRpm:        prometheus.NewDesc("prusa_fan_speed_rpm", "Returns information about speed of hotend fan in rpm.", append(defaultLabels, "fan"), nil),
		printerPrintSpeedRatio:    prometheus.NewDesc("prusa_print_speed_ratio", "Current setting of printer speed in values from 0.0 - 1.0", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}, nil)}
}

// Describe implements prometheus.Collector
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printerTemp

	ch <- collector.printerFiles
	ch <- collector.printerPrintTime
	ch <- collector.printerPrintTimeRemaining
	ch <- collector.printerPrintProgressRatio
	ch <- collector.printerPrintSpeedRatio
	ch <- collector.printerMaterial
	ch <- collector.printerUp
	ch <- collector.printerNozzleSize
	ch <- collector.printerStatus
	ch <- collector.printerAxis
	ch <- collector.printerFlow
	ch <- collector.printerInfo
	ch <- collector.printerMMU
	ch <- collector.printerFanSpeedRpm
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

			version, err := GetVersion(s)
			if err != nil {
				log.Error().Msg("Error while scraping version endpoint at " + s.Address + " - " + err.Error())
				ch <- printerUp
				return
			}

			status, err := GetStatus(s)

			if err != nil {
				log.Error().Msg("Error while scraping status endpoint at " + s.Address + " - " + err.Error())
			}

			info, err := GetInfo(s)

			if err != nil {
				log.Error().Msg("Error while scraping info endpoint at " + s.Address + " - " + err.Error())
			}

			printerInfo := prometheus.MustNewConstMetric(
				collector.printerInfo, prometheus.GaugeValue,
				1,
				GetLabels(s, job, version.API, version.Server, version.Text, info.Name, info.Location, info.Serial, info.Hostname)...)

			ch <- printerInfo

			printerFanHotend := prometheus.MustNewConstMetric(collector.printerFanSpeedRpm, prometheus.GaugeValue,
				status.Printer.FanHotend, GetLabels(s, job, "hotend")...)

			ch <- printerFanHotend

			printerFanPrint := prometheus.MustNewConstMetric(collector.printerFanSpeedRpm, prometheus.GaugeValue,
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
				collector.printerPrintProgressRatio, prometheus.GaugeValue,
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

			printerFlow := prometheus.MustNewConstMetric(collector.printerFlow, prometheus.GaugeValue,
				status.Printer.Flow/100, GetLabels(s, job)...)

			ch <- printerFlow

			printerMMU := prometheus.MustNewConstMetric(collector.printerMMU, prometheus.GaugeValue,
				BoolToFloat(info.Mmu), GetLabels(s, job)...)
			ch <- printerMMU

			printerBedTemp := prometheus.MustNewConstMetric(collector.printerTemp, prometheus.GaugeValue,
				printer.Temperature.Bed.Actual, GetLabels(s, job, "bed")...)

			ch <- printerBedTemp

			printerBedTempTarget := prometheus.MustNewConstMetric(collector.printerTempTarget, prometheus.GaugeValue,
				printer.Temperature.Bed.Target, GetLabels(s, job, "bed")...)

			ch <- printerBedTempTarget

			printerToolTempTarget := prometheus.MustNewConstMetric(collector.printerTempTarget, prometheus.GaugeValue,
				printer.Temperature.Tool0.Target, GetLabels(s, job, "tool0")...)

			ch <- printerToolTempTarget

			printerToolTemp := prometheus.MustNewConstMetric(collector.printerTemp, prometheus.GaugeValue,
				printer.Temperature.Tool0.Actual, GetLabels(s, job, "tool0")...)

			ch <- printerToolTemp

			printerStatus := prometheus.MustNewConstMetric(
				collector.printerStatus, prometheus.GaugeValue,
				getStateFlag(printer),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.State.Text)

			ch <- printerStatus

			printerUp = prometheus.MustNewConstMetric(collector.printerUp, prometheus.GaugeValue,
				1, s.Address, s.Type, s.Name)

			ch <- printerUp

			log.Debug().Msg("Scraping done at " + s.Address)
		}(s)
	}
	wg.Wait()
}
