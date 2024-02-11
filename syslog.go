package main

import (
	"fmt"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/mcuadros/go-syslog.v2"
)

var syslogData = make(map[string]map[string]string)

func startSyslog(port int) { // yep i'll leave it in one function for now
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenUDP("0.0.0.0:" + fmt.Sprint(port))
	server.Boot()

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
			client_ip := logParts["client"].(string)
			if client_ip == "" { // Skip empty client ip
				continue
			} else {
				if syslogData[client_ip] == nil {
					syslogData[client_ip] = make(map[string]string)
				} // Initialize map for ip address if it doesn't exist - is it unique? No. Is it a problem? No. Is it experimental? Yes.

				syslogData[client_ip]["mac"] = logParts["hostname"].(string)

				for _, pattern := range patterns {
					reg, err := regexp.Compile(pattern.pattern)
					if err != nil {
						fmt.Println("Error compiling regexp:", err)
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

						syslogData[client_ip][match[1]] = fmt.Sprint(valueStr)
					}
				}
			}
		}
	}(channel)

	server.Wait()
}

type syslogCollector struct {
	// power metrics
	printerVolt5V            *prometheus.Desc
	printerVolt24V           *prometheus.Desc
	printerVoltBed           *prometheus.Desc
	printerVoltNozzle        *prometheus.Desc
	printerVoltSandwich5V    *prometheus.Desc
	printerVoltSplitter5V    *prometheus.Desc
	printerCurrentXlbuddy5V  *prometheus.Desc
	printerCurrentInput      *prometheus.Desc
	printerCurrentMMU        *prometheus.Desc
	printerCurrentBed        *prometheus.Desc
	printerCurrentNozzle     *prometheus.Desc
	printerOvercurrentNozzle *prometheus.Desc
	printerOvercurrentInput  *prometheus.Desc

	// printer metrics
	printerActiveExtruder     *prometheus.Desc
	printerDwarfMcuTemp       *prometheus.Desc
	printerDwarfBoardTemp     *prometheus.Desc
	printerAxisZAdjustment    *prometheus.Desc
	printerHeaterEnabled      *prometheus.Desc
	printerLoadcellScale      *prometheus.Desc
	printerLoadcellThreshold  *prometheus.Desc
	printerLoadcellHysteresis *prometheus.Desc

	// system metrics
	printerBuddyInfo       *prometheus.Desc // revision, bom
	printerCpuUsage        *prometheus.Desc
	printerHeapTotal       *prometheus.Desc
	printerHeapUsed        *prometheus.Desc
	printerPointsDropped   *prometheus.Desc
	printerMediaPrefetched *prometheus.Desc
}

func newSyslogCollector() *syslogCollector {
	defaultLabels := []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}
	return &syslogCollector{
		printerVolt5V: prometheus.NewDesc("prusa_buddy_voltage_5volts",
			"Voltage of 5 volt rail",
			defaultLabels,
			nil),
		printerVolt24V: prometheus.NewDesc("prusa_buddy_voltage_24volts",
			"Voltage of 24 volt rail",
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
			"Voltage of sandwich 5 volt rail",
			defaultLabels,
			nil),
		printerVoltSplitter5V: prometheus.NewDesc("prusa_buddy_voltage_splitter_5volts",
			"Voltage of splitter 5 volt rail",
			defaultLabels,
			nil),
		printerCurrentXlbuddy5V: prometheus.NewDesc("prusa_buddy_current_xlbuddy_5volts",
			"Current of XL buddy 5 volt rail",
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
		printerBuddyInfo: prometheus.NewDesc("prusa_buddy_info",
			"Buddy info",
			append(defaultLabels, "buddy_revision", "buddy_bom"),
			nil),
		printerCpuUsage: prometheus.NewDesc("prusa_buddy_cpu_usage",
			"CPU usage",
			defaultLabels,
			nil),
		printerHeapTotal: prometheus.NewDesc("prusa_buddy_heap_total",
			"Total heap",
			defaultLabels,
			nil),
		printerHeapUsed: prometheus.NewDesc("prusa_buddy_heap_used",
			"Used heap",
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
	ch <- collector.printerBuddyInfo
	ch <- collector.printerCpuUsage
	ch <- collector.printerHeapTotal
	ch <- collector.printerHeapUsed
	ch <- collector.printerPointsDropped
	ch <- collector.printerMediaPrefetched
}
