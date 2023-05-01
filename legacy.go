package main

import (
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type legacyCollector struct {
	printerNozzleTemp         *prometheus.Desc
	printerBedTemp            *prometheus.Desc
	printerZHeight            *prometheus.Desc
	printerPrintSpeed         *prometheus.Desc
	printerPrintTime          *prometheus.Desc
	printerPrintTimeRemaining *prometheus.Desc
	printerPrintProgress      *prometheus.Desc
	printerPrinting           *prometheus.Desc
	printerMaterial           *prometheus.Desc
}

func newLegacyCollector() *legacyCollector {
	return &legacyCollector{
		printerNozzleTemp:         prometheus.NewDesc("prusa_legacy_nozzle_temperature", "Current temperature of printer nozzle in Celsius", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerBedTemp:            prometheus.NewDesc("prusa_legacy_bed_temperature", "Current temperature of printer bed in Celsius", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerZHeight:            prometheus.NewDesc("prusa_legacy_z_height", "Current height of Z", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerPrintSpeed:         prometheus.NewDesc("prusa_legacy_print_speed", "Current setting of printer speed in percents (%)", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerPrintTime:          prometheus.NewDesc("prusa_legacy_print_time", "Returns information about loaded filament. Returns 0 if there is no loaded filament", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_legacy_printing_time_remaining", "Returns time that remains for completion of current print", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerPrintProgress:      prometheus.NewDesc("prusa_legacy_printing_progress", "Returns information about completion of current print in percents", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerPrinting:           prometheus.NewDesc("prusa_legacy_printing", "Return information about printing", []string{"printer_address", "printer_model", "printer_name", "printer_job_name"}, nil),
		printerMaterial:           prometheus.NewDesc("prusa_legacy_material", "Returns information about loaded filament. Returns 0 if there is no loaded filament", []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_filament"}, nil),
	}
}

func (collector *legacyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printerNozzleTemp
	ch <- collector.printerBedTemp
	ch <- collector.printerZHeight
	ch <- collector.printerPrintSpeed
	ch <- collector.printerPrintTime
	ch <- collector.printerPrintTimeRemaining
	ch <- collector.printerPrintProgress
	ch <- collector.printerPrinting
	ch <- collector.printerMaterial
}

func parseDuration(time string) float64 {
	var result float64
	result = 0
	parsedTime := strings.Split(time, " ")

	for _, v := range parsedTime {
		if strings.Contains(v, "h") {
			hours, _ := strconv.ParseFloat(strings.Trim(v, "h"), 32)
			result = result + (hours * 60 * 60)
		} else if strings.Contains(v, "m") {
			mins, _ := strconv.ParseFloat(strings.Trim(v, "m"), 32)
			result = result + (mins * 60)
		} else if strings.Contains(v, "s") {
			sec, _ := strconv.ParseFloat(strings.Trim(v, "s"), 32)
			result = result + (sec)
		}
	}
	return result
}

func (collector *legacyCollector) Collect(ch chan<- prometheus.Metric) {
	cfg := loadedConfig

	for _, s := range cfg.Printers.Legacy {
		logger.Debug("Legacy scraping at " + s.Address)
		if connTest("http://" + s.Address) {
			telemetry := getLegacyTelemetry(s.Address)

			nozzleTemp := prometheus.MustNewConstMetric(
				collector.printerNozzleTemp, prometheus.GaugeValue,
				float64(telemetry.TempNozzle),
				s.Address, s.Type, s.Name, telemetry.ProjectName)

			bedTemp := prometheus.MustNewConstMetric(
				collector.printerBedTemp, prometheus.GaugeValue, // collector
				float64(telemetry.TempBed),                       // value
				s.Address, s.Type, s.Name, telemetry.ProjectName) // labels

			printProgress := prometheus.MustNewConstMetric(
				collector.printerPrintProgress, prometheus.GaugeValue,
				float64(telemetry.Progress)/100,
				s.Address, s.Type, s.Name, telemetry.ProjectName)

			printSpeed := prometheus.MustNewConstMetric(
				collector.printerPrintSpeed, prometheus.GaugeValue,
				float64(telemetry.PrintingSpeed),
				s.Address, s.Type, s.Name, telemetry.ProjectName)

			time_est, _ := strconv.ParseFloat(telemetry.TimeEst, 32)

			printTimeRemaining := prometheus.MustNewConstMetric(
				collector.printerPrintTimeRemaining, prometheus.GaugeValue,
				time_est,
				s.Address, s.Type, s.Name, telemetry.ProjectName)

			printingMetric := 0
			if telemetry.TimeEst != "" {
				printingMetric = 1
			}

			printing := prometheus.MustNewConstMetric(
				collector.printerPrinting, prometheus.GaugeValue,
				float64(printingMetric),
				s.Address, s.Type, s.Name, telemetry.ProjectName)

			printTime := prometheus.MustNewConstMetric(
				collector.printerPrintTime, prometheus.GaugeValue,
				float64(parseDuration(telemetry.PrintDur)),
				s.Address, s.Type, s.Name, telemetry.ProjectName)

			filamentLoaded := 0
			if telemetry.Material != "---" {
				filamentLoaded = 1
			}

			material := prometheus.MustNewConstMetric(
				collector.printerMaterial, prometheus.GaugeValue,
				float64(filamentLoaded),
				s.Address, s.Type, s.Name, telemetry.ProjectName, telemetry.Material)

			zHeight := prometheus.MustNewConstMetric(
				collector.printerZHeight, prometheus.GaugeValue,
				telemetry.PosZMm,
				s.Address, s.Type, s.Name, telemetry.ProjectName)

			ch <- bedTemp
			ch <- nozzleTemp
			ch <- printProgress
			ch <- printSpeed
			ch <- printTimeRemaining
			ch <- printing
			ch <- printTime
			ch <- material
			ch <- zHeight

		} else {
			logger.Error(s.Address + " is unreachable")
		}
	}
}
