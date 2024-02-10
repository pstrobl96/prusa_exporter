package main

import (
	"fmt"
	"regexp"
	"strconv"

	"gopkg.in/mcuadros/go-syslog.v2"
)

var syslogData = make(map[string]map[string]string)

type Measurement struct {
	Name      string
	Value     float64
	Timestamp int64
}

// EXPERIMENTAL - I DID IT MYSELF - NOT TESTED
type Client struct {
	IP                string  // ip address
	MAC               string  // mac address
	CurrNozz          float64 // nozzle current - curr_nozz in SYSLOG
	CurrInp           float64 // input current - curr_inp in SYSLOG
	CurrMmu           float64 // mmu current - cur_mmu_imp
	HeapFree          int     // how much heap is free - heap_free in SYSLOG
	CpuUsage          int     // cpu use in percentage - cpu_usage in SYSLOG
	VoltBed           int     // bed voltage - volt_bed in SYSLOG
	HeaterEnabled     bool    // is the heater enabled - heater_enabled in SYSLOG
	ActiveExtruder    int     // active extruder -	active_extruder in SYSLOG
	BedCurrFirstRail  float64 // bed current - implemented only for one bed for now - bed_curr in SYSLOG
	BedCurrSecondRail float64 // bed current - implemented only for one bed for now - bed_curr in SYSLOG
	DwarfBoardTemp    int     // dwarf board temperature - dwarf_board_temp in SYSLOG
	DwarfMcuTemp      int     // dwarf mcu temperature - dwarf_mcu_temp in SYSLOG
	PointsDropped     int     // points dropped - points_dropped in SYSLOG
	Xlbuddy5VCurrent  float64 // xl buddy 5v current - xlbuddy5VCurrent in SYSLOG
	Sandwich5VCurrent float64 // sandwi(T)ch 5v current- Sandwitch5VCurrent in SYSLOG
	Splitter5VCurrent float64 // splitter 5v current - splitter_5V_current in SYSLOG
	Voltage5V         float64 // 5v voltage - named 5VVoltage in SYSLOG - Voltage_5V in SYSLOG
	Voltage24V        float64 // 24v voltage - named 24VVoltage in SYSLOG - Voltage_24V in SYSLOG
	OcNozz            int     // nozzle overcurrent / overclock / overcooked? - oc_nozz in SYSLOG
	OcInp             int     // input? overcurrent / overclock / overcooked? - oc_inp in SYSLOG
	BuddyBom          string  // version of buddy - buddy_bom in SYSLOG
}

func startSyslog(port int) {
	//clients := []Client{}
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
		{pattern: `(?P<name>\w+_[a-z]+) v=(?P<value>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}}, // integer

		//{pattern: `(?P<name>fan,fan=[\w,]+) (?P<value>[^=]+)=(?P<subvalue>[^=]+)=(?P<subvalue2>[^ ]+) (?P<timestamp>\d+)`, fields: []string{"name", "value", "subvalue", "subvalue2", "timestamp"}},
		// Add more patterns as needed
	}

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {

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

					value, err := strconv.ParseFloat(valueStr, 64)
					if err != nil {
						fmt.Printf("Error parsing value for %s: %v\n", match[1], err)
						continue
					}

					if syslogData[logParts["client"].(string)] == nil {
						syslogData[logParts["client"].(string)] = make(map[string]string)
					}

					syslogData[logParts["client"].(string)][match[1]] = fmt.Sprint(value)

				}
			}
			//time.Sleep(100 * time.Millisecond)
			//fmt.Println(logParts["client"])   // ip address
			//fmt.Println(logParts["hostname"]) // mac address
			//fmt.Println(logParts["message"])  // metrics
			//for data := range syslogData {
			//	fmt.Println(data)
			//}
		}
	}(channel)

	server.Wait()
}
