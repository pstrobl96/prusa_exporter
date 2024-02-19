package prusalink

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

// NewSLCollector creates a new SL collector
func NewSLCollector() *slCollector {
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
	ch <- c.printerUp
	ch <- c.printerBedTempTarget
	ch <- c.printerBedTempOffset
	ch <- c.printerChamberTempTarget
	ch <- c.printerChamberTempOffset
	ch <- c.printerToolTempTarget
	ch <- c.printerToolTempOffset
}

func (c *slCollector) Collect(ch chan<- prometheus.Metric) {
	for _, s := range configuration.Printers {
		log.Debug().Msg("SL scraping at " + s.Address)
		if !s.Reachable && s.Type == "SL1" {
			printerUp := prometheus.MustNewConstMetric(c.printerUp, prometheus.GaugeValue,
				0, s.Address, s.Type, s.Name)

			ch <- printerUp
			log.Debug().Msg(s.Address + " is unreachable while scraping")
		} else if s.Type == "SL1" {
			version, _, _, printer, _, err := getSLResponse(s)
			if err != nil {
				log.Error().Msg(err.Error())
			} else {
				printerUp := prometheus.MustNewConstMetric(c.printerUp, prometheus.GaugeValue,
					1, GetLabels(s, Job{})...)

				printerStatus := prometheus.MustNewConstMetric(c.printerStatus, prometheus.GaugeValue,
					getStateFlag(printer), GetLabels(s, Job{}, printer.State.Text)...)

				printerCover := prometheus.MustNewConstMetric(c.printerCover, prometheus.GaugeValue,
					BoolToFloat(printer.Telemetry.CoverClosed), GetLabels(s, Job{})...)

				printerFanBlower := prometheus.MustNewConstMetric(c.printerFanBlower, prometheus.GaugeValue,
					float64(printer.Telemetry.FanBlower), GetLabels(s, Job{})...)

				printerFanRear := prometheus.MustNewConstMetric(c.printerFanRear, prometheus.GaugeValue,
					float64(printer.Telemetry.FanRear), GetLabels(s, Job{})...)

				printerFanUV := prometheus.MustNewConstMetric(c.printerFanUV, prometheus.GaugeValue,
					float64(printer.Telemetry.FanUvLed), GetLabels(s, Job{})...)

				printerAmbientTemp := prometheus.MustNewConstMetric(c.printerAmbientTemp, prometheus.GaugeValue,
					float64(printer.Telemetry.TempAmbient), GetLabels(s, Job{})...)

				printerCPUTemp := prometheus.MustNewConstMetric(c.printerCPUTemp, prometheus.GaugeValue,
					float64(printer.Telemetry.TempCPU), GetLabels(s, Job{})...)

				pritnerUVTemp := prometheus.MustNewConstMetric(c.pritnerUVTemp, prometheus.GaugeValue,
					float64(printer.Telemetry.TempUvLed), GetLabels(s, Job{})...)

				printerBedTemp := prometheus.MustNewConstMetric(c.printerBedTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Actual), GetLabels(s, Job{})...)

				printerBedTempTarget := prometheus.MustNewConstMetric(c.printerBedTempTarget, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Target), GetLabels(s, Job{})...)

				printerBedTempOffset := prometheus.MustNewConstMetric(c.printerBedTempOffset, prometheus.GaugeValue,
					float64(printer.Temperature.Bed.Offset), GetLabels(s, Job{})...)

				printerChamberTempTarget := prometheus.MustNewConstMetric(c.printerChamberTempTarget, prometheus.GaugeValue,
					float64(printer.Temperature.Chamber.Target), GetLabels(s, Job{})...)

				printerChamberTempOffset := prometheus.MustNewConstMetric(c.printerChamberTempOffset, prometheus.GaugeValue,
					float64(printer.Temperature.Chamber.Offset), GetLabels(s, Job{})...)

				printerToolTempTarget := prometheus.MustNewConstMetric(c.printerToolTempTarget, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Target), GetLabels(s, Job{})...)

				printerToolTempOffset := prometheus.MustNewConstMetric(c.printerToolTempOffset, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Offset), GetLabels(s, Job{})...)

				printerChamberTemp := prometheus.MustNewConstMetric(c.printerChamberTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Chamber.Actual), GetLabels(s, Job{})...)

				printerToolTemp := prometheus.MustNewConstMetric(c.printerToolTemp, prometheus.GaugeValue,
					float64(printer.Temperature.Tool0.Actual), GetLabels(s, Job{})...)

				printerVersion := prometheus.MustNewConstMetric(c.printerVersion, prometheus.GaugeValue,
					1, GetLabels(s, Job{}, version.API, version.Server, version.Text, version.Hostname)...)

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
