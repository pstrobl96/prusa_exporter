package syslog

import "github.com/prometheus/client_golang/prometheus"

type label struct {
	name  string
	value string
}

type collectorBranch struct {
	collector    *prometheus.Desc
	nameOfMetric string
	labels       []string
}

type overrideMapBranch struct {
	collectorName string
	labels        []string
}

var (
	defaultLabels = []string{"mac", "ip"}

	overrideMap = map[string][]overrideMapBranch{
		"cur_mmu_imp":         {{"current", []string{"mmu", ""}}},
		"Sandwitch5VCurrent":  {{"current", []string{"sandwich", "5v"}}},
		"volt_bed":            {{"voltage", []string{"bed", "24v"}}},
		"24VVoltage":          {{"voltage", []string{"printer", "24v"}}},
		"5VVoltage":           {{"voltage", []string{"printer", "5v"}}},
		"volt_bed_raw":        {{"voltage_raw", []string{"bed", "24v"}}},
		"volt_nozz_raw":       {{"voltage_raw", []string{"nozzle", "24v"}}},
		"dwarf_heat_curr":     {{"current", []string{"dwarf_heater", "24v"}}},
		"splitter_5V_current": {{"current", []string{"splitter", "5v"}}},
		"xlbuddy5VCurrent":    {{"current", []string{"buddy", "5v"}}},
		"curr_nozz_raw":       {{"current_raw", []string{"nozzle", "24v"}}},
		"curr_inp_raw":        {{"current_raw", []string{"input", "24v"}}},
		"print_fan_act":       {{"fan_active", []string{"print"}}},
		"hbr_fan_act":         {{"fan_active", []string{"heatbreak"}}},
		"pos_x":               {{"pos", []string{"x"}}},
		"pos_y":               {{"pos", []string{"y"}}},
		"pos_z":               {{"pos", []string{"z"}}},
		"ipos_x":              {{"ipos", []string{"x"}}},
		"ipos_y":              {{"ipos", []string{"y"}}},
		"ipos_z":              {{"ipos", []string{"z"}}},
	}

	// exceptionsMap is a map of exceptions for the metric names
	// key is the name from collectorMap and in array are labels
	overrideMapOld = map[string]map[string][]string{
		"cur_mmu_imp":         {"current": {"mmu", ""}},
		"Sandwitch5VCurrent":  {"current": {"sandwich", "5v"}},
		"volt_bed":            {"voltage": {"bed", "24v"}},
		"24VVoltage":          {"voltage": {"printer", "24v"}},
		"5VVoltage":           {"voltage": {"printer", "5v"}},
		"volt_bed_raw":        {"voltage_raw": {"bed", "24v"}},
		"volt_nozz_raw":       {"voltage_raw": {"nozzle", "24v"}},
		"dwarf_heat_curr":     {"current": {"dwarf_heater", "24v"}},
		"splitter_5V_current": {"current": {"splitter", "5v"}},
		"xlbuddy5VCurrent":    {"current": {"buddy", "5v"}},
		"curr_nozz_raw":       {"current_raw": {"nozzle", "24v"}},
		"curr_inp_raw":        {"current_raw": {"input", "24v"}},
		"print_fan_act":       {"fan_active": {"print"}},
		"hbr_fan_act":         {"fan_active": {"heatbreak"}},
		"pos_x":               {"pos": {"x"}},
		"pos_y":               {"pos": {"y"}},
		"pos_z":               {"pos": {"z"}},
		"ipos_x":              {"ipos": {"x"}},
		"ipos_y":              {"ipos": {"y"}},
		"ipos_z":              {"ipos": {"z"}},
	}

	collectorMap = map[string]collectorBranch{
		"active_extruder": {
			collector:    prometheus.NewDesc("prusa_active_extruder", "Active extruder", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"app_start": {
			collector:    prometheus.NewDesc("prusa_app_start", "Application start", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"axis_z_adjustment": {
			collector:    prometheus.NewDesc("prusa_axis_z_adjustment", "Axis Z adjustment", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"bedlet_regulation": {
			collector:    prometheus.NewDesc("prusa_bedlet_regulation", "Bedlet regulation", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"bedlet_state": {
			collector:    prometheus.NewDesc("prusa_bedlet_state", "Bedlet state", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"bed_state": {
			collector:    prometheus.NewDesc("prusa_bed_state", "Bed state", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"cpu_usage": {
			collector:    prometheus.NewDesc("prusa_cpu_usage_ratio", "CPU usage from 0.0 to 1.0", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"crash_counter": {
			collector:    prometheus.NewDesc("prusa_crash_counter", "Crash counter", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"crash_length": {
			collector:    prometheus.NewDesc("prusa_crash_length", "Crash length", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"crash_repeated_counter": {
			collector:    prometheus.NewDesc("prusa_crash_repeated_counter", "Crash repeated counter", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"crash_stat": {
			collector:    prometheus.NewDesc("prusa_crash_stat", "Crash statistics", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"current": {
			collector:    prometheus.NewDesc("prusa_current", "Current of different devices in / on the printer", append(defaultLabels, "device", "rail"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"current_raw": {
			collector:    prometheus.NewDesc("prusa_current_raw", "Current of different devices in / on the printer in raw sensor value", append(defaultLabels, "device", "rail"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"dwarf_fast_refresh_delay": {
			collector:    prometheus.NewDesc("prusa_dwarf_fast_refresh_delay", "Dwarf fast refresh delay", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"dwarf_parked_raw": {
			collector:    prometheus.NewDesc("prusa_dwarf_parked_raw", "Dwarf parked raw sensor value", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"dwarf_picked_raw": {
			collector:    prometheus.NewDesc("prusa_dwarf_picked_raw", "Dwarf picked raw sensor value", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"eeeprom_write": {
			collector:    prometheus.NewDesc("prusa_eeeprom_write", "Eeeprom write", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"excite_freq": {
			collector:    prometheus.NewDesc("prusa_excite_freq", "Excite frequency", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"fan_active": {
			collector:    prometheus.NewDesc("prusa_fan_active", "Fan active", append(defaultLabels, "fan"), nil),
			nameOfMetric: "v",
			labels:       []string{},
		},
		"fan_speed": {
			collector:    prometheus.NewDesc("prusa_syslog_fan_speed", "Fan", append(defaultLabels, "fan"), nil),
			nameOfMetric: "v",
			labels:       []string{},
		},
		"filename": {
			collector:    prometheus.NewDesc("prusa_filename", "Name of printed (b)gcode", append(defaultLabels, "file"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"fsensor": {
			collector:    prometheus.NewDesc("prusa_fsensor", "Filament Sensor", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"fsensor_raw": {
			collector:    prometheus.NewDesc("prusa_fsensor_raw", "Filament Sensor - raw sensor value", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"freq_gain": {
			collector:    prometheus.NewDesc("prusa_freq_gain", "Frequency gain", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"g425_cen": {
			collector:    prometheus.NewDesc("prusa_g425_cen", "Absolute tool center - an input for offset computation [mm]", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"g425_offset": {
			collector:    prometheus.NewDesc("prusa_g425_off", "Offset from the absolute tool center [mm]", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"g425_rxy": {
			collector:    prometheus.NewDesc("prusa_g425_rxy", "Raw XY probe [mm]", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"g425_rz": {
			collector:    prometheus.NewDesc("prusa_g425_rz", "Raw Z probe [mm]", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"g425_xy": {
			collector:    prometheus.NewDesc("prusa_g425_xy", "Verified XY probe - two raw probes agree on position [mm]", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"g425_z": {
			collector:    prometheus.NewDesc("prusa_g425_z", "Averaged Z probe - N raw probes averaged [mm]", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"gcode": {
			collector:    prometheus.NewDesc("prusa_gcode", "Printed GCode", append(defaultLabels, "gcode"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"gui_loop_duration": {
			collector:    prometheus.NewDesc("prusa_gui_loop_duration", "Gui loop duration", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"heap_free": {
			collector:    prometheus.NewDesc("prusa_heap_free", "Free heap", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"heap_total": {
			collector:    prometheus.NewDesc("prusa_heap_total", "Total heap", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"heat_model_discard": {
			collector:    prometheus.NewDesc("prusa_heat_model_disc", "Heating model discrepancy", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"heater_enabled": {
			collector:    prometheus.NewDesc("prusa_heater_enabled", "Heater enabled", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"home_diff": {
			collector:    prometheus.NewDesc("prusa_home_diff", "Home difference", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"ipos": {
			collector:    prometheus.NewDesc("prusa_stepper_ipos", "Stepper possition from startup", append(defaultLabels, "axis"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"loadcell_hysteresis": {
			collector:    prometheus.NewDesc("prusa_loadcell_hysteresis", "Loadcell hysteresis", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"loadcell_scale": {
			collector:    prometheus.NewDesc("prusa_loadcell_scale", "Loadcell scale", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"loadcell_threshold": {
			collector:    prometheus.NewDesc("prusa_loadcell_threshold", "Loadcell threshold", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"loadcell_threshold_cont": {
			collector:    prometheus.NewDesc("prusa_loadcell_threshold_cont", "Loadcell threshold continuous", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"loadcell_value": {
			collector:    prometheus.NewDesc("prusa_loadcell", "Value from loadcell sensor", defaultLabels, nil),
			nameOfMetric: "r",
			labels:       []string{},
		},
		"loadcell_xy": {
			collector:    prometheus.NewDesc("prusa_loadcell_xy", "Loadcell XY", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"maintask_loop": {
			collector:    prometheus.NewDesc("prusa_maintask_loop", "Maintask loop", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"media_prefetched": {
			collector:    prometheus.NewDesc("prusa_eeeprom_write", "Eeeprom write", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"mmu_comm": {
			collector:    prometheus.NewDesc("prusa_mmu_comm", "MMU communication", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"modbus_reqfail": {
			collector:    prometheus.NewDesc("prusa_modbus_reqfail", "Modbus request fail", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"network_in": {
			collector:    prometheus.NewDesc("prusa_network_in", "Network in", append(defaultLabels, "device"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"network_out": {
			collector:    prometheus.NewDesc("prusa_network_out", "Network out", append(defaultLabels, "device"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"overcurrent": {
			collector:    prometheus.NewDesc("prusa_overcurrent", "Overcurrent of different devices in / on the printer", append(defaultLabels, "device"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"points_dropped": {
			collector:    prometheus.NewDesc("prusa_points_dropped", "Points dropped", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"pos": {
			collector:    prometheus.NewDesc("prusa_stepper_pos", "Stepper possition", append(defaultLabels, "axis"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"power_panic_count": {
			collector:    prometheus.NewDesc("prusa_power_panic_count", "Power panic triggered", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"probe_analysis": {
			collector:    prometheus.NewDesc("prusa_probe_analysis", "Probe analysis", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"probe_info": {
			collector:    prometheus.NewDesc("prusa_probe_info", "Probe info", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"probe_start": {
			collector:    prometheus.NewDesc("prusa_probe_start", "Probe start", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"probe_z": {
			collector:    prometheus.NewDesc("prusa_probe_z", "Probe Z", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"probe_z_diff": {
			collector:    prometheus.NewDesc("prusa_probe_z_diff", "Probe Z difference", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"pwm": {
			collector:    prometheus.NewDesc("prusa_pwm", "PWM value of nozzle and bed mostly", append(defaultLabels, "device"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"side_fsensor": {
			collector:    prometheus.NewDesc("prusa_side_fsensor", "Side Filament Sensor", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"side_fsensor_raw": {
			collector:    prometheus.NewDesc("prusa_side_fsensor_raw", "Side Filament Sensor - raw sensor value", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"syslog_info": {
			collector:    prometheus.NewDesc("prusa_syslog_info", "Buddy syslog info", append(defaultLabels, "revision", "bom"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"tmc_read": {
			collector:    prometheus.NewDesc("prusa_tmc_read", "Trinamic read", append(defaultLabels, "axis"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"tmc_sg": {
			collector:    prometheus.NewDesc("prusa_tmc_sg", "Trinamic SG", append(defaultLabels, "axis"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"tmc_write": {
			collector:    prometheus.NewDesc("prusa_tmc_write", "Trinamic write", append(defaultLabels, "axis"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"tk_acceleration": {
			collector:    prometheus.NewDesc("prusa_tk_acceleration", "TK acceleration", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"temp": {
			collector:    prometheus.NewDesc("prusa_temp", "Temperature of different devices in / on the printer", append(defaultLabels, "device"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"usbh_err_count": {
			collector:    prometheus.NewDesc("prusa_usbh_err_count", "USBH error count", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"voltage": {
			collector:    prometheus.NewDesc("prusa_voltage", "Voltage of different devices in / on the printer", append(defaultLabels, "device", "rail"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"voltage_raw": {
			collector:    prometheus.NewDesc("prusa_voltage_raw", "Voltage of different devices in / on the printer in raw sensor value", append(defaultLabels, "device", "rail"), nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
		"xy_dev": {
			collector:    prometheus.NewDesc("prusa_xy_dev", "XY deviation - max difference between two raw probes [mm]", defaultLabels, nil),
			nameOfMetric: "value",
			labels:       []string{},
		},
	}
)
