package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type buddyCollector struct {
	printerNozzleTemp         *prometheus.Desc
	printerBedTemp            *prometheus.Desc
	printerVersion            *prometheus.Desc // info from version
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
}

func newBuddyCollector() *buddyCollector {
	return &buddyCollector{
		printerNozzleTemp: prometheus.NewDesc("prusa_buddy_nozzle_temperature",
			"Current temperature of printer nozzle in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerBedTemp: prometheus.NewDesc("prusa_buddy_bed_temperature",
			"Current temperature of printer bed in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerVersion: prometheus.NewDesc("prusa_buddy_version",
			"Return information about printer. This metric contains information mostly about Prusa Link",
			[]string{"printer_address", "printer_model", "printer_name", "printer_api", "printer_server", "printer_text"},
			nil),
		printerZHeight: prometheus.NewDesc("prusa_buddy_z_height",
			"Current height of Z",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrintSpeed: prometheus.NewDesc("prusa_buddy_print_speed",
			"Current setting of printer speed in percents (%)",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerTargetTempNozzle: prometheus.NewDesc("prusa_buddy_nozzle_target_temperature",
			"Target temperature of printer nozzle in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerTargetTempBed: prometheus.NewDesc("prusa_buddy_bed_target_temperature",
			"Target temperature of printer bed in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerFiles: prometheus.NewDesc("prusa_buddy_files",
			"Number of files in storage",
			[]string{"printer_address", "printer_model", "printer_name", "printer_storage"},
			nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_buddy_printing_time_remaining",
			"Returns time that remains for completion of current print",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrintProgress: prometheus.NewDesc("prusa_buddy_printing_progress",
			"Returns information about completion of current print in percents",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrinting: prometheus.NewDesc("prusa_buddy_printing",
			"Return information about printing",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerMaterial: prometheus.NewDesc("prusa_buddy_material",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "printer_filament"},
			nil),
		printerPrintTime: prometheus.NewDesc("prusa_buddy_print_time",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
	}
}

func (collector *buddyCollector) Describe(ch chan<- *prometheus.Desc) {
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

}

func (collector *buddyCollector) Collect(ch chan<- prometheus.Metric) {
	cfg := loadedConfig

	for _, s := range cfg.Printers.Buddy {
		logger.Debug("Buddy scraping at " + s.Address)
		if head("http://" + s.Address) {
			printer := getBuddyPrinter(s.Address, s.Apikey, s.Username, s.Pass)
			files := getBuddyFiles(s.Address, s.Apikey, s.Username, s.Pass)
			version := getBuddyVersion(s.Address, s.Apikey, s.Username, s.Pass)
			job := getBuddyJob(s.Address, s.Apikey, s.Username, s.Pass)
			bedTemp := prometheus.MustNewConstMetric(
				collector.printerBedTemp, prometheus.GaugeValue,
				float64(printer.Temperature.Bed.Actual),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			nozzleTemp := prometheus.MustNewConstMetric(
				collector.printerNozzleTemp, prometheus.GaugeValue,
				float64(printer.Temperature.Tool0.Actual),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printProgress := prometheus.MustNewConstMetric(
				collector.printerPrintProgress, prometheus.GaugeValue,
				float64(job.Progress.Completion),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printSpeed := prometheus.MustNewConstMetric(
				collector.printerPrintSpeed, prometheus.GaugeValue,
				float64(printer.Telemetry.PrintSpeed),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printTimeRemaining := prometheus.MustNewConstMetric(
				collector.printerPrintTimeRemaining, prometheus.GaugeValue,
				float64(job.Progress.PrintTimeLeft),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printingMetric := 0
			if job.State == "Printing" {
				printingMetric = 1
			}

			printing := prometheus.MustNewConstMetric(
				collector.printerPrinting, prometheus.GaugeValue,
				float64(printingMetric),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			printTime := prometheus.MustNewConstMetric(
				collector.printerPrintTime, prometheus.GaugeValue,
				float64(job.Progress.PrintTime),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			targetTempBed := prometheus.MustNewConstMetric(
				collector.printerTargetTempBed, prometheus.GaugeValue,
				float64(printer.Temperature.Bed.Target),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			targetTempNozzle := prometheus.MustNewConstMetric(
				collector.printerTargetTempNozzle, prometheus.GaugeValue,
				float64(printer.Temperature.Tool0.Target),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

			filamentLoaded := 0
			if printer.Telemetry.Material != "---" {
				filamentLoaded = 1
			}

			material := prometheus.MustNewConstMetric(
				collector.printerMaterial, prometheus.GaugeValue,
				float64(filamentLoaded),
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.Telemetry.Material)

			printerVersion := prometheus.MustNewConstMetric(
				collector.printerVersion, prometheus.GaugeValue,
				1,
				s.Address, s.Type, s.Name, version.API, version.Server, version.Text)

			if len(files.Files) > 0 {
				printerFiles := prometheus.MustNewConstMetric(
					collector.printerFiles, prometheus.GaugeValue,
					float64(len(files.Files[0].Children)),
					s.Address, s.Type, s.Name, files.Files[0].Display)
				ch <- printerFiles
			}

			zHeight := prometheus.MustNewConstMetric(
				collector.printerZHeight, prometheus.GaugeValue,
				printer.Telemetry.ZHeight,
				s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

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
		} else {
			logger.Error(s.Address + " is unreachable")
		}
	}
}
