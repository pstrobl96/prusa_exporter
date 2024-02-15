package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

type slCollector struct {
	printerStatus            *prometheus.Desc
	printerCover             *prometheus.Desc
	printerFanBlower         *prometheus.Desc
	printerFanRear           *prometheus.Desc
	printerFanUV             *prometheus.Desc
	printerAmbientTemp       *prometheus.Desc
	printerCPUTemp           *prometheus.Desc
	pritnerUVTemp            *prometheus.Desc
	printerBedTemp           *prometheus.Desc
	printerBedTempTarget     *prometheus.Desc
	printerBedTempOffset     *prometheus.Desc
	printerChamberTemp       *prometheus.Desc
	printerChamberTempTarget *prometheus.Desc
	printerChamberTempOffset *prometheus.Desc
	printerToolTemp          *prometheus.Desc
	printerToolTempTarget    *prometheus.Desc
	printerToolTempOffset    *prometheus.Desc
	printerHeatedChamber     *prometheus.Desc
	printerVersion           *prometheus.Desc
	printerFiles             *prometheus.Desc
	printerUp                *prometheus.Desc
}

func getSlFlag(printer SLPrinter) float64 {
	if printer.State.Flags.Operational {
		return 1
	} else if printer.State.Flags.Operational {
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
	} else {
		return 0
	}
}

func getSlLabels(s sl, labelValues ...string) []string {
	return append([]string{s.Address, s.Type, s.Name}, labelValues...)
}

func newSLCollector() *slCollector {
	defaultLabels := []string{"printer_address", "printer_model", "printer_name"}
	return &slCollector{
		printerStatus:            prometheus.NewDesc("prusa_sl_status", "Status of the printer", append(defaultLabels, "printer_status"), nil),
		printerCover:             prometheus.NewDesc("prusa_sl_cover", "Status of the printer - 0 = open, 1 = closed", defaultLabels, nil),
		printerFanBlower:         prometheus.NewDesc("prusa_sl_fan_blower", "Status of the printer blower fan", defaultLabels, nil),
		printerFanRear:           prometheus.NewDesc("prusa_sl_fan_rear", "Status of the printer fan rear", defaultLabels, nil),
		printerFanUV:             prometheus.NewDesc("prusa_sl_fan_uv", "Status of the printer fan uv", defaultLabels, nil),
		printerAmbientTemp:       prometheus.NewDesc("prusa_sl_ambient_temperature", "Status of the printer ambient temperature", defaultLabels, nil),
		printerCPUTemp:           prometheus.NewDesc("prusa_sl_cpu_temperature", "Status of the printer cpu temperature", defaultLabels, nil),
		pritnerUVTemp:            prometheus.NewDesc("prusa_sl_uv_temperature", "Status of the printer uv temperature", defaultLabels, nil),
		printerBedTemp:           prometheus.NewDesc("prusa_sl_bed_temperature", "Status of the printer bed temperature", defaultLabels, nil),
		printerBedTempTarget:     prometheus.NewDesc("prusa_sl_bed_target_temperature", "Target bed temperature", defaultLabels, nil),
		printerBedTempOffset:     prometheus.NewDesc("prusa_sl_bed_offset_temperature", "Offset bed temperature", defaultLabels, nil),
		printerChamberTemp:       prometheus.NewDesc("prusa_sl_chamber_temperature", "Status of the printer chamber temperature", defaultLabels, nil),
		printerChamberTempTarget: prometheus.NewDesc("prusa_sl_chamber_target_temperature", "Traget chamber temperature", defaultLabels, nil),
		printerChamberTempOffset: prometheus.NewDesc("prusa_sl_chamber_offset_temperature", "Offset chamber temperature", defaultLabels, nil),
		printerToolTemp:          prometheus.NewDesc("prusa_sl_tool_temperature", "Status of the printer tool temperature", defaultLabels, nil),
		printerToolTempTarget:    prometheus.NewDesc("prusa_sl_tool_target_temperature", "Target tool temperature", defaultLabels, nil),
		printerToolTempOffset:    prometheus.NewDesc("prusa_sl_tool_offset_temperature", "Offset tool temperature", defaultLabels, nil),
		printerHeatedChamber:     prometheus.NewDesc("prusa_sl_heated_chamber", "Status of the printer heated chamber", defaultLabels, nil),
		printerVersion:           prometheus.NewDesc("prusa_sl_version", "Returns value from version endpoint", append(defaultLabels, "printer_api", "printer_server", "printer_text", "printer_hostname"), nil),
		printerFiles:             prometheus.NewDesc("prusa_sl_files", "Number of the printer files", append(defaultLabels, "printer_storage"), nil),
		printerUp:                prometheus.NewDesc("prusa_sl_up", "Status of the printer", defaultLabels, nil),
	}
}

func (c *slCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.printerStatus
	ch <- c.printerCover
	ch <- c.printerFanBlower
	ch <- c.printerFanRear
	ch <- c.printerFanUV
	ch <- c.printerAmbientTemp
	ch <- c.printerCPUTemp
	ch <- c.pritnerUVTemp
	ch <- c.printerBedTemp
	ch <- c.printerChamberTemp
	ch <- c.printerToolTemp
	ch <- c.printerHeatedChamber
	ch <- c.printerVersion
	ch <- c.printerFiles
	ch <- c.printerUp
	ch <- c.printerBedTempTarget
	ch <- c.printerBedTempOffset
	ch <- c.printerChamberTempTarget
	ch <- c.printerChamberTempOffset
	ch <- c.printerToolTempTarget
	ch <- c.printerToolTempOffset
}

func (c *slCollector) Collect(ch chan<- prometheus.Metric) {
	cfg := &config
	for _, s := range cfg.Printers.Sl {
		log.Debug().Msg("SL scraping at " + s.Address)
		if !s.Reachable {
			printerUp := prometheus.MustNewConstMetric(c.printerUp, prometheus.GaugeValue,
				0, s.Address, s.Type, s.Name)

			ch <- printerUp

			log.Debug().Msg(s.Address + " is unreachable while scraping")
		} else {
			_, _, printer, _, version, err := getSLResponse(s)
			if err != nil {
				log.Error().Msg(err.Error())
			} else {
				printerUp := prometheus.MustNewConstMetric(c.printerUp, prometheus.GaugeValue,
					1, getSlLabels(s)...)

				printerStatus := prometheus.MustNewConstMetric(c.printerStatus, prometheus.GaugeValue,
					getSlFlag(printer), getSlLabels(s, printer.State.Text)...)

				printerCover := prometheus.MustNewConstMetric(c.printerCover, prometheus.GaugeValue,
					BoolToFloat(printer.Telemetry.CoverClosed), getSlLabels(s)...)

				printerFanBlower := prometheus.MustNewConstMetric(c.printerFanBlower, prometheus.GaugeValue,
					float64(printer.Telemetry.FanBlower), getSlLabels(s)...)

				printerFanRear := prometheus.MustNewConstMetric(c.printerFanRear, prometheus.GaugeValue,
					float64(printer.Telemetry.FanRear), getSlLabels(s)...)

				printerFanUV := prometheus.MustNewConstMetric(c.printerFanUV, prometheus.GaugeValue,
					float64(printer.Telemetry.FanUvLed), getSlLabels(s)...)

				printerAmbientTemp := prometheus.MustNewConstMetric(c.printerAmbientTemp, prometheus.GaugeValue,
					float64(printer.Telemetry.TempAmbient), getSlLabels(s)...)

				printerCPUTemp := prometheus.MustNewConstMetric(c.printerCPUTemp, prometheus.GaugeValue,
					float64(printer.Telemetry.TempCPU), getSlLabels(s)...)

				pritnerUVTemp := prometheus.MustNewConstMetric(c.pritnerUVTemp, prometheus.GaugeValue,
					float64(printer.Telemetry.TempUvLed), getSlLabels(s)...)

				printerBedTemp := prometheus.MustNewConstMetric(c.printerBedTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Actual), getSlLabels(s)...)

				printerBedTempTarget := prometheus.MustNewConstMetric(c.printerBedTempTarget, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Target), getSlLabels(s)...)

				printerBedTempOffset := prometheus.MustNewConstMetric(c.printerBedTempOffset, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Offset), getSlLabels(s)...)

				printerChamberTempTarget := prometheus.MustNewConstMetric(c.printerChamberTempTarget, prometheus.GaugeValue,
					float64(printer.Temperature.Chamber.Target), getSlLabels(s)...)

				printerChamberTempOffset := prometheus.MustNewConstMetric(c.printerChamberTempOffset, prometheus.GaugeValue,
					float64(printer.Temperature.Chamber.Offset), getSlLabels(s)...)

				printerToolTempTarget := prometheus.MustNewConstMetric(c.printerToolTempTarget, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Target), getSlLabels(s)...)

				printerToolTempOffset := prometheus.MustNewConstMetric(c.printerToolTempOffset, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Offset), getSlLabels(s)...)

				printerChamberTemp := prometheus.MustNewConstMetric(c.printerChamberTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Chamber.Actual), getSlLabels(s)...)

				printerToolTemp := prometheus.MustNewConstMetric(c.printerToolTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Actual), getSlLabels(s)...)

				printerVersion := prometheus.MustNewConstMetric(c.printerVersion, prometheus.GaugeValue,
					1, getSlLabels(s, version.API, version.Server, version.Text, version.Hostname)...)

				//printerFiles := prometheus.MustNewConstMetric(c.printerFiles, prometheus.GaugeValue,
				//	float64(len(files.Files[0].Children)), getSlLabels(s, files.Files[0].Display)...)

				ch <- printerStatus
				ch <- printerCover
				ch <- printerFanBlower
				ch <- printerFanRear
				ch <- printerFanUV
				ch <- printerAmbientTemp
				ch <- printerCPUTemp
				ch <- pritnerUVTemp
				ch <- printerBedTemp
				ch <- printerChamberTemp
				ch <- printerToolTemp
				ch <- printerVersion
				//ch <- printerFiles
				ch <- printerUp
				ch <- printerBedTempTarget
				ch <- printerBedTempOffset
				ch <- printerChamberTempTarget
				ch <- printerChamberTempOffset
				ch <- printerToolTempTarget
				ch <- printerToolTempOffset
			}
		}
	}
}
