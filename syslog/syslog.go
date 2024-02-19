package syslog

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/pstrobl96/prusa_exporter/config"
	"github.com/pstrobl96/prusa_exporter/prusalink"
	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

var syslogData = make(map[string]map[string]string)

func startSyslogServer(listenUDP string) (syslog.LogPartsChannel, *syslog.Server) {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenUDP(listenUDP)
	server.Boot()
	return channel, server
}

// HandleMetrics is function that listens for syslog messages and parses them into map
func HandleMetrics(listenUDP string) {
	channel, server := startSyslogServer(listenUDP)

	patterns := []struct {
		pattern string
		fields  []string
	}{
		{pattern: `(?P<name>\w+_[a-z]+) v=(?P<value>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}},
		{pattern: `(?P<name>\w+_[a-z]+) v=(?P<value>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}},  // integer
		{pattern: `(?P<name>\w+_[a-z]+) v="(?P<value>[-\d\.]+)" (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}}, // made for string values
		{pattern: `(?P<name>\w+(?:,[a-z]=\d+)?)[ ]v=(?P<value>[-\d\.]+),e=(?P<e>[-\d\.]+)[ ](?P<timestamp>\d+)`, fields: []string{"name", "value", "e", "timestamp"}},
		{pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) v=(?P<value>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}},
		{pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) free=(?P<subvalue>[-\d\.]+)i,total=(?P<subvalue2>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "subvalue", "subvalue2", "timestamp"}},
	}

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			clientIP := strings.Split(logParts["client"].(string), ":")[0] // getting rid of port and leaving only ip address
			if clientIP == "" {                                            // Skip empty client ip
				continue
			} else {
				if syslogData[clientIP] == nil {
					syslogData[clientIP] = make(map[string]string)
				} // Initialize map for ip address if it doesn't exist - is it unique? No. Is it a problem? No. Is it experimental? Yes.

				syslogData[clientIP]["mac"] = logParts["hostname"].(string)

				for _, pattern := range patterns {
					reg, err := regexp.Compile(pattern.pattern)
					if err != nil {
						log.Error().Msg("Error compiling regexp: " + err.Error())
						return
					}

					matches := reg.FindAllStringSubmatch(logParts["message"].(string), -1)
					if matches == nil {
						continue // No matches for this pattern
					}

					for _, match := range matches {
						// Extract values based on named groups
						var valueStr string
						for i, field := range pattern.fields {
							switch field {
							case "value":
								valueStr = match[i+1]
							case "subvalue":
								valueStr = match[i+1]
							case "subvalue2":
								// Handle combining subvalues if needed
								if valueStr != "" {
									valueStr += "," + match[i+1]
								} else {
									valueStr = match[i+1]
								}
							}
						}

						syslogData[clientIP][match[1]] = fmt.Sprint(valueStr)
					}
				}
			}
		}
	}(channel)

	server.Wait()
}

type syslogCollector struct {
	// power metrics
	printerVolt5V             *prometheus.Desc
	printerVolt24V            *prometheus.Desc
	printerVoltBed            *prometheus.Desc
	printerVoltNozzle         *prometheus.Desc
	printerVoltSandwich5V     *prometheus.Desc
	printerVoltSplitter5V     *prometheus.Desc
	printerCurrentXlbuddy5V   *prometheus.Desc
	printerCurrentInput       *prometheus.Desc
	printerCurrentMMU         *prometheus.Desc
	printerCurrentBed         *prometheus.Desc
	printerCurrentNozzle      *prometheus.Desc
	printerCurrentDwarfHeater *prometheus.Desc // dwarf_heat_curr
	printerCurrentBedlet      *prometheus.Desc // bedlet_curr
	printerOvercurrentNozzle  *prometheus.Desc
	printerOvercurrentInput   *prometheus.Desc

	// printer metrics
	printerBedletTemp         *prometheus.Desc // bedlet_temp
	printerBedletState        *prometheus.Desc // bedlet_state
	printerProbeZ             *prometheus.Desc // probe_z
	printerBedMcuTemp         *prometheus.Desc // bed_mcu_temp
	printerLoadcellValue      *prometheus.Desc // loadcell_value
	printerSandwitchTemp      *prometheus.Desc // temp_sandwich
	printerSplitterTemp       *prometheus.Desc // temp_splitter
	printerDwarfsBoardTemp    *prometheus.Desc // dwarfs_board_temp
	printerHeatbreakTemp      *prometheus.Desc // temp_hbr
	printerBoardTemp          *prometheus.Desc // temp_brd
	printerChamberTemp        *prometheus.Desc // temp_chamber
	printerMcuTemp            *prometheus.Desc // temp_mcu
	printerFSensor            *prometheus.Desc // fsensor
	printerSideFSensor        *prometheus.Desc // side_fsensor
	printerActiveExtruder     *prometheus.Desc
	printerDwarfMcuTemp       *prometheus.Desc
	printerDwarfBoardTemp     *prometheus.Desc
	printerAxisZAdjustment    *prometheus.Desc
	printerHeaterEnabled      *prometheus.Desc
	printerLoadcellScale      *prometheus.Desc
	printerLoadcellThreshold  *prometheus.Desc
	printerLoadcellHysteresis *prometheus.Desc

	// system metrics
	printerBuddySyslogInfo *prometheus.Desc // revision, bom
	printerCPUUsage        *prometheus.Desc
	printerHeapTotal       *prometheus.Desc
	printerHeapFree        *prometheus.Desc
	printerPointsDropped   *prometheus.Desc
	printerMediaPrefetched *prometheus.Desc
}

// NewSyslogCollector is a function that returns new syslogCollector
func NewSyslogCollector() *syslogCollector {
	defaultLabels := []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}
	return &syslogCollector{
		printerBedletTemp: prometheus.NewDesc("prusa_buddy_bedlet_temp",
			"Bedlet temperature",
			defaultLabels,
			nil),
		printerBedletState: prometheus.NewDesc("prusa_buddy_bedlet_state",
			"Bedlet state",
			defaultLabels,
			nil),
		printerProbeZ: prometheus.NewDesc("prusa_buddy_probe_z",
			"Probe Z",
			defaultLabels,
			nil),
		printerBedMcuTemp: prometheus.NewDesc("prusa_buddy_bed_mcu_temp",
			"Bed MCU temperature",
			defaultLabels,
			nil),
		printerLoadcellValue: prometheus.NewDesc("prusa_buddy_loadcell_value",
			"Loadcell value",
			defaultLabels,
			nil),
		printerSandwitchTemp: prometheus.NewDesc("prusa_buddy_temp_sandwich",
			"Sandwich temperature",
			defaultLabels,
			nil),
		printerSplitterTemp: prometheus.NewDesc("prusa_buddy_temp_splitter",
			"Splitter temperature",
			defaultLabels,
			nil),
		printerDwarfsBoardTemp: prometheus.NewDesc("prusa_buddy_dwarfs_board_temp",
			"Dwarfs board temperature",
			defaultLabels,
			nil),
		printerHeatbreakTemp: prometheus.NewDesc("prusa_buddy_temp_hbr",
			"Heatbreak temperature",
			defaultLabels,
			nil),
		printerBoardTemp: prometheus.NewDesc("prusa_buddy_temp_brd",
			"Board temperature",
			defaultLabels,
			nil),
		printerChamberTemp: prometheus.NewDesc("prusa_buddy_temp_chamber",
			"Chamber temperature",
			defaultLabels,
			nil),
		printerMcuTemp: prometheus.NewDesc("prusa_buddy_temp_mcu",
			"MCU temperature",
			defaultLabels,
			nil),
		printerFSensor: prometheus.NewDesc("prusa_buddy_fsensor",
			"Filament Sensor",
			defaultLabels,
			nil),
		printerSideFSensor: prometheus.NewDesc("prusa_buddy_side_fsensor",
			"Side Filament Sensor",
			defaultLabels,
			nil),
		printerCurrentDwarfHeater: prometheus.NewDesc("prusa_buddy_current_dwarf_heater",
			"Current of dwarf heater",
			defaultLabels,
			nil),
		printerCurrentBedlet: prometheus.NewDesc("prusa_buddy_current_bedlet",
			"Current of bedlet",
			defaultLabels,
			nil),
		printerVolt5V: prometheus.NewDesc("prusa_buddy_voltage_5volts",
			"Voltage of 5V rail",
			defaultLabels,
			nil),
		printerVolt24V: prometheus.NewDesc("prusa_buddy_voltage_24volts",
			"Voltage of 24V rail",
			defaultLabels,
			nil),
		printerVoltBed: prometheus.NewDesc("prusa_buddy_voltage_bed",
			"Voltage of bed",
			defaultLabels,
			nil),
		printerVoltNozzle: prometheus.NewDesc("prusa_buddy_voltage_nozzle",
			"Voltage of nozzle",
			defaultLabels,
			nil),
		printerVoltSandwich5V: prometheus.NewDesc("prusa_buddy_voltage_sandwich_5volts",
			"Voltage of sandwich 5V rail",
			defaultLabels,
			nil),
		printerVoltSplitter5V: prometheus.NewDesc("prusa_buddy_voltage_splitter_5volts",
			"Voltage of splitter 5V rail",
			defaultLabels,
			nil),
		printerCurrentXlbuddy5V: prometheus.NewDesc("prusa_buddy_current_xlbuddy_5volts",
			"Current of xlBuddy 5V rail",
			defaultLabels,
			nil),
		printerCurrentInput: prometheus.NewDesc("prusa_buddy_current_input",
			"Input current",
			defaultLabels,
			nil),
		printerCurrentMMU: prometheus.NewDesc("prusa_buddy_current_mmu",
			"Current of MMU",
			defaultLabels,
			nil),
		printerCurrentBed: prometheus.NewDesc("prusa_buddy_current_bed",
			"Current of bed",
			append(defaultLabels, "rail" /*XL has two 24V rails for heatbed*/),
			nil),
		printerCurrentNozzle: prometheus.NewDesc("prusa_buddy_current_nozzle",
			"Current of nozzle",
			defaultLabels,
			nil),
		printerOvercurrentNozzle: prometheus.NewDesc("prusa_buddy_overcurrent_nozzle",
			"Overcurrent of nozzle",
			defaultLabels,
			nil),
		printerOvercurrentInput: prometheus.NewDesc("prusa_buddy_overcurrent_input",
			"Overcurrent of input",
			defaultLabels,
			nil),
		printerActiveExtruder: prometheus.NewDesc("prusa_buddy_active_extruder",
			"Active extruder - used for XL",
			defaultLabels,
			nil),
		printerDwarfMcuTemp: prometheus.NewDesc("prusa_buddy_dwarf_mcu_temp",
			"Dwarf MCU temperature - used for XL",
			defaultLabels,
			nil),
		printerDwarfBoardTemp: prometheus.NewDesc("prusa_buddy_dwarf_board_temp",
			"Dwarf board temperature - used for XL",
			defaultLabels,
			nil),
		printerAxisZAdjustment: prometheus.NewDesc("prusa_buddy_axis_z_adjustment",
			"Axis Z adjustment",
			defaultLabels,
			nil),
		printerHeaterEnabled: prometheus.NewDesc("prusa_buddy_heater_enabled",
			"Heater enabled",
			defaultLabels,
			nil),
		printerLoadcellScale: prometheus.NewDesc("prusa_buddy_loadcell_scale",
			"Loadcell scale",
			defaultLabels,
			nil),
		printerLoadcellThreshold: prometheus.NewDesc("prusa_buddy_loadcell_threshold",
			"Loadcell threshold",
			defaultLabels,
			nil),
		printerLoadcellHysteresis: prometheus.NewDesc("prusa_buddy_loadcell_hysteresis",
			"Loadcell hysteresis",
			defaultLabels,
			nil),
		printerBuddySyslogInfo: prometheus.NewDesc("prusa_buddy_syslog_info",
			"Buddy info",
			append(defaultLabels, "buddy_revision", "buddy_bom"),
			nil),
		printerCPUUsage: prometheus.NewDesc("prusa_buddy_cpu_usage_ratio",
			"CPU usage from 0.0 to 1.0",
			defaultLabels,
			nil),
		printerHeapTotal: prometheus.NewDesc("prusa_buddy_heap_total",
			"Total heap",
			defaultLabels,
			nil),
		printerHeapFree: prometheus.NewDesc("prusa_buddy_heap_free",
			"Free heap",
			defaultLabels,
			nil),
		printerPointsDropped: prometheus.NewDesc("prusa_buddy_points_dropped",
			"Points dropped",
			defaultLabels,
			nil),
		printerMediaPrefetched: prometheus.NewDesc("prusa_buddy_media_prefetched",
			"Media prefetched",
			defaultLabels,
			nil),
	}
}

func (collector *syslogCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printerBedletTemp
	ch <- collector.printerBedletState
	ch <- collector.printerProbeZ
	ch <- collector.printerBedMcuTemp
	ch <- collector.printerLoadcellValue
	ch <- collector.printerSandwitchTemp
	ch <- collector.printerSplitterTemp
	ch <- collector.printerDwarfsBoardTemp
	ch <- collector.printerHeatbreakTemp
	ch <- collector.printerBoardTemp
	ch <- collector.printerChamberTemp
	ch <- collector.printerMcuTemp
	ch <- collector.printerFSensor
	ch <- collector.printerSideFSensor
	ch <- collector.printerCurrentDwarfHeater
	ch <- collector.printerCurrentBedlet
	ch <- collector.printerVolt5V
	ch <- collector.printerVolt24V
	ch <- collector.printerVoltBed
	ch <- collector.printerVoltNozzle
	ch <- collector.printerVoltSandwich5V
	ch <- collector.printerVoltSplitter5V
	ch <- collector.printerCurrentXlbuddy5V
	ch <- collector.printerCurrentInput
	ch <- collector.printerCurrentMMU
	ch <- collector.printerCurrentBed
	ch <- collector.printerCurrentNozzle
	ch <- collector.printerOvercurrentNozzle
	ch <- collector.printerOvercurrentInput
	ch <- collector.printerActiveExtruder
	ch <- collector.printerDwarfMcuTemp
	ch <- collector.printerDwarfBoardTemp
	ch <- collector.printerAxisZAdjustment
	ch <- collector.printerHeaterEnabled
	ch <- collector.printerLoadcellScale
	ch <- collector.printerLoadcellThreshold
	ch <- collector.printerLoadcellHysteresis
	ch <- collector.printerBuddySyslogInfo
	ch <- collector.printerCPUUsage
	ch <- collector.printerHeapTotal
	ch <- collector.printerHeapFree
	ch <- collector.printerPointsDropped
	ch <- collector.printerMediaPrefetched
}

func (collector *syslogCollector) Collect(ch chan<- prometheus.Metric, config config.Config) {
	for _, s := range config.Printers {
		log.Debug().Msg("SYSLOG - Buddy scraping at " + s.Address)
		if _, ok := syslogData[s.Address]; ok {

			log.Debug().Msg("SYSLOG - found data for: " + s.Address)
			if s.Reachable { // if not reachable then just do nothing

				job, err := prusalink.GetJob(s) // get job for labels
				if err != nil {
					log.Error().Msg(err.Error())
				} else {

					printerBedletTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["bedlet_temp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerBedletTemp := prometheus.MustNewConstMetric(collector.printerBedletTemp, prometheus.GaugeValue,
							printerBedletTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerBedletTemp
					}

					printerBedletStateParsed, e := strconv.ParseFloat(syslogData[s.Address]["bedlet_state"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerBedletState := prometheus.MustNewConstMetric(collector.printerBedletState, prometheus.GaugeValue,
							printerBedletStateParsed, prusalink.GetLabels(s, job)...)
						ch <- printerBedletState
					}

					printerProbeZParsed, e := strconv.ParseFloat(syslogData[s.Address]["probe_z"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerProbeZ := prometheus.MustNewConstMetric(collector.printerProbeZ, prometheus.GaugeValue,
							printerProbeZParsed, prusalink.GetLabels(s, job)...)
						ch <- printerProbeZ
					}

					printerBedMcuTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["bed_mcu_temp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerBedMcuTemp := prometheus.MustNewConstMetric(collector.printerBedMcuTemp, prometheus.GaugeValue,
							printerBedMcuTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerBedMcuTemp
					}

					printerLoadcellValueParsed, e := strconv.ParseFloat(syslogData[s.Address]["loadcell_value"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerLoadcellValue := prometheus.MustNewConstMetric(collector.printerLoadcellValue, prometheus.GaugeValue,
							printerLoadcellValueParsed, prusalink.GetLabels(s, job)...)
						ch <- printerLoadcellValue
					}

					printerSandwitchTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["temp_sandwich"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerSandwitchTemp := prometheus.MustNewConstMetric(collector.printerSandwitchTemp, prometheus.GaugeValue,
							printerSandwitchTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerSandwitchTemp
					}

					printerSplitterTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["temp_splitter"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerSplitterTemp := prometheus.MustNewConstMetric(collector.printerSplitterTemp, prometheus.GaugeValue,
							printerSplitterTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerSplitterTemp
					}

					printerDwarfsBoardTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["dwarfs_board_temp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerDwarfsBoardTemp := prometheus.MustNewConstMetric(collector.printerDwarfsBoardTemp, prometheus.GaugeValue,
							printerDwarfsBoardTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerDwarfsBoardTemp
					}

					printerHeatbreakTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["temp_hbr"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {

						printerHeatbreakTemp := prometheus.MustNewConstMetric(collector.printerHeatbreakTemp, prometheus.GaugeValue,
							printerHeatbreakTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerHeatbreakTemp
					}

					printerBoardTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["temp_brd"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerBoardTemp := prometheus.MustNewConstMetric(collector.printerBoardTemp, prometheus.GaugeValue,
							printerBoardTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerBoardTemp
					}

					printerChamberTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["temp_chamber"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerChamberTemp := prometheus.MustNewConstMetric(collector.printerChamberTemp, prometheus.GaugeValue,
							printerChamberTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerChamberTemp
					}

					printerMcuTempParsed, e := strconv.ParseFloat(syslogData[s.Address]["temp_mcu"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerMcuTemp := prometheus.MustNewConstMetric(collector.printerMcuTemp, prometheus.GaugeValue,
							printerMcuTempParsed, prusalink.GetLabels(s, job)...)
						ch <- printerMcuTemp
					}

					printerFSensorParsed, e := strconv.ParseFloat(syslogData[s.Address]["fsensor"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerFSensor := prometheus.MustNewConstMetric(collector.printerFSensor, prometheus.GaugeValue,
							printerFSensorParsed, prusalink.GetLabels(s, job)...)
						ch <- printerFSensor
					}

					printerSideFSensorParsed, e := strconv.ParseFloat(syslogData[s.Address]["side_fsensor"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerSideFSensor := prometheus.MustNewConstMetric(collector.printerSideFSensor, prometheus.GaugeValue,

							printerSideFSensorParsed, prusalink.GetLabels(s, job)...)
						ch <- printerSideFSensor
					}

					printerCurrentDwarfHeaterParsed, e := strconv.ParseFloat(syslogData[s.Address]["dwarf_heat_curr"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerCurrentDwarfHeater := prometheus.MustNewConstMetric(collector.printerCurrentDwarfHeater, prometheus.GaugeValue,
							printerCurrentDwarfHeaterParsed, prusalink.GetLabels(s, job)...)
						ch <- printerCurrentDwarfHeater
					}

					printerCurrentBedletParsed, e := strconv.ParseFloat(syslogData[s.Address]["bedlet_curr"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())
					} else {
						printerCurrentBedlet := prometheus.MustNewConstMetric(collector.printerCurrentBedlet, prometheus.GaugeValue,
							printerCurrentBedletParsed, prusalink.GetLabels(s, job)...)
						ch <- printerCurrentBedlet
					}

					printerVolt5vParsed, e := strconv.ParseFloat(syslogData[s.Address]["5VVoltage"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerVolt5V := prometheus.MustNewConstMetric(collector.printerVolt5V, prometheus.GaugeValue,
							printerVolt5vParsed, prusalink.GetLabels(s, job)...)
						ch <- printerVolt5V
					}

					printerVolt24vParsed, e := strconv.ParseFloat(syslogData[s.Address]["24VVoltage"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerVolt24V := prometheus.MustNewConstMetric(collector.printerVolt24V, prometheus.GaugeValue,
							printerVolt24vParsed, prusalink.GetLabels(s, job)...)
						ch <- printerVolt24V
					}

					printerVoltBedParsed, e := strconv.ParseFloat(syslogData[s.Address]["volt_bed"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerVoltBed := prometheus.MustNewConstMetric(collector.printerVoltBed, prometheus.GaugeValue,
							printerVoltBedParsed, prusalink.GetLabels(s, job)...)
						ch <- printerVoltBed
					}

					printerVoltNozzleParsed, e := strconv.ParseFloat(syslogData[s.Address]["volt_nozz"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerVoltNozzle := prometheus.MustNewConstMetric(collector.printerVoltNozzle, prometheus.GaugeValue,
							printerVoltNozzleParsed, prusalink.GetLabels(s, job)...)
						ch <- printerVoltNozzle
					}

					printerVoltSandwich5vParsed, e := strconv.ParseFloat(syslogData[s.Address]["Sandwitch5VCurrent"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerVoltSandwich5V := prometheus.MustNewConstMetric(collector.printerVoltSandwich5V, prometheus.GaugeValue,
							printerVoltSandwich5vParsed, prusalink.GetLabels(s, job)...)
						ch <- printerVoltSandwich5V
					}

					printerVoltSplitter5vParsed, e := strconv.ParseFloat(syslogData[s.Address]["splitter_5V_current"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {

						printerVoltSplitter5V := prometheus.MustNewConstMetric(collector.printerVoltSplitter5V, prometheus.GaugeValue,
							printerVoltSplitter5vParsed, prusalink.GetLabels(s, job)...)
						ch <- printerVoltSplitter5V
					}

					printerCurrentXlbuddy5vParsed, e := strconv.ParseFloat(syslogData[s.Address]["xlbuddy5VCurrent"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerCurrentXlbuddy5V := prometheus.MustNewConstMetric(collector.printerCurrentXlbuddy5V, prometheus.GaugeValue,
							printerCurrentXlbuddy5vParsed, prusalink.GetLabels(s, job)...)
						ch <- printerCurrentXlbuddy5V
					}

					printerCurrentInputParsed, e := strconv.ParseFloat(syslogData[s.Address]["curr_inp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerCurrentInput := prometheus.MustNewConstMetric(collector.printerCurrentInput, prometheus.GaugeValue,
							printerCurrentInputParsed, prusalink.GetLabels(s, job)...)
						ch <- printerCurrentInput
					}

					printerCurrentMMUParsed, e := strconv.ParseFloat(syslogData[s.Address]["cur_mmu_imp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerCurrentMMU := prometheus.MustNewConstMetric(collector.printerCurrentMMU, prometheus.GaugeValue,
							printerCurrentMMUParsed, prusalink.GetLabels(s, job)...)
						ch <- printerCurrentMMU
					}

					printerCurrentBedParsed, e := strconv.ParseFloat(syslogData[s.Address]["bed_curr,n=0"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerCurrentBed := prometheus.MustNewConstMetric(collector.printerCurrentBed, prometheus.GaugeValue,
							printerCurrentBedParsed, prusalink.GetLabels(s, job, "0")...)
						ch <- printerCurrentBed
					}

					printerCurrentBedParsed, e = strconv.ParseFloat(syslogData[s.Address]["bed_curr,n=1"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerCurrentBed := prometheus.MustNewConstMetric(collector.printerCurrentBed, prometheus.GaugeValue,
							printerCurrentBedParsed, prusalink.GetLabels(s, job, "1")...)
						ch <- printerCurrentBed
					}

					printerCurrentNozzleParsed, e := strconv.ParseFloat(syslogData[s.Address]["curr_nozz"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerCurrentNozzle := prometheus.MustNewConstMetric(collector.printerCurrentNozzle, prometheus.GaugeValue,
							printerCurrentNozzleParsed, prusalink.GetLabels(s, job)...)
						ch <- printerCurrentNozzle
					}

					printerOvercurrentNozzleParsed, e := strconv.ParseFloat(syslogData[s.Address]["oc_nozz"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerOvercurrentNozzle := prometheus.MustNewConstMetric(collector.printerOvercurrentNozzle, prometheus.GaugeValue,
							printerOvercurrentNozzleParsed, prusalink.GetLabels(s, job)...)
						ch <- printerOvercurrentNozzle
					}

					printerOvercurrentInputParsed, e := strconv.ParseFloat(syslogData[s.Address]["oc_inp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerOvercurrentInput := prometheus.MustNewConstMetric(collector.printerOvercurrentInput, prometheus.GaugeValue,
							printerOvercurrentInputParsed, prusalink.GetLabels(s, job)...)
						ch <- printerOvercurrentInput
					}

					printerActiveExtruder, e := strconv.ParseFloat(syslogData[s.Address]["active_extruder"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerActiveExtruder := prometheus.MustNewConstMetric(collector.printerActiveExtruder, prometheus.GaugeValue,
							printerActiveExtruder, prusalink.GetLabels(s, job)...)
						ch <- printerActiveExtruder
					}

					printerDwarfMcuTemp, e := strconv.ParseFloat(syslogData[s.Address]["dwarf_mcu_temp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerDwarfMcuTemp := prometheus.MustNewConstMetric(collector.printerDwarfMcuTemp, prometheus.GaugeValue,
							printerDwarfMcuTemp, prusalink.GetLabels(s, job)...)
						ch <- printerDwarfMcuTemp
					}

					printerDwarfBoardTemp, e := strconv.ParseFloat(syslogData[s.Address]["dwarf_board_temp"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {

						printerDwarfBoardTemp := prometheus.MustNewConstMetric(collector.printerDwarfBoardTemp, prometheus.GaugeValue,
							printerDwarfBoardTemp, prusalink.GetLabels(s, job)...)
						ch <- printerDwarfBoardTemp
					}

					printerAxisZAdjustment, e := strconv.ParseFloat(syslogData[s.Address]["adj_z"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerAxisZAdjustment := prometheus.MustNewConstMetric(collector.printerAxisZAdjustment, prometheus.GaugeValue,
							printerAxisZAdjustment, prusalink.GetLabels(s, job)...)
						ch <- printerAxisZAdjustment
					}

					printerHeaterEnabled, e := strconv.ParseFloat(syslogData[s.Address]["heater_enabled"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerHeaterEnabled := prometheus.MustNewConstMetric(collector.printerHeaterEnabled, prometheus.GaugeValue,
							printerHeaterEnabled, prusalink.GetLabels(s, job)...)
						ch <- printerHeaterEnabled
					}

					printerLoadcellScale, e := strconv.ParseFloat(syslogData[s.Address]["loadcell_scale"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerLoadcellScale := prometheus.MustNewConstMetric(collector.printerLoadcellScale, prometheus.GaugeValue,
							printerLoadcellScale, prusalink.GetLabels(s, job)...)
						ch <- printerLoadcellScale
					}

					printerLoadcellThreshold, e := strconv.ParseFloat(syslogData[s.Address]["loadcell_threshold"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerLoadcellThreshold := prometheus.MustNewConstMetric(collector.printerLoadcellThreshold, prometheus.GaugeValue,
							printerLoadcellThreshold, prusalink.GetLabels(s, job)...)
						ch <- printerLoadcellThreshold
					}

					printerLoadcellHysteresis, e := strconv.ParseFloat(syslogData[s.Address]["loadcell_hysteresis"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerLoadcellHysteresis := prometheus.MustNewConstMetric(collector.printerLoadcellHysteresis, prometheus.GaugeValue,
							printerLoadcellHysteresis, prusalink.GetLabels(s, job)...)
						ch <- printerLoadcellHysteresis
					}

					if syslogData[s.Address]["buddy_revision"] != "" && syslogData[s.Address]["buddy_bom"] != "" {
						printerBuddySyslogInfo := prometheus.MustNewConstMetric(collector.printerBuddySyslogInfo, prometheus.GaugeValue,
							1, prusalink.GetLabels(s, job, syslogData[s.Address]["buddy_revision"], syslogData[s.Address]["buddy_bom"])...)
						ch <- printerBuddySyslogInfo
					}

					printerCPUUsage, e := strconv.ParseFloat(syslogData[s.Address]["cpu_usage"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerCPUUsage := prometheus.MustNewConstMetric(collector.printerCPUUsage, prometheus.GaugeValue,
							printerCPUUsage/100, prusalink.GetLabels(s, job)...)
						ch <- printerCPUUsage
					}
					////////////////////////////// PARSE! Heap ofc

					printerHeapTotal, e := strconv.ParseFloat(strings.Split(syslogData[s.Address]["heap"], ",")[1], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerHeapTotal := prometheus.MustNewConstMetric(collector.printerHeapTotal, prometheus.GaugeValue,
							printerHeapTotal, prusalink.GetLabels(s, job)...)
						ch <- printerHeapTotal
					}

					printerHeapFree, e := strconv.ParseFloat(strings.Split(syslogData[s.Address]["heap"], ",")[0], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerHeapFree := prometheus.MustNewConstMetric(collector.printerHeapFree, prometheus.GaugeValue,
							printerHeapFree, prusalink.GetLabels(s, job)...)
						ch <- printerHeapFree
					}

					printerPointsDropped, e := strconv.ParseFloat(syslogData[s.Address]["points_dropped"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {
						printerPointsDropped := prometheus.MustNewConstMetric(collector.printerPointsDropped, prometheus.GaugeValue,
							printerPointsDropped, prusalink.GetLabels(s, job)...)
						ch <- printerPointsDropped
					}

					printerMediaPrefetched, e := strconv.ParseFloat(syslogData[s.Address]["media_prefetched"], 32)
					if e != nil {
						log.Debug().Msg(e.Error())

					} else {

						printerMediaPrefetched := prometheus.MustNewConstMetric(collector.printerMediaPrefetched, prometheus.GaugeValue,
							printerMediaPrefetched, prusalink.GetLabels(s, job)...)
						ch <- printerMediaPrefetched
					}

				}
			}
		}

	}
}
