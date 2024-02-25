package syslog

import (
	"regexp"
	"sync"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

type labels struct {
	name  string
	value string
}

type patterns struct {
	pattern string
	fields  []string
}

var (
	// syslogMetrics is a map of mac addresses and their metrics
	syslogMetrics = sync.Map{}

	// regexpPatterns is a map that stores the regular expression patterns for different types of log messages.
	// Each pattern is associated with a set of named capture groups and corresponding field names.
	regexpPatterns = map[string]patterns{
		"v_integer":              {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) v=(?P<value>-?\d+)i (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}},
		"float":                  {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) v=(?P<value>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}},
		"integer":                {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) v=(?P<value>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}},
		"string":                 {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) v="(?P<value>[-\d\.]+)" (?P<timestamp>\d+)`, fields: []string{"name", "value", "timestamp"}},
		"xyv":                    {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) x=(?P<x>[-\d\.]+),y=(?P<y>[-\d\.]+),v=(?P<value>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "x", "y", "value", "timestamp"}},
		"free_total":             {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) free=(?P<free>[-\d\.]+)i,total=(?P<total>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "free", "total", "timestamp"}},
		"axis_sens_period_speed": {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),axis=(?P<axis>[-\d\.]+) sens=(?P<sens>[-\d\.]+)i,period=(?P<period>[-\d\.]+)i,speed=(?P<speed>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "axis", "sens", "period", "speed", "timestamp"}},
		"axis_last_total":        {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),axis=(?P<axis>[-\d\.]+) last=(?P<last>[-\d\.]+)i,total=(?P<total>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "axis", "last", "total", "timestamp"}},
		"xyz":                    {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) x=(?P<x>[-\d\.]+),y=(?P<y>[-\d\.]+),z=(?P<z>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "x", "y", "z", "timestamp"}},
		"a_f_x_y_z":              {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) a=(?P<a>[-\d\.]+),f=(?P<f>[-\d\.]+),x=(?P<x>[-\d\.]+),y=(?P<y>[-\d\.]+),z=(?P<z>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "a", "f", "x", "y", "z", "timestamp"}},
		"ax_ok_v_n":              {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),ax=(?P<ax>[-\d\.]+),ok=(?P<ok>[-\d\.]+) v=(?P<v>[-\d\.]+),n=(?P<n>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "ax", "ok", "value", "n", "timestamp"}},
		"ok_desc":                {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) ok=(?P<ok>[-\d\.]+),desc="(?P<desc>[-\d\.]+)" (?P<timestamp>\d+)`, fields: []string{"name", "ok", "desc", "timestamp"}},
		"sent":                   {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) sent=(?P<sent>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "sent", "timestamp"}},
		"recv":                   {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) recv=(?P<recv>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "recv", "timestamp"}},
		"n_t_m":                  {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) t=(?P<t>[-\d\.]+),m=(?P<m>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "n", "t", "m", "timestamp"}},
		"n_u":                    {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) u=(?P<u>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "n", "u", "timestamp"}},
		"n_a_value":              {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+),a=(?P<a>[-\d\.]+) value=(?P<value>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "n", "a", "value", "timestamp"}},
		"n_a_value_integer":      {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+),a=(?P<a>[-\d\.]+) value=(?P<value>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "n", "a", "value", "timestamp"}},
		"n_st_f_r_ri_sp":         {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) st=(?P<st>[-\d\.]+),f=(?P<f>[-\d\.]+),r=(?P<r>[-\d\.]+),ri=(?P<ri>[-\d\.]+),sp=(?P<sp>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "n", "st", "f", "r", "ri", "sp", "timestamp"}},
		"n_v_integer":            {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) v=(?P<v>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "n", "value", "timestamp"}},
		"xy":                     {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) x=(?P<x>[-\d\.]+),y=(?P<y>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "x", "y", "timestamp"}},
		"as_fe_rs_ae":            {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) as=(?P<as>[-\d\.]+),fe=(?P<fe>[-\d\.]+),rs=(?P<rs>[-\d\.]+),ae=(?P<ae>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "as", "fe", "rs", "ae", "timestamp"}},
		"ax_reg_regn_value":      {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),ax=(?P<ax>[-\d\.]+) reg=(?P<reg>[-\d\.]+),regn="(?P<regn>[-\d\.]+)",value=(?P<value>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "ax", "reg", "regn", "value", "timestamp"}},
		"fan_state_pwm_measured": {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),fan=(?P<fan>[-\d\.]+) state=(?P<state>[-\d\.]+),pwm=(?P<pwm>[-\d\.]+),measured=(?P<measured>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "fan", "state", "pwm", "measured", "timestamp"}},
		"t_p_a_x_y":              {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),t=(?P<t>[-\d\.]+),p=(?P<p>[-\d\.]+),a=(?P<a>[-\d\.]+) x=(?P<x>[-\d\.]+),y=(?P<y>[-\d\.]+)`, fields: []string{"name", "t", "p", "a", "x", "y"}},
		"t_p_x_y_z":              {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),t=(?P<t>[-\d\.]+),p=(?P<p>[-\d\.]+) x=(?P<x>[-\d\.]+),y=(?P<y>[-\d\.]+),z=(?P<z>[-\d\.]+)`, fields: []string{"name", "t", "p", "x", "y", "z"}},
		"t_x_y_z":                {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),t=(?P<t>[-\d\.]+) x=(?P<x>[-\d\.]+),y=(?P<y>[-\d\.]+),z=(?P<z>[-\d\.]+)`, fields: []string{"name", "t", "x", "y", "z"}},
		"n_v":                    {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) v=(?P<v>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "n", "value", "timestamp"}},
		"n_v_e_integer":          {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) v=(?P<v>[-\d\.]+)i,e=(?P<e>[-\d\.]+)i (?P<timestamp>\d+)`, fields: []string{"name", "n", "value", "e", "timestamp"}},
		"n_p_i_d_tc":             {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) p=(?P<p>[-\d\.]+),i=(?P<i>[-\d\.]+),d=(?P<d>[-\d\.]+),tc=(?P<tc>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "n", "p", "i", "d", "tc", "timestamp"}},
		"n_v_e":                  {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+),n=(?P<n>[-\d\.]+) v=(?P<v>[-\d\.]+),e=(?P<e>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "n", "value", "e", "timestamp"}},
		"r_o_s":                  {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) r=(?P<r>[-\d\.]+)i,o=(?P<o>[-\d\.]+)i,s=(?P<s>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "r", "o", "s", "timestamp"}},
	}
)

// startSyslogServer is a function that starts a syslog server and returns a channel to receive log parts and the server instance.
// The syslog server listens for UDP connections on the specified address.
// It uses the RFC5424 format for log messages.
// The log parts are sent to the provided channel for further processing.
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
	log.Debug().Msg("Syslog server started at: " + listenUDP)
	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			mac := logParts["hostname"].(string)
			if mac == "" { // Skip empty mac addresses
				continue
			} else {
				mac, syslogMetricsPart := func(mac string, logParts map[string]interface{}) (string, map[string]map[string]string) {
					syslogMetricsPart, ok := syslogMetrics.Load(mac) // loading from sync.Map - thread safe

					if !ok {
						return mac, nil // if not found, return empty map
					}

					loadedPart, ok := syslogMetricsPart.(map[string]map[string]string) // type assertion

					if !ok {
						return mac, nil // if not found, return empty map
					}

					if loadedPart == nil {
						loadedPart = make(map[string]map[string]string) // if found but empty, create a new map, at start it will be empty everytime
					}

					if loadedPart["ip"] == nil {
						loadedPart["ip"] = make(map[string]string)

					}

					loadedPart["ip"]["value"] = logParts["client"].(string)

					log.Trace().Msg("Received message from: " + mac)

					for name, pattern := range regexpPatterns {

						reg, err := regexp.Compile(pattern.pattern)
						if err != nil {
							log.Error().Msg("Error compiling regexp: " + err.Error())
							continue
						}

						log.Trace().Msg("Matching pattern: " + name + " for message: " + logParts["message"].(string))

						matches := reg.FindAllStringSubmatch(logParts["message"].(string), -1)
						if matches == nil {
							continue // No matches for this pattern
						}
						var metricName string

						for _, match := range matches {
							// Extract values based on named groups
							for i, field := range pattern.fields {
								if field == "name" {
									metricName = match[i+1]
								} else if match[i+1] != "" && field != "timestamp" { // todo - check if timestamp is needed

									if field == "n" {
										metricName = metricName + "_" + match[i+1]
									}
									if loadedPart[metricName] == nil {
										loadedPart[metricName] = make(map[string]string)
									}
									loadedPart[metricName][field] = match[i+1]
								}
							}
						}

					}
					return mac, loadedPart
				}(mac, logParts)
				if syslogMetricsPart != nil {
					syslogMetrics.Store(mac, syslogMetricsPart) // store the updated map back to sync.Map
				}
			}
		}
	}(channel)

	server.Wait()
}
