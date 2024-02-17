package prusalink

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/rs/zerolog/log"
)

func getSLResponse(printer config.Printers) (Version, Files, Job, Printer, PrinterProfiles, error) {
	var (
		version         Version
		files           Files
		job             Job
		printerData     Printer
		printerprofiles PrinterProfiles
		err             error
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

	printerprofiles, err = GetPrinterProfiles(printer)

	if err != nil {
		log.Error().Msg("Error getting printerprofiles" + err.Error())
	}

	return version, files, job, printerData, printerprofiles, err
}

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
