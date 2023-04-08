package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type einsyCollector struct {
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
}

func newEinsyCollector() *einsyCollector {
	return &einsyCollector{
		printerNozzleTemp: prometheus.NewDesc("prusa_nozzle_temperature",
			"Current temperature of printer nozzle in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerBedTemp: prometheus.NewDesc("prusa_bed_temperature",
			"Current temperature of printer bed in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerVersion: prometheus.NewDesc("prusa_version",
			"Return information about printer. This metric contains information mostly about Prusa Link",
			[]string{"printer_address", "printer_model", "printer_name", "printer_api", "printer_server", "printer_text"},
			nil),
		printerZHeight: prometheus.NewDesc("prusa_z_height",
			"Current height of Z",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrintSpeed: prometheus.NewDesc("prusa_print_speed",
			"Current setting of printer speed in percents (%)",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerTargetTempNozzle: prometheus.NewDesc("prusa_nozzle_target_temperature",
			"Target temperature of printer nozzle in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerTargetTempBed: prometheus.NewDesc("prusa_bed_target_temperature",
			"Target temperature of printer bed in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerFiles: prometheus.NewDesc("prusa_files",
			"Number of files in storage",
			[]string{"printer_address", "printer_model", "printer_name", "printer_storage"},
			nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_printing_time_remaining",
			"Returns time that remains for completion of current print",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrintProgress: prometheus.NewDesc("prusa_printing_progress",
			"Returns information about completion of current print in percents",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrinting: prometheus.NewDesc("prusa_printing",
			"Return information about printing",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerMaterial: prometheus.NewDesc("prusa_material",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "printer_filament"},
			nil),
		printerPrintTime: prometheus.NewDesc("prusa_print_time",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
	}
}

func (collector *einsyCollector) Describe(ch chan<- *prometheus.Desc) {
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

func (collector *einsyCollector) Collect(ch chan<- prometheus.Metric) {
	cfg := loadCfg(getCfgFile())

	for _, s := range cfg.Printers.Einsy {
		log.Println("Einsy scraping at " + s.Address)

	}
}
