package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
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
	defaultLabels := []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}
	return &buddyCollector{
		printerNozzleTemp: prometheus.NewDesc("prusa_buddy_nozzle_temperature",
			"Current temperature of printer nozzle in Celsius",
			defaultLabels,
			nil),
		printerBedTemp: prometheus.NewDesc("prusa_buddy_bed_temperature",
			"Current temperature of printer bed in Celsius",
			defaultLabels,
			nil),
		printerVersion: prometheus.NewDesc("prusa_buddy_version",
			"Return information about printer. This metric contains information mostly about Prusa Link",
			append(defaultLabels, "printer_api", "printer_server", "printer_text"),
			nil),
		printerZHeight: prometheus.NewDesc("prusa_buddy_z_height",
			"Current height of Z",
			defaultLabels,
			nil),
		printerPrintSpeed: prometheus.NewDesc("prusa_buddy_print_speed",
			"Current setting of printer speed in percents (%)",
			defaultLabels,
			nil),
		printerTargetTempNozzle: prometheus.NewDesc("prusa_buddy_nozzle_target_temperature",
			"Target temperature of printer nozzle in Celsius",
			defaultLabels,
			nil),
		printerTargetTempBed: prometheus.NewDesc("prusa_buddy_bed_target_temperature",
			"Target temperature of printer bed in Celsius",
			defaultLabels,
			nil),
		printerFiles: prometheus.NewDesc("prusa_buddy_files",
			"Number of files in storage",
			append(defaultLabels, "printer_storage"),
			nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_buddy_printing_time_remaining",
			"Returns time that remains for completion of current print",
			defaultLabels,
			nil),
		printerPrintProgress: prometheus.NewDesc("prusa_buddy_printing_progress",
			"Returns information about completion of current print in percents",
			defaultLabels,
			nil),
		printerPrinting: prometheus.NewDesc("prusa_buddy_printing",
			"Return information about printing",
			defaultLabels,
			nil),
		printerMaterial: prometheus.NewDesc("prusa_buddy_material",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			append(defaultLabels, "printer_filament"),
			nil),
		printerPrintTime: prometheus.NewDesc("prusa_buddy_print_time",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			defaultLabels,
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

func BoolToFloat(boolean bool) float64 {
	if boolean {
		return 1.0
	} else {
		return 0.0
	}
}

func getLabels(s buddy, job buddyJob, labelValues ...string) []string {
	return append([]string{s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path}, labelValues...)
}

func (collector *buddyCollector) Collect(ch chan<- prometheus.Metric) {
	cfg := &config
	for _, s := range cfg.Printers.Buddy {
		log.Debug().Msg("Buddy scraping at " + s.Address)
		if !s.Reachable {
			log.Debug().Msg(s.Address + " is unreachable while scraping")
		} else {
			version, files, job, printer, err := getBuddyResponse(s)

			if err != nil {
				log.Error().Msg(err.Error())
			} else {
				bedTemp := prometheus.MustNewConstMetric(collector.printerBedTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Actual), getLabels(s, job)...)

				nozzleTemp := prometheus.MustNewConstMetric(collector.printerNozzleTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Actual), getLabels(s, job)...)

				printProgress := prometheus.MustNewConstMetric(collector.printerPrintProgress, prometheus.GaugeValue,
					float64(job.Progress.Completion), getLabels(s, job)...)

				printSpeed := prometheus.MustNewConstMetric(collector.printerPrintSpeed, prometheus.GaugeValue,
					float64(printer.Telemetry.PrintSpeed), getLabels(s, job)...)

				printTimeRemaining := prometheus.MustNewConstMetric(collector.printerPrintTimeRemaining, prometheus.GaugeValue,
					float64(job.Progress.PrintTimeLeft), getLabels(s, job)...)

				printTime := prometheus.MustNewConstMetric(collector.printerPrintTime, prometheus.GaugeValue,
					float64(job.Progress.PrintTime), getLabels(s, job)...)

				targetTempBed := prometheus.MustNewConstMetric(collector.printerTargetTempBed, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Target), getLabels(s, job)...)

				targetTempNozzle := prometheus.MustNewConstMetric(collector.printerTargetTempNozzle, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Target), getLabels(s, job)...)

				zHeight := prometheus.MustNewConstMetric(collector.printerZHeight, prometheus.GaugeValue,
					printer.Telemetry.ZHeight, getLabels(s, job)...)

				printing := prometheus.MustNewConstMetric(collector.printerPrinting, prometheus.GaugeValue,
					float64(BoolToFloat(strings.Contains(job.State, "Printing"))), getLabels(s, job)...)

				printerVersion := prometheus.MustNewConstMetric(collector.printerVersion, prometheus.GaugeValue,
					1, getLabels(s, job, version.API, version.Server, version.Text)...)

				material := prometheus.MustNewConstMetric(collector.printerMaterial, prometheus.GaugeValue,
					float64(BoolToFloat(strings.Contains(printer.Telemetry.Material, "---"))),
					getLabels(s, job, printer.Telemetry.Material)...)

				if len(files.Files) > 0 {
					printerFiles := prometheus.MustNewConstMetric(collector.printerFiles, prometheus.GaugeValue,
						float64(len(files.Files[0].Children)), getLabels(s, job, files.Files[0].Display)...)
					ch <- printerFiles
				}

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
			}
		}
	}
}
