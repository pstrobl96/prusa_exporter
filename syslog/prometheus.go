package syslog

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

type label struct {
	name  string
	value string
}

type collectorBranch struct {
	collector    *prometheus.Desc
	nameOfMetric string
	labels       []label
}

func getLabels(mac string, ip string, labels []label, labelValues ...string) []string {
	for _, l := range labels {
		labelValues = append(labelValues, l.value)
	}
	return append([]string{mac, ip}, labelValues...)
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

var (
	collectorMap = map[string]collectorBranch{
		"active_extruder": {
			collector:    NewCollector().printerActiveExtruder,
			nameOfMetric: "value",
			labels:       []label{},
		},
		"app_start": {
			collector:    NewCollector().printerAppStart,
			nameOfMetric: "value",
			labels:       []label{},
		},
		"axis_z_adjustment": {
			collector:    NewCollector().printerAxisZAdjustment,
			nameOfMetric: "value",
			labels:       []label{},
		},
		"bedlet_regulation": {
			collector:    NewCollector().printerBedletRegulation,
			nameOfMetric: "value",
			labels:       []label{},
		},
		"bedlet_state": {
			collector:    NewCollector().printerBedletState,
			nameOfMetric: "value",
			labels:       []label{},
		},
		"bed_state": {
			collector:    NewCollector().printerBedState,
			nameOfMetric: "value",
			labels:       []label{},
		},
	}
)

// NewCollector is a function that returns new Collector
// NewCollector creates a new instance of the Collector struct with the provided configuration.
// It initializes all the Prometheus metrics used for monitoring different aspects of the printer.
// The defaultLabels parameter is a list of labels that will be included in all the metrics.
// Returns a pointer to the created Collector.
func NewCollector() *Collector {
	defaultLabels := []string{"mac", "ip"}

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

	i := 0
	syslogMetrics.Range(func(key, value interface{}) bool {
		mac := key.(string)
		innermap, ok := value.(map[string]map[string]string)

		log.Trace().Msg("Collecting metrics for " + mac)
		log.Trace().Msg("Innermap: " + innermap["ip"]["value"])

		if !ok {
			log.Error().Msg("Error casting syslog data")
			return false
		}

		ip := innermap["ip"]["value"]

		for k, v := range innermap {
			var (
				collectorItem *prometheus.Desc
				labels        []string
				value         float64
			)

			mapExtract := collectorMap[k]

			valueParsed, e := strconv.ParseFloat(v[mapExtract.nameOfMetric], 64)

			if e != nil {
				log.Error().Msg(e.Error())
				break
			}

			collectorItem = NewCollector().printerActiveExtruder
			labels = getLabels(mac, ip, mapExtract.labels)
			value = valueParsed

			if collectorItem != nil {
				printerMetric := prometheus.MustNewConstMetric(collectorItem, prometheus.GaugeValue,
					value, labels...)
				ch <- printerMetric
			} else {
				log.Debug().Msg("Error creating metric: " + k + " for " + mac + " at " + ip + " with value: " + v[mapExtract.nameOfMetric])
			}

		}

		i++
		return true
	})

	// needs reworking :pug-dance:
	/*for _, s := range configuration.Printers {
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
						printerFSensor/* 				if e != nil {
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

					printerDwarfMcuTemp, ecollector.printerPos
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
	}*/
}
