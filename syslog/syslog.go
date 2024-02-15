package syslog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

var syslogData = make(map[string]map[string]string)

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

func handleMetrics(listenUDP string) {
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
