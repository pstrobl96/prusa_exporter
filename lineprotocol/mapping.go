package syslog

import (
	"log"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

func getLabels(mac string, ip string, labels []string, labelValues ...string) []string {
	labelValues = append(labelValues, labels...)
	return append([]string{mac, ip}, labelValues...)
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

	return -1, s, nil
}

// Collector is a struct that defines all the syslog metrics
type Collector struct {
	prusaActiveExtruder        *prometheus.Desc
	prusaAppStart              *prometheus.Desc
	prusaAxisZAdjustment       *prometheus.Desc
	prusaBedletRegulationD     *prometheus.Desc // bedlet_regulation_d
	prusaBedletRegulationI     *prometheus.Desc // bedlet_regulation_i
	prusaBedletRegulationP     *prometheus.Desc // bedlet_regulation_p
	prusaBedletRegulationTc    *prometheus.Desc // bedlet_regulation_tc
	prusaBedletState           *prometheus.Desc // bedlet_state
	prusaBedState              *prometheus.Desc
	prusaBuddyInfo             *prometheus.Desc
	prusaCPUUsage              *prometheus.Desc
	prusaCrashSpeed            *prometheus.Desc
	prusaCrashLength           *prometheus.Desc
	prusaCrashStat             *prometheus.Desc
	prusaCurrent               *prometheus.Desc
	prusaCurrentRaw            *prometheus.Desc
	prusaDwarfFastRefreshDelay *prometheus.Desc
	prusaDwarfParkedRaw        *prometheus.Desc
	prusaDwarfPickedRaw        *prometheus.Desc
	prusaEeepromWrite          *prometheus.Desc
	prusaExciteFreq            *prometheus.Desc
	prusaFanActive             *prometheus.Desc
	prusaFanSpeed              *prometheus.Desc
	prusaFilename              *prometheus.Desc
	prusaFilament              *prometheus.Desc
	prusaFSensor               *prometheus.Desc
	prusaFSensorRaw            *prometheus.Desc
	prusaFreqGain              *prometheus.Desc
	prusaG425Cen               *prometheus.Desc
	prusaG425Offset            *prometheus.Desc
	prusaG425Rxy               *prometheus.Desc
	prusaG425Rz                *prometheus.Desc
	prusaG425Xy                *prometheus.Desc
	prusaG425XyDev             *prometheus.Desc
	prusaG425Z                 *prometheus.Desc
	prusaGcode                 *prometheus.Desc
	prusaGuiLoopDuration       *prometheus.Desc
	prusaHeapFree              *prometheus.Desc
	prusaHeapTotal             *prometheus.Desc
	prusaHeaterEnabled         *prometheus.Desc
	prusaHomeDiffOk            *prometheus.Desc
	prusaHomeDiff              *prometheus.Desc
	prusaIpos                  *prometheus.Desc
	prusaLoadcellAge           *prometheus.Desc
	prusaLoadcellHp            *prometheus.Desc
	prusaLoadcellHysteresis    *prometheus.Desc
	prusaLoadcellScale         *prometheus.Desc
	prusaLoadcellThreshold     *prometheus.Desc
	prusaLoadcellThresholdCont *prometheus.Desc
	prusaLoadcellValue         *prometheus.Desc
	prusaLoadcellValueRaw      *prometheus.Desc
	prusaLoadcellXY            *prometheus.Desc
	prusaMaintaskLoop          *prometheus.Desc
	prusaMediaPrefetched       *prometheus.Desc
	prusaMMUComm               *prometheus.Desc
	prusaModbusReqfail         *prometheus.Desc
	prusaNetworkIn             *prometheus.Desc
	prusaNetworkOut            *prometheus.Desc
	prusaOvercurrent           *prometheus.Desc
	prusaPointsDropped         *prometheus.Desc
	prusaPos                   *prometheus.Desc
	prusaPowerPanicCount       *prometheus.Desc
	prusaPrinting              *prometheus.Desc
	prusaProbeAnalysis         *prometheus.Desc
	prusaProbeWindowStart      *prometheus.Desc
	prusaProbeWindowFallEnd    *prometheus.Desc
	prusaProbeWindowRiseStart  *prometheus.Desc
	prusaProbeWindowEnd        *prometheus.Desc
	prusaProbeStart            *prometheus.Desc
	prusaProbeZ                *prometheus.Desc // probe_z
	prusaProbeZDiff            *prometheus.Desc
	prusaPwm                   *prometheus.Desc
	prusaSideFSensor           *prometheus.Desc // side_fsensor
	prusaSideFSensorRaw        *prometheus.Desc
	prusaSyslogInfo            *prometheus.Desc // revision, bom
	prusaTmcRead               *prometheus.Desc
	prusaTmcSg                 *prometheus.Desc
	prusaTmcWrite              *prometheus.Desc
	prusaTKAcceleration        *prometheus.Desc
	prusaTemp                  *prometheus.Desc
	prusaTempTarget            *prometheus.Desc
	prusaUsbhErrCount          *prometheus.Desc
	prusaVoltage               *prometheus.Desc
	prusaVoltageRaw            *prometheus.Desc
	prusaXyDev                 *prometheus.Desc
	prusaBuddyTimeUs           *prometheus.Desc
	prusaPuppyTimeUs           *prometheus.Desc
	prusaSyncRoundtripUs       *prometheus.Desc
	prusaPuppyOffsetUs         *prometheus.Desc
	prusaPuppyDriftPpb         *prometheus.Desc
	prusaPuppyAverageOffsetUs  *prometheus.Desc
	prusaPuppyAverageDriftPpb  *prometheus.Desc
	prusaSyslogUp              *prometheus.Desc
	prusaPrintFilename         *prometheus.Desc
}

// NewCollector is a function that returns new Collector
// NewCollector creates a new instance of the Collector struct with the provided configuration.
// It initializes all the Prometheus metrics used for monitoring different aspects of the prusa.
// The defaultLabels parameter is a list of labels that will be included in all the metrics.
// Returns a pointer to the created Collector.
func NewCollector(syslogTTL int) *Collector {
	defaultLabels := []string{"mac", "ip"}
	if syslogTTL < 1 {
		log.Panic("syslog TTL must be greater than 0")
	}
	ttl = syslogTTL
	return &Collector{
		//// Important
		prusaActiveExtruder: prometheus.NewDesc("prusa_active_extruder", "Active extruder - used for XL", defaultLabels, nil),
		prusaBuddyInfo:      prometheus.NewDesc("prusa_buddy_info", "Buddy firmware version", append(defaultLabels, "fw_version", "buddy_bom", "buddy_revision"), nil), // had to cache this one
		prusaCPUUsage:       prometheus.NewDesc("prusa_cpu_usage_ratio", "CPU usage from 0.0 to 1.0", defaultLabels, nil),
		prusaCurrent:        prometheus.NewDesc("prusa_current_amperes", "Current of different devices in / on the prusa in mampers", append(defaultLabels, "rail", "device"), nil),
		prusaLoadcellValue:  prometheus.NewDesc("prusa_loadcell", "Value from loadcell sensor", defaultLabels, nil),
		prusaSyslogUp:       prometheus.NewDesc("prusa_up_lineprotocol", "Printer up - from syslog metric - ttl is by default 60 seconds but can be different and it depends on choosen interval. That means if prusa wont sent any data for 60 seconds is considered down.", defaultLabels, nil),
		prusaVoltage:        prometheus.NewDesc("prusa_voltage_volts", "Voltage of different devices in / on the prusa", append(defaultLabels, "rail", "device"), nil),
		prusaNetworkIn:      prometheus.NewDesc("prusa_network_in_bytes", "Network in", append(defaultLabels, "device"), nil),
		prusaNetworkOut:     prometheus.NewDesc("prusa_network_out_bytes", "Network out", append(defaultLabels, "device"), nil),
		prusaOvercurrent:    prometheus.NewDesc("prusa_overcurrent", "Overcurrent of different devices in / on the prusa", append(defaultLabels, "device"), nil),
		prusaFanActive:      prometheus.NewDesc("prusa_fan_active", "Fan active", append(defaultLabels, "fan"), nil),
		prusaFanSpeed:       prometheus.NewDesc("prusa_fan_speed_ratio", "Fan", append(defaultLabels, "fan"), nil),
		prusaFilament:       prometheus.NewDesc("prusa_filament", "Currently loaded filament", append(defaultLabels, "filament"), nil),

		//// Not so important
		prusaAppStart:              prometheus.NewDesc("prusa_app_start", "Application start", defaultLabels, nil),
		prusaAxisZAdjustment:       prometheus.NewDesc("prusa_axis_z_adjustment", "Axis Z adjustment", defaultLabels, nil),
		prusaBedletRegulationD:     prometheus.NewDesc("prusa_bedlet_regulation_d", "Bedlet regulation d value", append(defaultLabels, "bedlet"), nil),
		prusaBedletRegulationI:     prometheus.NewDesc("prusa_bedlet_regulation_i", "Bedlet regulation i value", append(defaultLabels, "bedlet"), nil),
		prusaBedletRegulationP:     prometheus.NewDesc("prusa_bedlet_regulation_p", "Bedlet regulation p value", append(defaultLabels, "bedlet"), nil),
		prusaBedletRegulationTc:    prometheus.NewDesc("prusa_bedlet_regulation_tc", "Bedlet regulation tc value", append(defaultLabels, "bedlet"), nil),
		prusaBedletState:           prometheus.NewDesc("prusa_bedlet_state", "Bedlet state", append(defaultLabels, "bedlet"), nil),
		prusaBedState:              prometheus.NewDesc("prusa_bed_state", "Bed state", defaultLabels, nil),
		prusaCrashSpeed:            prometheus.NewDesc("prusa_crash_speed", "Crash Speed", append(defaultLabels, "axis", "sens", "period"), nil),
		prusaCrashLength:           prometheus.NewDesc("prusa_crash_length", "Crash length", append(defaultLabels, "x", "y"), nil),
		prusaCrashStat:             prometheus.NewDesc("prusa_crash_stat", "Crash statistics", append(defaultLabels, "axis"), nil),
		prusaCurrentRaw:            prometheus.NewDesc("prusa_current_raw", "Current of different devices in / on the prusa in raw sensor value", append(defaultLabels, "rail", "device"), nil),
		prusaDwarfFastRefreshDelay: prometheus.NewDesc("prusa_dwarf_fast_refresh_delay", "Dwarf fast refresh delay", defaultLabels, nil),
		prusaDwarfParkedRaw:        prometheus.NewDesc("prusa_dwarf_parked_raw", "Dwarf parked raw sensor value", append(defaultLabels, "tool"), nil),
		prusaDwarfPickedRaw:        prometheus.NewDesc("prusa_dwarf_picked_raw", "Dwarf picked raw sensor value", append(defaultLabels, "tool"), nil),
		prusaEeepromWrite:          prometheus.NewDesc("prusa_eeeprom_write", "Eeeprom write", defaultLabels, nil),
		prusaExciteFreq:            prometheus.NewDesc("prusa_excite_freq", "Excite frequency", defaultLabels, nil),
		prusaFilename:              prometheus.NewDesc("prusa_filename", "Name of printed (b)gcode", append(defaultLabels, "file"), nil),
		prusaFSensor:               prometheus.NewDesc("prusa_fsensor", "Filament Sensor", defaultLabels, nil),
		prusaFSensorRaw:            prometheus.NewDesc("prusa_fsensor_raw", "Filament Sensor - raw sensor value", append(defaultLabels, "sensor"), nil),
		prusaFreqGain:              prometheus.NewDesc("prusa_freq_gain", "Frequency gain", defaultLabels, nil),
		prusaG425Cen:               prometheus.NewDesc("prusa_g425_cen", "Absolute tool center - an input for offset computation [mm]", append(defaultLabels, "t", "x", "y", "z"), nil),   // ",t=%u x=%.3f,y=%.3f,z=%.3f"
		prusaG425Offset:            prometheus.NewDesc("prusa_g425_off", "Offset from the absolute tool center [mm]", append(defaultLabels, "t", "x", "y", "z"), nil),                     //  ",t=%u x=%.3f,y=%.3f,z=%.3f",
		prusaG425Rxy:               prometheus.NewDesc("prusa_g425_rxy", "Raw XY probe [mm]", append(defaultLabels, "t", "p", "a", "x", "y"), nil),                                        // ",t=%u,p=%u,a=%.3f x=%.3f,y=%.3f"
		prusaG425Rz:                prometheus.NewDesc("prusa_g425_rz", "Raw Z probe [mm]", append(defaultLabels, "t", "p", "x", "y", "z"), nil),                                          // ",t=%u,p=%u x=%.3f,y=%.3f,z=%.3f",
		prusaG425Xy:                prometheus.NewDesc("prusa_g425_xy", "Verified XY probe - two raw probes agree on position [mm]", append(defaultLabels, "t", "p", "a", "x", "y"), nil), // ",t=%u,p=%u,a=%.3f x=%.3f,y=%.3f"
		prusaG425XyDev:             prometheus.NewDesc("prusa_g425_xy_dev", "Max deviation", defaultLabels, nil),                                                                          // ",t=%u,p=%u,x=%.3f,y=%.3f z=%.3f",
		prusaG425Z:                 prometheus.NewDesc("prusa_g425_z", "Averaged Z probe - N raw probes averaged [mm]", append(defaultLabels, "t", "p", "x", "y", "z"), nil),              // ",t=%u,p=%u,x=%.3f,y=%.3f z=%.3f",
		prusaGcode:                 prometheus.NewDesc("prusa_gcode", "Printed GCode", append(defaultLabels, "gcode"), nil),
		prusaGuiLoopDuration:       prometheus.NewDesc("prusa_gui_loop_duration", "Gui loop duration", defaultLabels, nil),
		prusaHeapFree:              prometheus.NewDesc("prusa_heap_free", "Free heap", defaultLabels, nil),
		prusaHeapTotal:             prometheus.NewDesc("prusa_heap_total", "Total heap", defaultLabels, nil),
		prusaHeaterEnabled:         prometheus.NewDesc("prusa_heater_enabled", "Heater enabled", defaultLabels, nil),
		prusaHomeDiffOk:            prometheus.NewDesc("prusa_home_diff_ok", "Home diff ok", append(defaultLabels, "axis", "attempts"), nil),
		prusaHomeDiff:              prometheus.NewDesc("prusa_home_diff", "Home diff value", append(defaultLabels, "axis", "attempts"), nil),
		prusaIpos:                  prometheus.NewDesc("prusa_stepper_ipos", "Stepper possition from startup", append(defaultLabels, "axis"), nil),
		prusaLoadcellAge:           prometheus.NewDesc("prusa_loadcell_age", "Loadcell age", defaultLabels, nil),
		prusaLoadcellHysteresis:    prometheus.NewDesc("prusa_loadcell_hysteresis", "Loadcell hysteresis", defaultLabels, nil),
		prusaLoadcellHp:            prometheus.NewDesc("prusa_loadcell_hp", "Loadcell filtered z load", defaultLabels, nil),
		prusaLoadcellScale:         prometheus.NewDesc("prusa_loadcell_scale", "Loadcell scale", defaultLabels, nil),
		prusaLoadcellThreshold:     prometheus.NewDesc("prusa_loadcell_threshold", "Loadcell threshold", defaultLabels, nil),
		prusaLoadcellThresholdCont: prometheus.NewDesc("prusa_loadcell_threshold_cont", "Loadcell threshold continuous", defaultLabels, nil),
		prusaLoadcellValueRaw:      prometheus.NewDesc("prusa_loadcell_raw", "Value from loadcell sensor in raw sensor value", defaultLabels, nil),
		prusaLoadcellXY:            prometheus.NewDesc("prusa_loadcell_xy", "Loadcell XY", defaultLabels, nil),
		prusaMaintaskLoop:          prometheus.NewDesc("prusa_maintask_loop", "Maintask loop", defaultLabels, nil),
		prusaMediaPrefetched:       prometheus.NewDesc("prusa_media_prefetched_bytes", "Media prefetched in bytes", defaultLabels, nil),
		prusaMMUComm:               prometheus.NewDesc("prusa_mmu_comm", "MMU communication", append(defaultLabels, "msg"), nil),
		prusaModbusReqfail:         prometheus.NewDesc("prusa_modbus_reqfail", "Modbus request fail", defaultLabels, nil),
		prusaPrinting:              prometheus.NewDesc("prusa_printing", "Printing prusa", defaultLabels, nil),
		prusaPointsDropped:         prometheus.NewDesc("prusa_points_dropped", "Points dropped", defaultLabels, nil),
		prusaPos:                   prometheus.NewDesc("prusa_stepper_pos", "Stepper possition", append(defaultLabels, "axis"), nil),
		prusaPowerPanicCount:       prometheus.NewDesc("prusa_power_panic_count", "Power panic triggered", defaultLabels, nil),
		prusaProbeAnalysis:         prometheus.NewDesc("prusa_probe_analysis", "Probe analysis", append(defaultLabels, "desc"), nil),
		prusaProbeWindowStart:      prometheus.NewDesc("prusa_probe_window_start", "Probe window analysis start", defaultLabels, nil),
		prusaProbeWindowFallEnd:    prometheus.NewDesc("prusa_probe_window_fall_end", "Probe window fall ended", defaultLabels, nil),
		prusaProbeWindowRiseStart:  prometheus.NewDesc("prusa_probe_window_rise_start", "Probe window rise start", defaultLabels, nil),
		prusaProbeWindowEnd:        prometheus.NewDesc("prusa_probe_window_analysis_end", "Probe window analysis", defaultLabels, nil),
		prusaProbeStart:            prometheus.NewDesc("prusa_probe_start", "Probe start", defaultLabels, nil),
		prusaProbeZ:                prometheus.NewDesc("prusa_probe_z", "Probe Z", append(defaultLabels, "x", "y"), nil),
		prusaProbeZDiff:            prometheus.NewDesc("prusa_probe_z_diff", "Probe Z difference", defaultLabels, nil),
		prusaPwm:                   prometheus.NewDesc("prusa_pwm", "PWM value of nozzle and bed mostly", append(defaultLabels, "device"), nil),
		prusaSideFSensor:           prometheus.NewDesc("prusa_side_fsensor", "Side Filament Sensor", defaultLabels, nil),
		prusaSideFSensorRaw:        prometheus.NewDesc("prusa_side_fsensor_raw", "Side Filament Sensor - raw sensor value", append(defaultLabels, "sensor"), nil),
		prusaTmcRead:               prometheus.NewDesc("prusa_tmc_read", "Trinamic read", append(defaultLabels, "axis", "reg_addr", "reg_addr_name"), nil), //     metric_record_custom(&metric_read, ",ax=%c reg=%ui,regn=\"%s\",value=%ui",
		prusaTmcSg:                 prometheus.NewDesc("prusa_tmc_sg", "Trinamic SG", append(defaultLabels, "axis"), nil),
		prusaTmcWrite:              prometheus.NewDesc("prusa_tmc_write", "Trinamic write", append(defaultLabels, "axis", "reg_addr", "reg_addr_name"), nil),
		prusaTKAcceleration:        prometheus.NewDesc("prusa_tk_acceleration", "TK acceleration", defaultLabels, nil),
		prusaTemp:                  prometheus.NewDesc("prusa_temp", "Temperature of different devices in / on the prusa", append(defaultLabels, "device"), nil),
		prusaTempTarget:            prometheus.NewDesc("prusa_temp_target", "Target temperature of different devices in / on the prusa", append(defaultLabels, "device"), nil),
		prusaUsbhErrCount:          prometheus.NewDesc("prusa_usbh_err_count", "USBH error counter", defaultLabels, nil),
		prusaVoltageRaw:            prometheus.NewDesc("prusa_voltage_raw", "Voltage of different devices in / on the prusa in raw sensor value", append(defaultLabels, "rail", "device"), nil),
		prusaXyDev:                 prometheus.NewDesc("prusa_xy_dev", "XY deviation - max difference between two raw probes [mm]", defaultLabels, nil),
		prusaBuddyTimeUs:           prometheus.NewDesc("prusa_buddy_time_seconds", "Buddy time in microseconds", defaultLabels, nil),
		prusaPuppyTimeUs:           prometheus.NewDesc("prusa_puppy_time_seconds", "Puppy time in microseconds", defaultLabels, nil),
		prusaSyncRoundtripUs:       prometheus.NewDesc("prusa_sync_roundtrip_seconds", "Sync roundtrip in microseconds", defaultLabels, nil),
		prusaPuppyOffsetUs:         prometheus.NewDesc("prusa_puppy_offset_seconds", "Puppy offset in microseconds", defaultLabels, nil),
		prusaPuppyDriftPpb:         prometheus.NewDesc("prusa_puppy_drift_ppb", "Puppy drift in ppb", defaultLabels, nil),
		prusaPuppyAverageOffsetUs:  prometheus.NewDesc("prusa_puppy_average_offset_seconds", "Puppy average offset in microseconds", defaultLabels, nil),
		prusaPuppyAverageDriftPpb:  prometheus.NewDesc("prusa_puppy_average_drift_ppb", "Puppy average drift in ppb", defaultLabels, nil),
		prusaPrintFilename:         prometheus.NewDesc("prusa_printed_filename", "Printed file name", append(defaultLabels, "filename"), nil),
	}
}

// Describe is a function that describes all the metrics
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.prusaActiveExtruder
	ch <- collector.prusaAppStart
	ch <- collector.prusaAxisZAdjustment
	ch <- collector.prusaBedletRegulationD
	ch <- collector.prusaBedletRegulationI
	ch <- collector.prusaBedletRegulationP
	ch <- collector.prusaBedletRegulationTc
	ch <- collector.prusaBedletState
	ch <- collector.prusaBedState
	ch <- collector.prusaBuddyBom
	ch <- collector.prusaBuddyRevision
	ch <- collector.prusaBuddyFW
	ch <- collector.prusaCPUUsage
	ch <- collector.prusaCrashSpeed
	ch <- collector.prusaCrashLength
	ch <- collector.prusaCrashStat
	ch <- collector.prusaCurrent
	ch <- collector.prusaCurrentRaw
	ch <- collector.prusaDwarfFastRefreshDelay
	ch <- collector.prusaDwarfParkedRaw
	ch <- collector.prusaDwarfPickedRaw
	ch <- collector.prusaEeepromWrite
	ch <- collector.prusaExciteFreq
	ch <- collector.prusaFanActive
	ch <- collector.prusaFanSpeed
	ch <- collector.prusaFilament
	ch <- collector.prusaFilename
	ch <- collector.prusaFSensor
	ch <- collector.prusaFSensorRaw
	ch <- collector.prusaFreqGain
	ch <- collector.prusaG425Cen
	ch <- collector.prusaG425Offset
	ch <- collector.prusaG425Rxy
	ch <- collector.prusaG425Rz
	ch <- collector.prusaG425Xy
	ch <- collector.prusaG425Z
	ch <- collector.prusaGcode
	ch <- collector.prusaGuiLoopDuration
	ch <- collector.prusaHeapFree
	ch <- collector.prusaHeapTotal
	ch <- collector.prusaHeaterEnabled
	ch <- collector.prusaHomeDiffOk
	ch <- collector.prusaHomeDiff
	ch <- collector.prusaIpos
	ch <- collector.prusaLoadcellAge
	ch <- collector.prusaLoadcellHp
	ch <- collector.prusaLoadcellHysteresis
	ch <- collector.prusaLoadcellScale
	ch <- collector.prusaLoadcellThreshold
	ch <- collector.prusaLoadcellThresholdCont
	ch <- collector.prusaLoadcellValue
	ch <- collector.prusaLoadcellValueRaw
	ch <- collector.prusaMaintaskLoop
	ch <- collector.prusaMediaPrefetched
	ch <- collector.prusaMMUComm
	ch <- collector.prusaModbusReqfail
	ch <- collector.prusaNetworkIn
	ch <- collector.prusaNetworkOut
	ch <- collector.prusaOvercurrent
	ch <- collector.prusaPointsDropped
	ch <- collector.prusaPos
	ch <- collector.prusaPowerPanicCount
	ch <- collector.prusaProbeAnalysis
	ch <- collector.prusaProbeWindowStart
	ch <- collector.prusaProbeWindowFallEnd
	ch <- collector.prusaProbeWindowRiseStart
	ch <- collector.prusaProbeWindowEnd
	ch <- collector.prusaProbeStart
	ch <- collector.prusaProbeZ
	ch <- collector.prusaProbeZDiff
	ch <- collector.prusaPwm
	ch <- collector.prusaSideFSensor
	ch <- collector.prusaSideFSensorRaw
	ch <- collector.prusaTmcRead
	ch <- collector.prusaTmcSg
	ch <- collector.prusaTmcWrite
	ch <- collector.prusaTKAcceleration
	ch <- collector.prusaTemp
	ch <- collector.prusaTempTarget
	ch <- collector.prusaUsbhErrCount
	ch <- collector.prusaVoltage
	ch <- collector.prusaVoltageRaw
	ch <- collector.prusaXyDev
	ch <- collector.prusaBuddyTimeUs
	ch <- collector.prusaPuppyTimeUs
	ch <- collector.prusaSyncRoundtripUs
	ch <- collector.prusaPuppyOffsetUs
	ch <- collector.prusaPuppyDriftPpb
	ch <- collector.prusaPuppyAverageOffsetUs
	ch <- collector.prusaPuppyAverageDriftPpb
	ch <- collector.prusaSyslogUp
	ch <- collector.prusaPrintFilename
}
