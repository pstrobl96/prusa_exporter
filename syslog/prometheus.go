package syslog

import (
	"fmt"
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
		return num, strings.Join(splitted[:indexOfLast], "_"), nil
	}

	return 0, s, nil
}

// Collector is a struct that defines all the syslog metrics
type Collector struct {
	printerActiveExtruder        *prometheus.Desc
	printerAppStart              *prometheus.Desc
	printerAxisZAdjustment       *prometheus.Desc
	printerBedletRegulation      *prometheus.Desc
	printerBedletState           *prometheus.Desc // bedlet_state
	printerBedState              *prometheus.Desc
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
	printerHomeDiff              *prometheus.Desc
	printerIpos                  *prometheus.Desc
	printerLoadcellHysteresis    *prometheus.Desc
	printerLoadcellScale         *prometheus.Desc
	printerLoadcellThreshold     *prometheus.Desc
	printerLoadcellThresholdCont *prometheus.Desc
	printerLoadcellValue         *prometheus.Desc
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
		printerBedletRegulation:      prometheus.NewDesc("prusa_bedlet_regulation", "Bedlet regulation", defaultLabels, nil),
		printerBedletState:           prometheus.NewDesc("prusa_bedlet_state", "Bedlet state", defaultLabels, nil),
		printerBedState:              prometheus.NewDesc("prusa_bed_state", "Bed state", defaultLabels, nil),
		printerCPUUsage:              prometheus.NewDesc("prusa_cpu_usage_ratio", "CPU usage from 0.0 to 1.0", defaultLabels, nil),
		printerCrashCounter:          prometheus.NewDesc("prusa_crash_counter", "Crash counter", defaultLabels, nil),
		printerCrashLength:           prometheus.NewDesc("prusa_crash_length", "Crash length", defaultLabels, nil),
		printerCrashRepeatedCounter:  prometheus.NewDesc("prusa_crash_repeated_counter", "Crash repeated counter", defaultLabels, nil),
		printerCrashStat:             prometheus.NewDesc("prusa_crash_stat", "Crash statistics", defaultLabels, nil),
		printerCurrent:               prometheus.NewDesc("prusa_current", "Current of different devices in / on the printer", append(defaultLabels, "rail", "device"), nil),
		printerCurrentRaw:            prometheus.NewDesc("prusa_current_raw", "Current of different devices in / on the printer in raw sensor value", append(defaultLabels, "rail", "device"), nil),
		printerDwarfFastRefreshDelay: prometheus.NewDesc("prusa_dwarf_fast_refresh_delay", "Dwarf fast refresh delay", defaultLabels, nil),
		printerDwarfParkedRaw:        prometheus.NewDesc("prusa_dwarf_parked_raw", "Dwarf parked raw sensor value", defaultLabels, nil),
		printerDwarfPickedRaw:        prometheus.NewDesc("prusa_dwarf_picked_raw", "Dwarf picked raw sensor value", defaultLabels, nil),
		printerEeepromWrite:          prometheus.NewDesc("prusa_eeeprom_write", "Eeeprom write", defaultLabels, nil),
		printerExciteFreq:            prometheus.NewDesc("prusa_excite_freq", "Excite frequency", defaultLabels, nil),
		printerFanActive:             prometheus.NewDesc("prusa_fan_active", "Fan active", append(defaultLabels, "fan"), nil),
		printerFanSpeed:              prometheus.NewDesc("prusa_syslog_fan_speed", "Fan", append(defaultLabels, "fan"), nil),
		printerFilename:              prometheus.NewDesc("prusa_filename", "Name of printed (b)gcode", append(defaultLabels, "file"), nil),
		printerFSensor:               prometheus.NewDesc("prusa_fsensor", "Filament Sensor", defaultLabels, nil),
		printerFSensorRaw:            prometheus.NewDesc("prusa_fsensor_raw", "Filament Sensor - raw sensor value", defaultLabels, nil),
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
		printerHomeDiff:              prometheus.NewDesc("prusa_home_diff", "Home difference", defaultLabels, nil),
		printerIpos:                  prometheus.NewDesc("prusa_stepper_ipos", "Stepper possition from startup", append(defaultLabels, "axis"), nil),
		printerLoadcellHysteresis:    prometheus.NewDesc("prusa_loadcell_hysteresis", "Loadcell hysteresis", defaultLabels, nil),
		printerLoadcellScale:         prometheus.NewDesc("prusa_loadcell_scale", "Loadcell scale", defaultLabels, nil),
		printerLoadcellThreshold:     prometheus.NewDesc("prusa_loadcell_threshold", "Loadcell threshold", defaultLabels, nil),
		printerLoadcellThresholdCont: prometheus.NewDesc("prusa_loadcell_threshold_cont", "Loadcell threshold continuous", defaultLabels, nil),
		printerLoadcellValue:         prometheus.NewDesc("prusa_loadcell", "Value from loadcell sensor", defaultLabels, nil),
		printerMaintaskLoop:          prometheus.NewDesc("prusa_maintask_loop", "Maintask loop", defaultLabels, nil),
		printerMediaPrefetched:       prometheus.NewDesc("prusa_eeeprom_write", "Eeeprom write", defaultLabels, nil),
		printerMMUComm:               prometheus.NewDesc("prusa_mmu_comm", "MMU communication", defaultLabels, nil),
		printerModbusReqfail:         prometheus.NewDesc("prusa_modbus_reqfail", "Modbus request fail", defaultLabels, nil),
		printerNetworkIn:             prometheus.NewDesc("prusa_network_in", "Network in", append(defaultLabels, "device"), nil),
		printerNetworkOut:            prometheus.NewDesc("prusa_network_out", "Network out", append(defaultLabels, "device"), nil),
		printerOvercurrent:           prometheus.NewDesc("prusa_overcurrent", "Overcurrent of different devices in / on the printer", append(defaultLabels, "device"), nil),
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
		printerSideFSensorRaw:        prometheus.NewDesc("prusa_side_fsensor_raw", "Side Filament Sensor - raw sensor value", defaultLabels, nil),
		printerSyslogInfo:            prometheus.NewDesc("prusa_syslog_info", "Buddy syslog info", append(defaultLabels, "revision", "bom"), nil),
		printerTmcRead:               prometheus.NewDesc("prusa_tmc_read", "Trinamic read", append(defaultLabels, "axis"), nil),
		printerTmcSg:                 prometheus.NewDesc("prusa_tmc_sg", "Trinamic SG", append(defaultLabels, "axis"), nil),
		printerTmcWrite:              prometheus.NewDesc("prusa_tmc_write", "Trinamic write", append(defaultLabels, "axis"), nil),
		printerTKAcceleration:        prometheus.NewDesc("prusa_tk_acceleration", "TK acceleration", defaultLabels, nil),
		printerTemp:                  prometheus.NewDesc("prusa_temp", "Temperature of different devices in / on the printer", append(defaultLabels, "device"), nil),
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
	ch <- collector.printerBedletRegulation
	ch <- collector.printerBedletState
	ch <- collector.printerBedState
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
	ch <- collector.printerHomeDiff
	ch <- collector.printerIpos
	ch <- collector.printerLoadcellHysteresis
	ch <- collector.printerLoadcellScale
	ch <- collector.printerLoadcellThreshold
	ch <- collector.printerLoadcellThresholdCont
	ch <- collector.printerLoadcellValue
	ch <- collector.printerMaintaskLoop
	ch <- collector.printerMediaPrefetched
	ch <- collector.printerMMUComm
	ch <- collector.printerModbusReqfail
	ch <- collector.printerNetworkIn
	ch <- collector.printerNetworkOut
	ch <- collector.printerOvercurrent
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
	ch <- collector.printerSyslogInfo
	ch <- collector.printerTmcRead
	ch <- collector.printerTmcSg
	ch <- collector.printerTmcWrite
	ch <- collector.printerTKAcceleration
	ch <- collector.printerTemp
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
		fmt.Println(mac)
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
				labels        []string
			)

			length, name, err := getNumberOf(k)

			if err != nil {
				log.Error().Msgf("Error parsing metric name %s: %s", k, err)
				continue // Skip to next iteration if metric name parsing fails
			}

			if length > 0 {
				k = name
			}

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
				fallthrough
			case "dwarf_board_temp":
				fallthrough
			case "dwarf_mcu_temp":
				fallthrough
			case "dwarfs_mcu_temp":
				fallthrough
			case "dwarfs_board_temp":
				fallthrough
			case "bed_mcu_temp":
				fallthrough
			case "bedlet_temp":
				valueParsed, err := strconv.ParseFloat(v["value"], 64)
				collectorItem = collector.printerTemp
				splitedName := strings.Split(k, "_")
				if splitedName[0] == "temp" {
					labels = getLabels(mac, ip, port, []string{splitedName[1] + "_" + strconv.Itoa(length)})
				} else if len(splitedName) == 2 {
					labels = getLabels(mac, ip, port, []string{splitedName[0] + "_" + strconv.Itoa(length)})
				} else {
					labels = getLabels(mac, ip, port, []string{splitedName[0] + "_" + splitedName[1] + "_" + strconv.Itoa(length)})
				}

				if err != nil {
					log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
					continue // Skip to next iteration if value parsing fails
				}

				printerMetric := prometheus.MustNewConstMetric(collectorItem, prometheus.GaugeValue, valueParsed, labels...)
				ch <- printerMetric
				break
			}
		}
	}
}
