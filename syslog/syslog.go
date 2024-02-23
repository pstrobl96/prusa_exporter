package syslog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

var (
	syslogData = make(map[string]map[string]string)
	metrics    = map[string][]string{ // {"metric_name", "metric_help", "metric_label", "metric_type", "metric_regexp", {"label1", "label2"}} - TODO: merge metrics, add metric type and regexp
		"cpu_usage":                 {"prusa_cpu_usage", "Usage of printer CPU in percentage"},
		"heap":                      {"prusa_heap", "heap_free and heap_total in bytes"}, // not used
		"heap_free":                 {"prusa_heap_free", "Free heap in bytes"},
		"heap_total":                {"prusa_heap_total", "Total heap in bytes"},
		"crash":                     {"prusa_crash", "Crash count"},
		"crash_stat":                {"prusa_crash_stat", "Crash statistics"},
		"crash_repeated":            {"prusa_crash_repeated", "Crash repeated count"},
		"excite_freq":               {"prusa_excite_freq", "Excite frequency"},
		"freq_gain":                 {"prusa_freq_gain", "Frequency gain"},
		"tk_accel":                  {"prusa_tk_accel", "Acceleration"},
		"home_diff":                 {"prusa_home_diff", "Home difference"},
		"probe_z":                   {"prusa_probe_z", "Probe Z"},
		"probe_z_diff":              {"prusa_probe_z_diff", "Probe Z difference"},
		"probe_start":               {"prusa_probe_start", "Probe start"},
		"probe_analysis":            {"prusa_probe_analysis", "Probe analysis"},
		"heating_model_discrepancy": {"prusa_heat_model_disc", "Heating model discrepancy"},
		"esp_out":                   {"prusa_esp_out", "ESP output"},
		"esp_in":                    {"prusa_esp_in", "ESP input"},
		"fan_speed":                 {"prusa_fan_speed", "Fan speed"},
		"fan_hbr_speed":             {"prusa_fan_hbr_speed", "Heatbreak fan speed"},
		"ipos_x":                    {"prusa_ipos_x", "Stepper possition from startup in x axis"}, // merge with ipos_x, ipos_y and ipos_z
		"ipos_y":                    {"prusa_ipos_y", "Stepper possition from startup in y axis"}, // merge with ipos_x, ipos_y and ipos_z
		"ipos_z":                    {"prusa_ipos_z", "Stepper possition from startup in z axis"}, // merge with ipos_x, ipos_y and ipos_z
		"pos_x":                     {"prusa_pos_x", "Stepper possition in x axis"},               // merge into prusa_axis with label "axis"
		"pos_y":                     {"prusa_pos_y", "Stepper possition in y axis"},               // merge into prusa_axis with label "axis"
		"pos_z":                     {"prusa_pos_z", "Stepper possition in z axis"},               // merge into prusa_axis with label "axis"
		"adj_z":                     {"prusa_adj_z", "Adjustment of z axis"},
		"heater_enabled":            {"prusa_heater_enabled", "Heater enabled"},
		"volt_bed_raw":              {"prusa_voltage_bed_raw", "Voltage of bed raw"},     // merge into prusa_voltage_raw with label "voltage" and "rail"
		"volt_bed":                  {"prusa_voltage_bed", "Voltage of bed"},             // merge into prusa_voltage with label "voltage" and "rail"
		"volt_nozz_raw":             {"prusa_voltage_nozz_raw", "Voltage of nozzle raw"}, // merge into prusa_voltage_raw with label "voltage" and "rail"
		"volt_nozz":                 {"prusa_voltage_nozz", "Voltage of nozzle"},         // merge into prusa_voltage with label "voltage" and "rail"
		"curr_nozz_raw":             {"prusa_current_nozz_raw", "Current of nozzle raw"}, // merge into prusa_current_raw with label "current" and "rail"
		"curr_nozz":                 {"prusa_current_nozz", "Current of nozzle"},         // merge into prusa_current with label "current" and "rail"
		"curr_inp_raw":              {"prusa_current_inp_raw", "Current of input raw"},   // merge into prusa_current_raw with label "current" and "rail"
		"curr_inp":                  {"prusa_current_inp", "Current of input"},           // merge into prusa_current with label "current" and "rail"
		"cur_mmu_imp":               {"prusa_current_mmu_imp", "Current of MMU"},         // merge into prusa_mmu with label "mmu"
		"oc_nozz":                   {"prusa_oc_nozz", "Overcurrent of nozzle"},
		"oc_inp":                    {"prusa_oc_inp", "Overcurrent of input"},
		"splitter_5V_current":       {"prusa_current_splitter_5v", "Current of 5V splitter"}, // merge into prusa_current with label "current" and "rail"
		"24VVoltage":                {"prusa_voltage_24v", "Voltage of 24V rail"},            // merge into prusa_voltage with label "voltage" and "rail"
		"5VVoltage":                 {"prusa_voltage_5v", "Voltage of 5V rail"},              // merge into prusa_voltage with label "voltage" and "rail"
		"Sandwitch5VCurrent":        {"prusa_current_sandwich_5v", "Current of 5V sandwich"}, // merge into prusa_current with label "current" and "rail"
		"xlbuddy5VCurrent":          {"prusa_current_xlbuddy_5v", "Current of 5V xlbuddy"},   // merge into prusa_current with label "current" and "rail"
		"print_filename":            {"prusa_print_filename", "Name of the file being printed"},
		"dwarf_board_temp":          {"prusa_temp_dwarf_board", "Temperature of the board"},
		"dwarf_mcu_temp":            {"prusa_temp_dwarf_mcu", "Temperature of the Dwarf MCU"},
		"dwarfs_mcu_temp":           {"prusa_temp_dwarfs_mcu", "Temperature of the all Dwarfs MCUs"},
		"dwarfs_board_temp":         {"prusa_temp_dwarfs_board", "Temperature of the all Dwarfs boards"},
		"app_start":                 {"prusa_app_start", "Application start"}, // absolutelly no idea what this is
		"maintask_loop":             {"prusa_maintask_loop", "Main task loop"},
		"fsensor_raw":               {"prusa_fsensor_raw", "Raw value of the filament sensor"},
		"fsensor":                   {"prusa_fsensor", "Value of the filament sensor"},
		"side_fsensor_raw":          {"prusa_side_fsensor_raw", "Raw value of the side filament sensor"},
		"side_fsensor":              {"prusa_side_fsensor", "Value of the side filament sensor"},
		"nozzle_pwm":                {"prusa_nozzle_pwm", "Nozzle PWM"}, // merge into prusa_pwm with label "target" that will contain nozzle bed or bedlet
		"bed_pwm":                   {"prusa_bed_pwm", "Bed PWM"},       // merge into prusa_pwm with label "target" that will contain nozzle bed or bedlet
		"loadcell":                  {"prusa_loadcell", "Loadcell"},
		"loadcell_hp":               {"prusa_loadcell_hp", "Loadcell high precision"},
		"loadcell_xy":               {"prusa_loadcell_xy", "Loadcell XY"},
		"loadcell_age":              {"prusa_loadcell_age", "Loadcell age"},
		"loadcell_value":            {"prusa_loadcell_value", "Loadcell value"},
		"power_panic":               {"prusa_power_panic", "Power panic counter"},
		"crash_length":              {"prusa_crash_length", "Crash length"},
		"usbh_err_cnt":              {"prusa_usbh_err_cnt", "USBH error counter"},
		"media_prefetched":          {"prusa_media_prefetched", "Media prefetched"},
		"points_dropped":            {"prusa_points_dropped", "Points dropped"},
		"probe_window":              {"prusa_probe_window", "Probe window"},
		"eeprom_write":              {"prusa_eeprom_write", "EEPROM write"},
		"tmc_sg_x":                  {"prusa_tmc_sg_x", "Trinamic SG for axis X"},   // merge into prusa_tmc_sg with label "axis"
		"tmc_sg_y":                  {"prusa_tmc_sg_y", "Trinamic SG for axis Y"},   // merge into prusa_tmc_sg with label "axis"
		"tmc_sg_z":                  {"prusa_tmc_sg_z", "Trinamic SG for axis Z"},   // merge into prusa_tmc_sg with label "axis"
		"tmc_sg_e":                  {"prusa_tmc_sg_e", "Trinamic SG for extruder"}, // merge into prusa_tmc_sg with label "axis"
		"tmc_write":                 {"prusa_tmc_write", "Trinamic write"},
		"tmc_read":                  {"prusa_tmc_read", "Trinamic read"},
		"fan":                       {"prusa_fan", "Fan status"},
		"print_fan_act":             {"prusa_print_fan_act", "Print fan active"},
		"hbr_fan_act":               {"prusa_hbr_fan_act", "Heatbreak fan active"},
		"gui_loop_dur":              {"prusa_gui_loop_dur", "GUI loop duration"},
		"g425_cen":                  {"prusa_g425_cen", "Absolute tool center - an input for offset computation [mm]"},
		"g425_off":                  {"prusa_g425_off", "Tool offset relative to the first tool - result of the tool offset calibration [mm]"},
		"g425_rxy":                  {"prusa_g425_rxy", "Raw XY probe [mm]"},
		"g425_xy":                   {"prusa_g425_xy", "Verified XY probe - two raw probes agree on position [mm]"},
		"g425_rz":                   {"prusa_g425_rz", "Raw Z probe [mm]"},
		"g425_z":                    {"prusa_g425_z", "Averaged Z probe - N raw probes averaged [mm]"},
		"g425_xy_dev":               {"prusa_g425_xy_dev", "XY deviation - max difference between two raw probes [mm]"},
		"gcode":                     {"prusa_gcode", "Gcode"},
		"loadcell_scale":            {"prusa_loadcell_scale", "Loadcell scale"},
		"loadcell_threshold":        {"prusa_loadcell_threshold", "Loadcell threshold"},
		"loadcell_threshold_cont":   {"prusa_loadcell_threshold_cont", "Loadcell threshold continuous"},
		"loadcell_hysteresis":       {"prusa_loadcell_hysteresis", "Loadcell hysteresis"},
		"mmu_comm":                  {"prusa_mmu_comm", "MMU communication"},
		"dwarf_fast_refresh_delay":  {"prusa_dwarf_fast_refresh_delay", "Dwarf fast refresh delay"},
		"dwarf_picked_raw":          {"prusa_dwarf_picked_raw", "Dwarf picked raw"},
		"dwarf_parked_raw":          {"prusa_dwarf_parked_raw", "Dwarf parked raw"},
		"dwarf_heat_curr":           {"prusa_current_heater_dwarf", "Dwarf heater current"},
		"bed_state":                 {"prusa_state_bed", "Bed state"},
		"bed_curr":                  {"prusa_current_bed", "Bed current"},
		"bedlet_state":              {"prusa_state_bedlet", "Bedlet state"},
		"bedlet_temp":               {"prusa_temp_bedlet", "Bedlet temperature"},
		"bedlet_target":             {"prusa_temp_target_bedlet", "Bedlet target temperature"},
		"bedlet_pwm":                {"prusa_pwm_bedlet", "Bedlet PWM"}, // merge into prusa_pwm with label "target" that will contain nozzle bed or bedlet
		"bedlet_reg":                {"prusa_reg_bedlet", "Bedlet regulation"},
		"bedlet_curr":               {"prusa_current_bedlet", "Bedlet current"},
		"bed_mcu_temp":              {"prusa_temp_bed_mcu", "Bed MCU temperature"},
		"modbus_reqfail":            {"prusa_modbus_reqfail", "Modbus request fail"},
	}
)

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
