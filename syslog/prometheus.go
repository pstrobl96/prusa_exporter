package syslog

import (
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

func getLabels(mac string, ip string, port string, labels []string, labelValues ...string) []string {
	labelValues = append(labelValues, labels...)
	return append([]string{mac, ip, port}, labelValues...)
}

func getNumberOf(s string) (int, string, error) {
	splitted := strings.Split(s, "_")
	if len(splitted) == 0 {
		return 0, s, nil
	}
	indexOfLast := len(splitted) - 1

	if num, err := strconv.Atoi(splitted[indexOfLast]); err == nil {
		return num, strings.Join(splitted[:indexOfLast-1], "_"), nil
	}

	return -1, s, nil
}

// Collector is a struct that defines all the syslog metrics
type Collector struct {
	printerActiveExtruder        *prometheus.Desc
	printerAppStart              *prometheus.Desc
	printerAxisZAdjustment       *prometheus.Desc
	printerBedletRegulationD     *prometheus.Desc // bedlet_regulation_d
	printerBedletRegulationI     *prometheus.Desc // bedlet_regulation_i
	printerBedletRegulationP     *prometheus.Desc // bedlet_regulation_p
	printerBedletRegulationTc    *prometheus.Desc // bedlet_regulation_tc
	printerBedletState           *prometheus.Desc // bedlet_state
	printerBedState              *prometheus.Desc
	printerBuddyBom              *prometheus.Desc
	printerBuddyRevision         *prometheus.Desc
	printerBuddyFW               *prometheus.Desc
	printerCPUUsage              *prometheus.Desc
	printerCrashCounter          *prometheus.Desc
	printerCrashLength           *prometheus.Desc
	printerCrashRepeatedCounter  *prometheus.Desc
	printerCrashStat             *prometheus.Desc
	printerCurrent               *prometheus.Desc
	printerCurrentRaw            *prometheus.Desc
	printerDwarfFastRefreshDelay *prometheus.Desc
	printerDwarfParkedRaw        *prometheus.Desc
	printerDwarfPickedRaw        *prometheus.Desc
	printerEeepromWrite          *prometheus.Desc
	printerExciteFreq            *prometheus.Desc
	printerFanActive             *prometheus.Desc
	printerFanSpeed              *prometheus.Desc
	printerFilename              *prometheus.Desc
	printerFilament              *prometheus.Desc
	printerFSensor               *prometheus.Desc
	printerFSensorRaw            *prometheus.Desc
	printerFreqGain              *prometheus.Desc
	printerG425Cen               *prometheus.Desc
	printerG425Offset            *prometheus.Desc
	printerG425Rxy               *prometheus.Desc
	printerG425Rz                *prometheus.Desc
	printerG425Xy                *prometheus.Desc
	printerG425Z                 *prometheus.Desc
	printerGcode                 *prometheus.Desc
	printerGuiLoopDuration       *prometheus.Desc
	printerHeapFree              *prometheus.Desc
	printerHeapTotal             *prometheus.Desc
	printerHeatModelDiscard      *prometheus.Desc
	printerHeaterEnabled         *prometheus.Desc
	printerHomeDiffAx            *prometheus.Desc
	printerHomeDiffOk            *prometheus.Desc
	printerHomeDiffValue         *prometheus.Desc
	printerIpos                  *prometheus.Desc
	printerLoadcellAge           *prometheus.Desc
	printerLoadcellHp            *prometheus.Desc
	printerLoadcellHysteresis    *prometheus.Desc
	printerLoadcellScale         *prometheus.Desc
	printerLoadcellThreshold     *prometheus.Desc
	printerLoadcellThresholdCont *prometheus.Desc
	printerLoadcellValue         *prometheus.Desc
	printerLoadcellValueRaw      *prometheus.Desc
	printerLoadcellXY            *prometheus.Desc
	printerMaintaskLoop          *prometheus.Desc
	printerMediaPrefetched       *prometheus.Desc
	printerMMUComm               *prometheus.Desc
	printerModbusReqfail         *prometheus.Desc
	printerNetworkIn             *prometheus.Desc
	printerNetworkOut            *prometheus.Desc
	printerOvercurrent           *prometheus.Desc
	printerPointsDropped         *prometheus.Desc
	printerPos                   *prometheus.Desc
	printerPowerPanicCount       *prometheus.Desc
	printerPrinting              *prometheus.Desc
	printerProbeAnalysis         *prometheus.Desc
	printerProbeInfo             *prometheus.Desc
	printerProbeStart            *prometheus.Desc
	printerProbeZ                *prometheus.Desc // probe_z
	printerProbeZDiff            *prometheus.Desc
	printerPwm                   *prometheus.Desc
	printerSideFSensor           *prometheus.Desc // side_fsensor
	printerSideFSensorRaw        *prometheus.Desc
	printerSyslogInfo            *prometheus.Desc // revision, bom
	printerTmcRead               *prometheus.Desc
	printerTmcSg                 *prometheus.Desc
	printerTmcWrite              *prometheus.Desc
	printerTKAcceleration        *prometheus.Desc
	printerTemp                  *prometheus.Desc
	printerTempTarget            *prometheus.Desc
	printerUsbhErrCount          *prometheus.Desc
	printerVoltage               *prometheus.Desc
	printerVoltageRaw            *prometheus.Desc
	printerXyDev                 *prometheus.Desc
}

// NewCollector is a function that returns new Collector
// NewCollector creates a new instance of the Collector struct with the provided configuration.
// It initializes all the Prometheus metrics used for monitoring different aspects of the printer.
// The defaultLabels parameter is a list of labels that will be included in all the metrics.
// Returns a pointer to the created Collector.
func NewCollector() *Collector {
	defaultLabels := []string{"mac", "ip", "port"}

	return &Collector{
		printerActiveExtruder:        prometheus.NewDesc("prusa_active_extruder", "Active extruder - used for XL", defaultLabels, nil),
		printerAppStart:              prometheus.NewDesc("prusa_app_start", "Application start", defaultLabels, nil),
		printerAxisZAdjustment:       prometheus.NewDesc("prusa_axis_z_adjustment", "Axis Z adjustment", defaultLabels, nil),
		printerBedletRegulationD:     prometheus.NewDesc("prusa_bedlet_regulation_d", "Bedlet regulation d value", append(defaultLabels, "bedlet"), nil),
		printerBedletRegulationI:     prometheus.NewDesc("prusa_bedlet_regulation_i", "Bedlet regulation i value", append(defaultLabels, "bedlet"), nil),
		printerBedletRegulationP:     prometheus.NewDesc("prusa_bedlet_regulation_p", "Bedlet regulation p value", append(defaultLabels, "bedlet"), nil),
		printerBedletRegulationTc:    prometheus.NewDesc("prusa_bedlet_regulation_tc", "Bedlet regulation tc value", append(defaultLabels, "bedlet"), nil),
		printerBedletState:           prometheus.NewDesc("prusa_bedlet_state", "Bedlet state", append(defaultLabels, "bedlet"), nil),
		printerBedState:              prometheus.NewDesc("prusa_bed_state", "Bed state", defaultLabels, nil),
		printerBuddyBom:              prometheus.NewDesc("prusa_buddy_bom", "Buddy bom", defaultLabels, nil),
		printerBuddyRevision:         prometheus.NewDesc("prusa_buddy_revision", "Buddy revision", defaultLabels, nil),
		printerBuddyFW:               prometheus.NewDesc("prusa_buddy_fw", "Buddy firmware version", append(defaultLabels, "version"), nil),
		printerCPUUsage:              prometheus.NewDesc("prusa_cpu_usage_ratio", "CPU usage from 0.0 to 1.0", defaultLabels, nil),
		printerCrashCounter:          prometheus.NewDesc("prusa_crash_counter", "Crash counter", defaultLabels, nil),
		printerCrashLength:           prometheus.NewDesc("prusa_crash_length", "Crash length", defaultLabels, nil),
		printerCrashRepeatedCounter:  prometheus.NewDesc("prusa_crash_repeated_counter", "Crash repeated counter", defaultLabels, nil),
		printerCrashStat:             prometheus.NewDesc("prusa_crash_stat", "Crash statistics", defaultLabels, nil),
		printerCurrent:               prometheus.NewDesc("prusa_current", "Current of different devices in / on the printer", append(defaultLabels, "rail", "device"), nil),
		printerCurrentRaw:            prometheus.NewDesc("prusa_current_raw", "Current of different devices in / on the printer in raw sensor value", append(defaultLabels, "rail", "device"), nil),
		printerDwarfFastRefreshDelay: prometheus.NewDesc("prusa_dwarf_fast_refresh_delay", "Dwarf fast refresh delay", defaultLabels, nil),
		printerDwarfParkedRaw:        prometheus.NewDesc("prusa_dwarf_parked_raw", "Dwarf parked raw sensor value", append(defaultLabels, "tool"), nil),
		printerDwarfPickedRaw:        prometheus.NewDesc("prusa_dwarf_picked_raw", "Dwarf picked raw sensor value", append(defaultLabels, "tool"), nil),
		printerEeepromWrite:          prometheus.NewDesc("prusa_eeeprom_write", "Eeeprom write", defaultLabels, nil),
		printerExciteFreq:            prometheus.NewDesc("prusa_excite_freq", "Excite frequency", defaultLabels, nil),
		printerFanActive:             prometheus.NewDesc("prusa_fan_active", "Fan active", append(defaultLabels, "fan"), nil),
		printerFanSpeed:              prometheus.NewDesc("prusa_fan_speed_ratio", "Fan", append(defaultLabels, "fan"), nil),
		printerFilename:              prometheus.NewDesc("prusa_filename", "Name of printed (b)gcode", append(defaultLabels, "file"), nil),
		printerFilament:              prometheus.NewDesc("prusa_filament", "Name of printed (b)gcode", append(defaultLabels, "filament"), nil),
		printerFSensor:               prometheus.NewDesc("prusa_fsensor", "Filament Sensor", defaultLabels, nil),
		printerFSensorRaw:            prometheus.NewDesc("prusa_fsensor_raw", "Filament Sensor - raw sensor value", append(defaultLabels, "sensor"), nil),
		printerFreqGain:              prometheus.NewDesc("prusa_freq_gain", "Frequency gain", defaultLabels, nil),
		printerG425Cen:               prometheus.NewDesc("prusa_g425_cen", "Absolute tool center - an input for offset computation [mm]", defaultLabels, nil),
		printerG425Offset:            prometheus.NewDesc("prusa_g425_off", "Offset from the absolute tool center [mm]", defaultLabels, nil),
		printerG425Rxy:               prometheus.NewDesc("prusa_g425_rxy", "Raw XY probe [mm]", defaultLabels, nil),
		printerG425Rz:                prometheus.NewDesc("prusa_g425_rz", "Raw Z probe [mm]", defaultLabels, nil),
		printerG425Xy:                prometheus.NewDesc("prusa_g425_xy", "Verified XY probe - two raw probes agree on position [mm]", defaultLabels, nil),
		printerG425Z:                 prometheus.NewDesc("prusa_g425_z", "Averaged Z probe - N raw probes averaged [mm]", defaultLabels, nil),
		printerGcode:                 prometheus.NewDesc("prusa_gcode", "Printed GCode", append(defaultLabels, "gcode"), nil),
		printerGuiLoopDuration:       prometheus.NewDesc("prusa_gui_loop_duration", "Gui loop duration", defaultLabels, nil),
		printerHeapFree:              prometheus.NewDesc("prusa_heap_free", "Free heap", defaultLabels, nil),
		printerHeapTotal:             prometheus.NewDesc("prusa_heap_total", "Total heap", defaultLabels, nil),
		printerHeatModelDiscard:      prometheus.NewDesc("prusa_heat_model_disc", "Heating model discrepancy", defaultLabels, nil),
		printerHeaterEnabled:         prometheus.NewDesc("prusa_heater_enabled", "Heater enabled", defaultLabels, nil),
		printerHomeDiffAx:            prometheus.NewDesc("prusa_home_diff_ax", "Home diff ax", append(defaultLabels, "tool"), nil),
		printerHomeDiffOk:            prometheus.NewDesc("prusa_home_diff_ok", "Home diff ok", append(defaultLabels, "tool"), nil),
		printerHomeDiffValue:         prometheus.NewDesc("prusa_home_diff_value", "Home diff value", append(defaultLabels, "tool"), nil),
		printerIpos:                  prometheus.NewDesc("prusa_stepper_ipos", "Stepper possition from startup", append(defaultLabels, "axis"), nil),
		printerLoadcellAge:           prometheus.NewDesc("prusa_loadcell_age", "Loadcell age", defaultLabels, nil),
		printerLoadcellHysteresis:    prometheus.NewDesc("prusa_loadcell_hysteresis", "Loadcell hysteresis", defaultLabels, nil),
		printerLoadcellHp:            prometheus.NewDesc("prusa_loadcell_hp", "Loadcell filtered z load", defaultLabels, nil),
		printerLoadcellScale:         prometheus.NewDesc("prusa_loadcell_scale", "Loadcell scale", defaultLabels, nil),
		printerLoadcellThreshold:     prometheus.NewDesc("prusa_loadcell_threshold", "Loadcell threshold", defaultLabels, nil),
		printerLoadcellThresholdCont: prometheus.NewDesc("prusa_loadcell_threshold_cont", "Loadcell threshold continuous", defaultLabels, nil),
		printerLoadcellValue:         prometheus.NewDesc("prusa_loadcell", "Value from loadcell sensor", defaultLabels, nil),
		printerLoadcellValueRaw:      prometheus.NewDesc("prusa_loadcell_raw", "Value from loadcell sensor in raw sensor value", defaultLabels, nil),
		printerLoadcellXY:            prometheus.NewDesc("prusa_loadcell_xy", "Loadcell XY", defaultLabels, nil),
		printerMaintaskLoop:          prometheus.NewDesc("prusa_maintask_loop", "Maintask loop", defaultLabels, nil),
		printerMediaPrefetched:       prometheus.NewDesc("prusa_eeeprom_write", "Eeeprom write", defaultLabels, nil),
		printerMMUComm:               prometheus.NewDesc("prusa_mmu_comm", "MMU communication", defaultLabels, nil),
		printerModbusReqfail:         prometheus.NewDesc("prusa_modbus_reqfail", "Modbus request fail", defaultLabels, nil),
		printerNetworkIn:             prometheus.NewDesc("prusa_network_in", "Network in", append(defaultLabels, "device"), nil),
		printerNetworkOut:            prometheus.NewDesc("prusa_network_out", "Network out", append(defaultLabels, "device"), nil),
		printerOvercurrent:           prometheus.NewDesc("prusa_overcurrent", "Overcurrent of different devices in / on the printer", append(defaultLabels, "device"), nil),
		printerPrinting:              prometheus.NewDesc("prusa_printing", "Printing printer", defaultLabels, nil),
		printerPointsDropped:         prometheus.NewDesc("prusa_points_dropped", "Points dropped", defaultLabels, nil),
		printerPos:                   prometheus.NewDesc("prusa_stepper_pos", "Stepper possition", append(defaultLabels, "axis"), nil),
		printerPowerPanicCount:       prometheus.NewDesc("prusa_power_panic_count", "Power panic triggered", defaultLabels, nil),
		printerProbeAnalysis:         prometheus.NewDesc("prusa_probe_analysis", "Probe analysis", defaultLabels, nil),
		printerProbeInfo:             prometheus.NewDesc("prusa_probe_info", "Probe info", defaultLabels, nil),
		printerProbeStart:            prometheus.NewDesc("prusa_probe_start", "Probe start", defaultLabels, nil),
		printerProbeZ:                prometheus.NewDesc("prusa_probe_z", "Probe Z", defaultLabels, nil),
		printerProbeZDiff:            prometheus.NewDesc("prusa_probe_z_diff", "Probe Z difference", defaultLabels, nil),
		printerPwm:                   prometheus.NewDesc("prusa_pwm", "PWM value of nozzle and bed mostly", append(defaultLabels, "device"), nil),
		printerSideFSensor:           prometheus.NewDesc("prusa_side_fsensor", "Side Filament Sensor", defaultLabels, nil),
		printerSideFSensorRaw:        prometheus.NewDesc("prusa_side_fsensor_raw", "Side Filament Sensor - raw sensor value", append(defaultLabels, "sensor"), nil),
		printerSyslogInfo:            prometheus.NewDesc("prusa_syslog_info", "Buddy syslog info", append(defaultLabels, "revision", "bom"), nil),
		printerTmcRead:               prometheus.NewDesc("prusa_tmc_read", "Trinamic read", append(defaultLabels, "axis"), nil),
		printerTmcSg:                 prometheus.NewDesc("prusa_tmc_sg", "Trinamic SG", append(defaultLabels, "axis"), nil),
		printerTmcWrite:              prometheus.NewDesc("prusa_tmc_write", "Trinamic write", append(defaultLabels, "axis"), nil),
		printerTKAcceleration:        prometheus.NewDesc("prusa_tk_acceleration", "TK acceleration", defaultLabels, nil),
		printerTemp:                  prometheus.NewDesc("prusa_temp", "Temperature of different devices in / on the printer", append(defaultLabels, "device"), nil),
		printerTempTarget:            prometheus.NewDesc("prusa_temp_target", "Target temperature of different devices in / on the printer", append(defaultLabels, "device"), nil),
		printerUsbhErrCount:          prometheus.NewDesc("prusa_usbh_err_count", "USBH error counter", defaultLabels, nil),
		printerVoltage:               prometheus.NewDesc("prusa_voltage", "Voltage of different devices in / on the printer", append(defaultLabels, "rail", "device"), nil),
		printerVoltageRaw:            prometheus.NewDesc("prusa_voltage_raw", "Voltage of different devices in / on the printer in raw sensor value", append(defaultLabels, "rail", "device"), nil),
		printerXyDev:                 prometheus.NewDesc("prusa_xy_dev", "XY deviation - max difference between two raw probes [mm]", defaultLabels, nil),
	}
}

// Describe is a function that describes all the metrics
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.printerActiveExtruder
	ch <- collector.printerAppStart
	ch <- collector.printerAxisZAdjustment
	ch <- collector.printerBedletRegulationD
	ch <- collector.printerBedletRegulationI
	ch <- collector.printerBedletRegulationP
	ch <- collector.printerBedletRegulationTc
	ch <- collector.printerBedletState
	ch <- collector.printerBedState
	ch <- collector.printerBuddyBom
	ch <- collector.printerBuddyRevision
	ch <- collector.printerBuddyFW
	ch <- collector.printerCPUUsage
	ch <- collector.printerCrashCounter
	ch <- collector.printerCrashLength
	ch <- collector.printerCrashRepeatedCounter
	ch <- collector.printerCrashStat
	ch <- collector.printerCurrent
	ch <- collector.printerCurrentRaw
	ch <- collector.printerDwarfFastRefreshDelay
	ch <- collector.printerDwarfParkedRaw
	ch <- collector.printerDwarfPickedRaw
	ch <- collector.printerEeepromWrite
	ch <- collector.printerExciteFreq
	ch <- collector.printerFanActive
	ch <- collector.printerFanSpeed
	ch <- collector.printerFilament
	ch <- collector.printerFilename
	ch <- collector.printerFSensor
	ch <- collector.printerFSensorRaw
	ch <- collector.printerFreqGain
	ch <- collector.printerG425Cen
	ch <- collector.printerG425Offset
	ch <- collector.printerG425Rxy
	ch <- collector.printerG425Rz
	ch <- collector.printerG425Xy
	ch <- collector.printerG425Z
	ch <- collector.printerGcode
	ch <- collector.printerGuiLoopDuration
	ch <- collector.printerHeapFree
	ch <- collector.printerHeapTotal
	ch <- collector.printerHeatModelDiscard
	ch <- collector.printerHeaterEnabled
	ch <- collector.printerHomeDiffAx
	ch <- collector.printerHomeDiffOk
	ch <- collector.printerHomeDiffValue
	ch <- collector.printerIpos
	ch <- collector.printerLoadcellAge
	ch <- collector.printerLoadcellHp
	ch <- collector.printerLoadcellHysteresis
	ch <- collector.printerLoadcellScale
	ch <- collector.printerLoadcellThreshold
	ch <- collector.printerLoadcellThresholdCont
	ch <- collector.printerLoadcellValue
	ch <- collector.printerLoadcellValueRaw
	ch <- collector.printerMaintaskLoop
	ch <- collector.printerMediaPrefetched
	ch <- collector.printerMMUComm
	ch <- collector.printerModbusReqfail
	ch <- collector.printerNetworkIn
	ch <- collector.printerNetworkOut
	ch <- collector.printerOvercurrent
	ch <- collector.printerPrinting
	ch <- collector.printerPointsDropped
	ch <- collector.printerPos
	ch <- collector.printerPowerPanicCount
	ch <- collector.printerProbeAnalysis
	ch <- collector.printerProbeInfo
	ch <- collector.printerProbeStart
	ch <- collector.printerProbeZ
	ch <- collector.printerProbeZDiff
	ch <- collector.printerPwm
	ch <- collector.printerSideFSensor
	ch <- collector.printerSideFSensorRaw
	ch <- collector.printerTmcRead
	ch <- collector.printerTmcSg
	ch <- collector.printerTmcWrite
	ch <- collector.printerTKAcceleration
	ch <- collector.printerTemp
	ch <- collector.printerTempTarget
	ch <- collector.printerUsbhErrCount
	ch <- collector.printerVoltage
	ch <- collector.printerVoltageRaw
	ch <- collector.printerXyDev
}

// Collect is a function that collects all the metrics
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {

	// little bit more memory intensive but we need to extract the data from the map as fast as possible
	var syslogMetricsExtracted = make(map[string]map[string]map[string]string)

	syslogMetrics.Range(func(key, value interface{}) bool {
		mac := key.(string)
		nestedmap, ok := value.(map[string]map[string]string)

		log.Trace().Msg("Collecting metrics for " + mac)
		log.Trace().Msg("nestedmap: " + nestedmap["ip"]["value"])

		if !ok {
			log.Error().Msg("Error casting syslog data")
			return false
		}

		syslogMetricsExtracted[mac] = nestedmap
		return true
	})

	for mac, nestedmap := range syslogMetricsExtracted {
		ipArr := strings.Split(nestedmap["ip"]["value"], ":")
		ip := ipArr[0]
		port := ipArr[1]
		for k, v := range nestedmap {
			var (
				collectorItem *prometheus.Desc
				labels        = []string{}
				suffix        string
				valueParsed   float64
				valueKey      = "value" // mostly its value
			)

			length, name, err := getNumberOf(k)

			if err != nil {
				log.Error().Msgf("Error parsing metric name %s: %s", k, err)
				continue // Skip to next iteration if metric name parsing fails
			}

			if length != -1 {
				k = name
				suffix = "_" + strconv.Itoa(length)
			}

			splittedName := strings.Split(k, "_")

			switch k {
			case "temp_hbr":
				fallthrough
			case "temp_brd":
				fallthrough
			case "temp_chamber":
				fallthrough
			case "temp_mcu":
				fallthrough
			case "temp_sandwich":
				fallthrough
			case "temp_splitter":
				fallthrough
			case "temp_bed":
				fallthrough
			case "temp_noz":
				labels = []string{splittedName[1] + suffix}
				collectorItem = collector.printerTemp
			case "bedlet_temp":
				labels = []string{splittedName[0] + suffix}
				collectorItem = collector.printerTemp
			case "dwarf_board_temp":
				fallthrough
			case "dwarf_mcu_temp":
				fallthrough
			case "dwarfs_mcu_temp":
				fallthrough
			case "dwarfs_board_temp":
				fallthrough
			case "bed_mcu_temp":
				collectorItem = collector.printerTemp
				labels = []string{splittedName[0] + "_" + splittedName[1] + suffix}
			case "pos_x":
				fallthrough
			case "pos_y":
				fallthrough
			case "pos_z":
				labels = []string{splittedName[1] + suffix}
				collectorItem = collector.printerPos
			case "ipos_x":
				fallthrough
			case "ipos_y":
				fallthrough
			case "ipos_z":
				labels = []string{splittedName[1] + suffix}
				collectorItem = collector.printerIpos
			case "esp_out":
				fallthrough
			case "eth_out":
				valueKey = "sent"
				labels = []string{splittedName[0]}
				collectorItem = collector.printerNetworkOut
			case "esp_in":
				fallthrough
			case "eth_in":
				valueKey = "recv"
				labels = []string{splittedName[0]}
				collectorItem = collector.printerNetworkIn
			case "24VVoltage":
				fallthrough
			case "5VVoltage":
				labels = []string{strings.ReplaceAll(name, "Voltage", ""), ""}
				collectorItem = collector.printerVoltage
			case "volt_bed":
				fallthrough
			case "volt_nozz":
				labels = []string{"", splittedName[1] + suffix}
				collectorItem = collector.printerVoltage
			case "ttemp_bed":
				fallthrough
			case "ttemp_noz":
				fallthrough
			case "bedlet_target":
				labels = []string{splittedName[1] + suffix}
				collectorItem = collector.printerTempTarget
			case "Sandwitch5VCurrent":
				labels = []string{"5V", "sandwich"}
				collectorItem = collector.printerCurrent
			case "xlbuddy5VCurrent":
				labels = []string{"5V", "xlBuddy"}
				collectorItem = collector.printerCurrent
			case "splitter_5V_current":
				labels = []string{"5V", "splitter"}
				collectorItem = collector.printerCurrent
			case "curr_nozz":
				fallthrough
			case "curr_inp":
				fallthrough
			case "cur_mmu_imp":
				labels = []string{"", splittedName[1] + suffix}
				collectorItem = collector.printerCurrent
			case "bed_curr":
				fallthrough
			case "bedlet_curr":
				labels = []string{"", splittedName[0] + suffix}
				collectorItem = collector.printerCurrent
			case "dwarf_heat_curr":
				labels = []string{"", splittedName[0] + "_" + splittedName[1] + suffix}
				collectorItem = collector.printerCurrent
			case "tmc_sg_x":
				fallthrough
			case "tmc_sg_y":
				fallthrough
			case "tmc_sg_z":
				fallthrough
			case "tmc_sg_e":
				labels = []string{splittedName[2] + suffix}
				collectorItem = collector.printerTmcSg
			case "oc_nozz":
				fallthrough
			case "oc_inp":
				labels = []string{splittedName[1] + suffix}
				collectorItem = collector.printerOvercurrent
			case "curr_nozz_raw":
				fallthrough
			case "curr_inp_raw":
				labels = []string{"", splittedName[1] + suffix}
				collectorItem = collector.printerCurrentRaw
			case "volt_bed_raw":
				fallthrough
			case "volt_nozz_raw":
				labels = []string{"", splittedName[1] + suffix}
				collectorItem = collector.printerVoltageRaw
			case "fan":
				labels = []string{splittedName[1]}
				collectorItem = collector.printerFanActive
			case "fan_speed":
				valueParsed, err = strconv.ParseFloat(v[valueKey], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
					continue // Skip to next iteration if value parsing fails
				}
				printerMetric := prometheus.MustNewConstMetric(collector.printerFanSpeed, prometheus.GaugeValue, valueParsed/255, getLabels(mac, ip, port, []string{"print"})...)
				ch <- printerMetric
				continue
			case "fan_hbr_speed":
				valueParsed, err = strconv.ParseFloat(v[valueKey], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)

					continue // Skip to next iteration if value parsing fails
				}
				ch <- prometheus.MustNewConstMetric(collector.printerFanSpeed, prometheus.GaugeValue, valueParsed/255, getLabels(mac, ip, port, []string{"heatbreak"})...)
				continue
			case "heater_enabled":
				collectorItem = collector.printerHeaterEnabled
			case "loadcell_age":
				collectorItem = collector.printerLoadcellAge
			case "loadcell_value":
				collectorItem = collector.printerLoadcellValue
			case "is_printing":
				collectorItem = collector.printerPrinting
			case "loadcell_hp":
				collectorItem = collector.printerLoadcellHp
			case "bed_pwm":
				labels = []string{"bed" + suffix}
				collectorItem = collector.printerPwm
			case "points_dropped":
				collectorItem = collector.printerPointsDropped
			case "hbr_fan_act":
				labels = []string{"heatbreak"}
				collectorItem = collector.printerFanActive
			case "adj_z":
				collectorItem = collector.printerAxisZAdjustment
			case "filament":
				valueParsed = 0
				if v[valueKey] != "0" {
					valueParsed = 1
				}
				ch <- prometheus.MustNewConstMetric(collector.printerFilament, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{v[valueKey]})...)
				continue
			case "nozzle_pwm":
				labels = []string{"nozzle" + suffix}
				collectorItem = collector.printerPwm
			case "heap":
				valueParsed, err = strconv.ParseFloat(v["free"], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)

					continue // Skip to next iteration if value parsing fails
				}
				ch <- prometheus.MustNewConstMetric(collector.printerHeapFree, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{})...)
				valueParsed, err = strconv.ParseFloat(v["total"], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)

					continue // Skip to next iteration if value parsing fails
				}
				ch <- prometheus.MustNewConstMetric(collector.printerHeapTotal, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{})...)
				continue
			case "print_fan_act":
				labels = []string{"print"}
				collectorItem = collector.printerFanActive
			case "gui_loop_dur":
				collectorItem = collector.printerGuiLoopDuration
			case "fsensor_raw":
				labels = []string{strconv.Itoa(length)}
				collectorItem = collector.printerFSensorRaw
			case "loadcell_xy":
				collectorItem = collector.printerLoadcellXY
			case "cpu_usage":
				collectorItem = collector.printerCPUUsage
			case "loadcell":
				collectorItem = collector.printerLoadcellValueRaw
				valueKey = "r"
			case "loadcell_scale":
				collectorItem = collector.printerLoadcellScale
			case "home_diff":
				valuesList := []string{"ax", "ok", "value"}
				for _, value := range valuesList {
					tool := ""
					if length != -1 {
						tool = strconv.Itoa(length)
					}
					valueParsed, err = strconv.ParseFloat(v[value], 64)
					if err != nil {
						log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
						continue // Skip to next iteration if value parsing fails
					}
					if value == "value" {
						ch <- prometheus.MustNewConstMetric(collector.printerHomeDiffValue, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{tool})...)
					} else if value == "ax" {
						ch <- prometheus.MustNewConstMetric(collector.printerHomeDiffAx, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{tool})...)
					} else if value == "ok" {
						ch <- prometheus.MustNewConstMetric(collector.printerHomeDiffOk, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{tool})...)
					}

				}

				continue
			case "bedlet_pwm":
				labels = []string{"bedlet" + suffix}
				collectorItem = collector.printerPwm
			case "bedlet_reg":
				valueParsed, err = strconv.ParseFloat(v["d"], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
					continue // Skip to next iteration if value parsing fails
				}
				ch <- prometheus.MustNewConstMetric(collector.printerBedletRegulationD, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{"bedlet" + suffix})...)
				valueParsed, err = strconv.ParseFloat(v["i"], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
					continue // Skip to next iteration if value parsing fails
				}
				ch <- prometheus.MustNewConstMetric(collector.printerBedletRegulationI, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{"bedlet" + suffix})...)
				valueParsed, err = strconv.ParseFloat(v["p"], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
					continue // Skip to next iteration if value parsing fails
				}
				ch <- prometheus.MustNewConstMetric(collector.printerBedletRegulationP, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{"bedlet" + suffix})...)
				valueParsed, err = strconv.ParseFloat(v["tc"], 64)
				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
					continue // Skip to next iteration if value parsing fails
				}
				ch <- prometheus.MustNewConstMetric(collector.printerBedletRegulationTc, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, []string{"bedlet" + suffix})...)
				continue
			case "dwarf_parked_raw":
				labels = []string{strconv.Itoa(length)}
				collectorItem = collector.printerDwarfParkedRaw
			case "side_fsensor_raw":
				labels = []string{strconv.Itoa(length)}
				collectorItem = collector.printerSideFSensorRaw
			case "active_extruder":
				collectorItem = collector.printerActiveExtruder
			case "bed_state":
				collectorItem = collector.printerBedState
			case "bedlet_state":
				labels = []string{splittedName[0] + suffix}
				collectorItem = collector.printerBedletState
			case "dwarf_fast_refresh_delay":
				collectorItem = collector.printerDwarfFastRefreshDelay
			case "dwarf_picked_raw":
				labels = []string{strconv.Itoa(length)}
				collectorItem = collector.printerDwarfPickedRaw
			case "buddy_revision":
				collectorItem = collector.printerBuddyRevision
			case "fw_version":
				ch <- prometheus.MustNewConstMetric(collector.printerBuddyFW, prometheus.GaugeValue, 1, getLabels(mac, ip, port, []string{v[valueKey]})...)
				continue
			case "buddy_bom":
				collectorItem = collector.printerBuddyBom
			case "loadcell_threshold_cont":
				collectorItem = collector.printerLoadcellThresholdCont
			case "loadcell_threshold":
				collectorItem = collector.printerLoadcellThreshold
			case "loadcell_hysteresis":
				collectorItem = collector.printerLoadcellHysteresis
			case "ip":
				continue // just ignore
			default:
				log.Error().Msgf("No collector item found for metric %s", k)
				continue // Skip to next iteration if collector item is nil
			}

			if collectorItem == nil {
				log.Debug().Msgf("No collector item found for metric %s", k) // not an error, just debug
				continue                                                     // Skip to next iteration if collector item is nil
			}

			valueParsed, err = strconv.ParseFloat(v[valueKey], 64)
			if err != nil {
				log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
				continue // Skip to next iteration if value parsing fails
			}
			printerMetric := prometheus.MustNewConstMetric(collectorItem, prometheus.GaugeValue, valueParsed, getLabels(mac, ip, port, labels)...)
			ch <- printerMetric
		}
	}
}
