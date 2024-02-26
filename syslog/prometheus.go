package syslog

import (
	"fmt"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

func getLabels(mac string, ip string, labels []string, labelValues ...string) []string {
	labelValues = append(labelValues, labels...)
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
	return &Collector{
		printerActiveExtruder:        collectorMap["active_extruder"].collector,
		printerAppStart:              collectorMap["app_start"].collector,
		printerAxisZAdjustment:       collectorMap["axis_z_adjustment"].collector,
		printerBedletRegulation:      collectorMap["bedlet_regulation"].collector,
		printerBedletState:           collectorMap["bedlet_state"].collector,
		printerBedState:              collectorMap["bed_state"].collector,
		printerCPUUsage:              collectorMap["cpu_usage"].collector,
		printerCrashCounter:          collectorMap["crash_counter"].collector,
		printerCrashLength:           collectorMap["crash_length"].collector,
		printerCrashRepeatedCounter:  collectorMap["crash_repeated_counter"].collector,
		printerCrashStat:             collectorMap["crash_stat"].collector,
		printerCurrent:               collectorMap["current"].collector,
		printerCurrentRaw:            collectorMap["current_raw"].collector,
		printerDwarfFastRefreshDelay: collectorMap["dwarf_fast_refresh_delay"].collector,
		printerDwarfParkedRaw:        collectorMap["dwarf_parked_raw"].collector,
		printerDwarfPickedRaw:        collectorMap["dwarf_picked_raw"].collector,
		printerEeepromWrite:          collectorMap["eeeprom_write"].collector,
		printerExciteFreq:            collectorMap["excite_freq"].collector,
		printerFanActive:             collectorMap["fan_active"].collector,
		printerFanSpeed:              collectorMap["fan_speed"].collector,
		printerFilename:              collectorMap["filename"].collector,
		printerFSensor:               collectorMap["fsensor"].collector,
		printerFSensorRaw:            collectorMap["fsensor_raw"].collector,
		printerFreqGain:              collectorMap["freq_gain"].collector,
		printerG425Cen:               collectorMap["g425_cen"].collector,
		printerG425Offset:            collectorMap["g425_offset"].collector,
		printerG425Rxy:               collectorMap["g425_rxy"].collector,
		printerG425Rz:                collectorMap["g425_rz"].collector,
		printerG425Xy:                collectorMap["g425_xy"].collector,
		printerG425Z:                 collectorMap["g425_z"].collector,
		printerGcode:                 collectorMap["gcode"].collector,
		printerGuiLoopDuration:       collectorMap["gui_loop_duration"].collector,
		printerHeapFree:              collectorMap["heap_free"].collector,
		printerHeapTotal:             collectorMap["heap_total"].collector,
		printerHeatModelDiscard:      collectorMap["heat_model_discard"].collector,
		printerHeaterEnabled:         collectorMap["heater_enabled"].collector,
		printerHomeDiff:              collectorMap["home_diff"].collector,
		printerIpos:                  collectorMap["ipos"].collector,
		printerLoadcellHysteresis:    collectorMap["loadcell_hysteresis"].collector,
		printerLoadcellScale:         collectorMap["loadcell_scale"].collector,
		printerLoadcellThreshold:     collectorMap["loadcell_threshold"].collector,
		printerLoadcellThresholdCont: collectorMap["loadcell_threshold_cont"].collector,
		printerLoadcellValue:         collectorMap["loadcell_value"].collector,
		printerLoadcellXY:            collectorMap["loadcell_xy"].collector,
		printerMaintaskLoop:          collectorMap["maintask_loop"].collector,
		printerMediaPrefetched:       collectorMap["media_prefetched"].collector,
		printerMMUComm:               collectorMap["mmu_comm"].collector,
		printerModbusReqfail:         collectorMap["modbus_reqfail"].collector,
		printerNetworkIn:             collectorMap["network_in"].collector,
		printerNetworkOut:            collectorMap["network_out"].collector,
		printerOvercurrent:           collectorMap["overcurrent"].collector,
		printerPointsDropped:         collectorMap["points_dropped"].collector,
		printerPos:                   collectorMap["pos"].collector,
		printerPowerPanicCount:       collectorMap["power_panic_count"].collector,
		printerProbeAnalysis:         collectorMap["probe_analysis"].collector,
		printerProbeInfo:             collectorMap["probe_info"].collector,
		printerProbeStart:            collectorMap["probe_start"].collector,
		printerProbeZ:                collectorMap["probe_z"].collector,
		printerProbeZDiff:            collectorMap["probe_z_diff"].collector,
		printerPwm:                   collectorMap["pwm"].collector,
		printerSideFSensor:           collectorMap["side_fsensor"].collector,
		printerSideFSensorRaw:        collectorMap["side_fsensor_raw"].collector,
		printerSyslogInfo:            collectorMap["syslog_info"].collector,
		printerTmcRead:               collectorMap["tmc_read"].collector,
		printerTmcSg:                 collectorMap["tmc_sg"].collector,
		printerTmcWrite:              collectorMap["tmc_write"].collector,
		printerTKAcceleration:        collectorMap["tk_acceleration"].collector,
		printerTemp:                  collectorMap["temp"].collector,
		printerUsbhErrCount:          collectorMap["usbh_err_count"].collector,
		printerVoltage:               collectorMap["voltage"].collector,
		printerVoltageRaw:            collectorMap["voltage_raw"].collector,
		printerXyDev:                 collectorMap["xy_dev"].collector,
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
		ip := nestedmap["ip"]["value"]

		for k, v := range nestedmap {

			mapExtract, ok := collectorMap[k]
			if !ok {
				log.Debug().Msgf("Error extracting metric: %s for %s at %s", k, mac, ip)
				try, ok := overrideMap[k]
				if ok {
					mapExtract.collector = collectorMap[try[0].collectorName].collector
					mapExtract.nameOfMetric = collectorMap[try[0].collectorName].nameOfMetric
					mapExtract.labels = append(collectorMap[try[0].collectorName].labels, try[0].labels...)
				} else {
					continue
				}

			}

			valueParsed, err := strconv.ParseFloat(v[mapExtract.nameOfMetric], 64)
			if err != nil {
				log.Error().Msgf("Error parsing value for metric %s: %s", k, err)
				continue // Skip to next iteration if value parsing fails
			}

			if mapExtract.collector != nil {
				labels := getLabels(mac, ip, mapExtract.labels)
				printerMetric := prometheus.MustNewConstMetric(mapExtract.collector, prometheus.GaugeValue, valueParsed, labels...)
				ch <- printerMetric
			} else {
				log.Debug().Msgf("Error creating metric: %s for %s at %s with value: %s", k, mac, ip, v[mapExtract.nameOfMetric])
			}
		}
	}

}
