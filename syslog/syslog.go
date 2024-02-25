package syslog

import (
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

type metrics struct {
	name  string
	value string
}

type patterns struct {
	pattern string
	fields  []string
}

var (
	// mac_address:
	//   metric_name:
	//     metric: string
	//     value: string
	syslogMetrics = make(map[string]map[string][]metrics)

	regexpPatterns = map[string]patterns{
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
		"sent":                   {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) sent=(?P<sent>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "sent", "timestamp"}},
		"recv":                   {pattern: `(?P<name>\w+[0-9]*[a-zA-Z]+) recv=(?P<recv>[-\d\.]+) (?P<timestamp>\d+)`, fields: []string{"name", "recv", "timestamp"}},
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
	log.Debug().Msg("Syslog server started at: " + listenUDP)

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			mac := logParts["hostname"].(string)
			if mac == "" { // Skip empty mac addresses
				continue
			} else {
				clientIP := strings.Split(logParts["client"].(string), ":")[0] // getting rid of port and leaving only ip address
				port := strings.Split(logParts["client"].(string), ":")[1]     // getting rid of port and leaving only ip address

				if syslogMetrics[mac] == nil {
					syslogMetrics[mac] = make(map[string][]metrics)
				} // Initialize map for ip address if it doesn't exist - is it unique? No. Is it a problem? No. Is it experimental? Yes.

				syslogMetrics[mac]["ip"] = append(syslogMetrics[mac]["ip"], metrics{name: "ip", value: clientIP})
				syslogMetrics[mac]["port"] = append(syslogMetrics[mac]["port"], metrics{name: "port", value: port})

				log.Debug().Msg("Received message from: " + mac)

				for name, pattern := range regexpPatterns {
					reg, err := regexp.Compile(pattern.pattern)
					if err != nil {
						log.Error().Msg("Error compiling regexp: " + err.Error())
						return
					}

					log.Debug().Msg("Matching pattern: " + name + " for message: " + logParts["message"].(string))

					matches := reg.FindAllStringSubmatch(logParts["message"].(string), -1)
					if matches == nil {
						log.Debug().Msg("No matches for pattern: " + name)
						continue // No matches for this pattern
					}
					//v, x, y, z, timestamp, free, total, axis, sens, period, speed, last, sent, recv, n, t, m, u, a, f, ax, ok, desc, st, r, ri, sp, e, p, i, d, tc, as, fe, rs, ae, reg, regn, pwm, measured, fan, state, n, v

					var metricName string

					for _, match := range matches {
						// Extract values based on named groups
						for i, field := range pattern.fields {
							if field == "name" {
								metricName = match[i+1]
							} else if match[i+1] != "" {
								syslogMetrics[mac][metricName] = append(syslogMetrics[mac][metricName], metrics{name: field, value: match[i+1]})
							}
						}
					}
				}
			}
		}
	}(channel)

	server.Wait()
}
