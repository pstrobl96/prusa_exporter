package lineprotocol

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2/format"
)

type point struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{} // Use interface{} to handle different field types
	Timestamp   time.Time
}

func process(data format.LogParts, received time.Time, prefix string) {
	mac, ip, timestamp, err := processTimestamp(data, received)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error processing timestamp: %v", err))
		return
	}
	log.Debug().Msg(fmt.Sprintf("Processing data for printer %s with timestamp %d", mac, timestamp))
	metrics, err := processMessage(data["message"].(string), timestamp, mac, prefix, ip)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error processing message: %v", err))
		return
	}

	for _, line := range metrics {
		point, err := parseLineProtocol(line)
		if err != nil {
			fmt.Printf("Error parsing '%s': %v\n", line, err)
			continue
		}
		fmt.Printf("Parsed: %+v\n", point)
		fmt.Printf("  Measurement: %s\n", point.Measurement)
		fmt.Printf("  Tags: %v\n", point.Tags)
		fmt.Printf("  Fields: %v\n", point.Fields)
		fmt.Println("---")
	}
}

// processTimestamp returns the MAC address and timestamp from the ingested data
// it is basically used for the synchronization of time between handler and the printer
func processTimestamp(data format.LogParts, received time.Time) (string, string, int64, error) {
	mac, ok := data["hostname"].(string)
	if !ok {
		return "", "", 0, fmt.Errorf("mac is not an string")
	}

	ip, ok := data["client"].(string)
	if !ok {
		return "", "", 0, fmt.Errorf("ip is not an string")
	}

	return mac, ip, received.UnixNano(), nil
}

func processMessage(message string, timestamp int64, mac string, prefix string, ip string) ([]string, error) {
	messageSplit := strings.Split(message, "\n")

	if len(messageSplit) == 0 {
		return nil, fmt.Errorf("message is empty")
	}

	firstMessage, err := parseFirstMessage(messageSplit[0])

	if err != nil {
		return nil, fmt.Errorf("error parsing first message: %v", err)
	}

	messageSplit = append(messageSplit[1:], firstMessage)

	for i, line := range messageSplit {
		splitted := strings.Split(line, " ")
		delta, err := strconv.ParseInt(splitted[len(splitted)-1], 10, 64)
		if err != nil {
			log.Error().Msg("Expected error while parsing time delta for metric: " + splitted[0] + " error:" + err.Error())
			continue
		}
		splitted[len(splitted)-1] = strconv.FormatInt(timestamp+delta, 10)
		log.Trace().Msg("Processing timestamps for " + message)
		splitted, err = updateMetric(splitted, prefix, mac, ip)
		if err != nil {
			log.Error().Msg("Expected error while adding mac label for metric: " + splitted[0] + " error:" + err.Error())
			continue
		}
		messageSplit[i] = strings.Join(splitted, " ")
	}
	return messageSplit, nil
}

func parseFirstMessage(message string) (string, error) {
	splitted := strings.Split(message, " ")
	if len(splitted) == 0 {
		return "", fmt.Errorf("splitted message is empty")
	}
	firstMsg := splitted[1:]
	return strings.Join(firstMsg, " "), nil
}

func updateMetric(splitted []string, prefix string, mac string, ip string) ([]string, error) {
	if len(splitted) == 0 {
		return nil, fmt.Errorf("splitted message is empty")
	}

	splitted[0] = fmt.Sprintf("%s%s,mac=%s,ip=%s", prefix, splitted[0], mac, strings.Split(ip, ":")[0])
	return splitted, nil
}

func newPoint() *point {
	return &point{
		Tags:   make(map[string]string),
		Fields: make(map[string]interface{}),
	}
}

func parseLineProtocol(line string) (*point, error) {
	p := newPoint()

	parts := strings.Split(line, " ")
	if len(parts) < 2 || len(parts) > 3 {
		return nil, fmt.Errorf("invalid line protocol format: %s", line) // this happens when printer sends error message
	}

	measurementTags := parts[0]
	measurementTagParts := strings.Split(measurementTags, ",")
	p.Measurement = measurementTagParts[0]

	for i := 1; i < len(measurementTagParts); i++ {
		tag := measurementTagParts[i]
		tagParts := strings.SplitN(tag, "=", 2)
		if len(tagParts) != 2 {
			return nil, fmt.Errorf("invalid tag format: %s", tag)
		}
		p.Tags[tagParts[0]] = tagParts[1]
	}

	fieldStr := parts[1]
	fieldParts := strings.Split(fieldStr, ",")
	for _, field := range fieldParts {
		kv := strings.SplitN(field, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid field format: %s", field)
		}
		key := kv[0]
		val := kv[1]

		// parsing metrics as different data types

		if strings.HasSuffix(val, "i") { // Integer
			if iVal, err := strconv.ParseInt(val[:len(val)-1], 10, 64); err == nil {
				p.Fields[key] = iVal
				continue
			}
		}

		if bVal, err := strconv.ParseBool(val); err == nil { // boolean
			p.Fields[key] = bVal
			continue
		}

		if fVal, err := strconv.ParseFloat(val, 64); err == nil { // float
			p.Fields[key] = fVal
			continue
		}

		if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") { // string
			p.Fields[key] = val[1 : len(val)-1]
			continue
		}

		// fallback
		p.Fields[key] = val
	}

	return p, nil
}
