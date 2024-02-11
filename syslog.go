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

/* func newSyslogCollector() *syslogCollector {
	defaultLabels := []string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"}
	return &syslogCollector{
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
		printerPrintSpeed: prometheus.NewDesc("prusa_buddy_print_speed_ratio",
			"Current setting of printer speed in ratio (0.0-1.0)",
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
			"Returns information about current print time.",
			defaultLabels,
			nil),
		printerUp: prometheus.NewDesc("prusa_buddy_up",
			"Return information about online printers. If printer is registered as offline then returned value is 0.",
			[]string{"printer_address", "printer_model", "printer_name"},
			nil),
		printerNozzleSize: prometheus.NewDesc("prusa_buddy_nozzle_size",
			"Returns information about selected nozzle size.",
			defaultLabels,
			nil),
		printerStatus: prometheus.NewDesc("prusa_buddy_status",
			"Returns information status of printer.",
			append(defaultLabels, "printer_state"), // flags are defined by number :pug-dance:
			nil),
		printerAxisX: prometheus.NewDesc("prusa_buddy_axis_x",
			"Returns information about position of axis X.",
			defaultLabels,
			nil),
		printerAxisY: prometheus.NewDesc("prusa_buddy_axis_y",
			"Returns information about position of axis Y.",
			defaultLabels,
			nil),
		printerAxisZ: prometheus.NewDesc("prusa_buddy_axis_z",
			"Returns information about position of axis Z.",
			defaultLabels,
			nil),
		printerFlow: prometheus.NewDesc("prusa_buddy_print_flow_ratio",
			"Returns information about of filament flow in ratio (0.0 - 1.0).",
			defaultLabels,
			nil),
		printerInfo: prometheus.NewDesc("prusa_buddy_info",
			"Returns information about printer.",
			append(defaultLabels, "printer_serial", "printer_hostname"),
			nil),
		printerMMU: prometheus.NewDesc("prusa_buddy_mmu",
			"Returns information if MMU is enabled.",
			defaultLabels,
			nil),
		printerFanHotend: prometheus.NewDesc("prusa_buddy_fan_hotend",
			"Returns information about speed of hotend fan in rpm.",
			defaultLabels,
			nil),
		printerFanPrint: prometheus.NewDesc("prusa_buddy_fan_print",
			"Returns information about speed of print fan in rpm.",
			defaultLabels,
			nil),
	}
}

func (collector *syslogCollector) Describe(ch chan<- *prometheus.Desc) {
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
	ch <- collector.printerUp
	ch <- collector.printerNozzleSize
	ch <- collector.printerStatus
	ch <- collector.printerAxisX
	ch <- collector.printerAxisY
	ch <- collector.printerAxisZ
	ch <- collector.printerFlow
	ch <- collector.printerInfo
	ch <- collector.printerMMU
	ch <- collector.printerFanHotend
	ch <- collector.printerFanPrint
}
*/
