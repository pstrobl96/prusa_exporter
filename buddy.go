package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

type version struct {
	API      string `json:"api"`
	Server   string `json:"server"`
	Text     string `json:"text"`
	Hostname string `json:"hostname"`
}

type files struct {
	Files []struct {
		Name     string `json:"name"`
		Path     string `json:"path"`
		Display  string `json:"display"`
		Type     string `json:"type"`
		Origin   string `json:"origin"`
		Children []struct {
			Name    string `json:"name"`
			Display string `json:"display"`
			Path    string `json:"path"`
			Origin  string `json:"origin"`
			Refs    struct {
				Resource       string `json:"resource"`
				ThumbnailSmall string `json:"thumbnailSmall"`
				ThumbnailBig   string `json:"thumbnailBig"`
				Download       string `json:"download"`
			} `json:"refs"`
		} `json:"children"`
	} `json:"files"`
}

type printer struct {
	Telemetry struct {
		TempBed    float64 `json:"temp-bed"`
		TempNozzle float64 `json:"temp-nozzle"`
		PrintSpeed int     `json:"print-speed"`
		ZHeight    float64 `json:"z-height"`
		Material   string  `json:"material"`
	} `json:"telemetry"`
	Temperature struct {
		Tool0 struct {
			Actual  float64 `json:"actual"`
			Target  float64 `json:"target"`
			Display float64 `json:"display"`
			Offset  int     `json:"offset"`
		} `json:"tool0"`
		Bed struct {
			Actual float64 `json:"actual"`
			Target float64 `json:"target"`
			Offset int     `json:"offset"`
		} `json:"bed"`
	} `json:"temperature"`
	State struct {
		Text  string `json:"text"`
		Flags struct {
			Operational   bool `json:"operational"`
			Paused        bool `json:"paused"`
			Printing      bool `json:"printing"`
			Cancelling    bool `json:"cancelling"`
			Pausing       bool `json:"pausing"`
			SdReady       bool `json:"sdReady"`
			Error         bool `json:"error"`
			ClosedOnError bool `json:"closedOnError"`
			Ready         bool `json:"ready"`
			Busy          bool `json:"busy"`
		} `json:"flags"`
	} `json:"state"`
}

type job struct {
	State string `json:"state"`
	Job   struct {
		EstimatedPrintTime int `json:"estimatedPrintTime"`
		File               struct {
			Name    string `json:"name"`
			Path    string `json:"path"`
			Display string `json:"display"`
		} `json:"file"`
	} `json:"job"`
	Progress struct {
		Completion    float64 `json:"completion"`
		PrintTime     int     `json:"printTime"`
		PrintTimeLeft int     `json:"printTimeLeft"`
	} `json:"progress"`
}

func getVersion(address string, apiKey string, username string, password string) version {
	resp := accessApi("version", address, apiKey, username, password)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result version

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getFiles(address string, apiKey string, username string, password string) files {
	resp := accessApi("files", address, apiKey, username, password)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result files

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getJob(address string, apiKey string, username string, password string) job {
	resp := accessApi("job", address, apiKey, username, password)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result job

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getPrinter(address string, apiKey string, username string, password string) printer {
	resp := accessApi("printer", address, apiKey, username, password)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result printer

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	return result
}

func getCfg() config {
	cfgFile := os.Getenv("PRUSA_EXPORTER_PRINTERS")
	if cfgFile == "" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pwd)
		cfgFile = pwd + "/printers.yaml"
	}

	log.Println("Using config - " + cfgFile)

	return loadCfg(cfgFile)
}

type buddyCollector struct {
	printerNozzleTemp         *prometheus.Desc
	printerBedTemp            *prometheus.Desc
	printerVersion            *prometheus.Desc // info from version
	printerZHeight            *prometheus.Desc
	printerPrintSpeed         *prometheus.Desc
	printerTargetTempNozzle   *prometheus.Desc
	printerTargetTempBed      *prometheus.Desc
	printerFiles              *prometheus.Desc
	printerPrintTime          *prometheus.Desc
	printerPrintTimeRemaining *prometheus.Desc
	printerPrintProgress      *prometheus.Desc
	printerPrinting           *prometheus.Desc
	printerMaterial           *prometheus.Desc
}

func newBuddyCollector() *buddyCollector {
	return &buddyCollector{
		printerNozzleTemp: prometheus.NewDesc("prusa_nozzle_temperature",
			"Current temperature of printer nozzle in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerBedTemp: prometheus.NewDesc("prusa_bed_temperature",
			"Current temperature of printer bed in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerVersion: prometheus.NewDesc("prusa_version",
			"Return information about printer. This metric contains information mostly about Prusa Link",
			[]string{"printer_address", "printer_model", "printer_name", "printer_api", "printer_server", "printer_text"},
			nil),
		printerZHeight: prometheus.NewDesc("prusa_z_height",
			"Current height of Z",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrintSpeed: prometheus.NewDesc("prusa_print_speed",
			"Current setting of printer speed in percents (%)",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerTargetTempNozzle: prometheus.NewDesc("prusa_nozzle_target_temperature",
			"Target temperature of printer nozzle in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerTargetTempBed: prometheus.NewDesc("prusa_bed_target_temperature",
			"Target temperature of printer bed in Celsius",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerFiles: prometheus.NewDesc("prusa_files",
			"Number of files in storage",
			[]string{"printer_address", "printer_model", "printer_name", "printer_storage"},
			nil),
		printerPrintTimeRemaining: prometheus.NewDesc("prusa_printing_time_remaining",
			"Returns time that remains for completion of current print",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrintProgress: prometheus.NewDesc("prusa_printing_progress",
			"Returns information about completion of current print in percents",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerPrinting: prometheus.NewDesc("prusa_printing",
			"Return information about printing",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
		printerMaterial: prometheus.NewDesc("prusa_material",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path", "printer_filament"},
			nil),
		printerPrintTime: prometheus.NewDesc("prusa_print_time",
			"Returns information about loaded filament. Returns 0 if there is no loaded filament",
			[]string{"printer_address", "printer_model", "printer_name", "printer_job_name", "printer_job_path"},
			nil),
	}
}

func (collector *buddyCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.printerNozzleTemp
	ch <- collector.printerBedTemp
	ch <- collector.printerVersion
	ch <- collector.printerZHeight
	ch <- collector.printerPrintSpeed
	ch <- collector.printerTargetTempNozzle
	ch <- collector.printerTargetTempBed
	ch <- collector.printerFiles
	ch <- collector.printerPrintTime
	ch <- collector.printerPrintTimeRemaining
	ch <- collector.printerPrintProgress
	ch <- collector.printerPrinting
	ch <- collector.printerMaterial

}

func (collector *buddyCollector) Collect(ch chan<- prometheus.Metric) {
	cfg := getCfg()
	for _, s := range cfg.Printers.Password {
		log.Println("Scraping " + s.Address)
		printer := getPrinter(s.Address, "", s.Username, s.Pass)
		files := getFiles(s.Address, "", s.Username, s.Pass)
		version := getVersion(s.Address, "", s.Username, s.Pass)
		job := getJob(s.Address, "", s.Username, s.Pass)
		bedTemp := prometheus.MustNewConstMetric(
			collector.printerBedTemp, prometheus.GaugeValue, // collector
			float64(printer.Temperature.Bed.Actual),                         // value
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path) // labels

		nozzleTemp := prometheus.MustNewConstMetric(
			collector.printerNozzleTemp, prometheus.GaugeValue,
			float64(printer.Temperature.Tool0.Actual),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printProgress := prometheus.MustNewConstMetric(
			collector.printerPrintProgress, prometheus.GaugeValue,
			float64(job.Progress.Completion),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printSpeed := prometheus.MustNewConstMetric(
			collector.printerPrintSpeed, prometheus.GaugeValue,
			float64(printer.Telemetry.PrintSpeed),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printTimeRemaining := prometheus.MustNewConstMetric(
			collector.printerPrintTimeRemaining, prometheus.GaugeValue,
			float64(job.Progress.PrintTimeLeft),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printingMetric := 0
		if job.State == "Printing" {
			printingMetric = 1
		}

		printing := prometheus.MustNewConstMetric(
			collector.printerPrinting, prometheus.GaugeValue,
			float64(printingMetric),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printTime := prometheus.MustNewConstMetric(
			collector.printerPrintTime, prometheus.GaugeValue,
			float64(job.Progress.PrintTime),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		targetTempBed := prometheus.MustNewConstMetric(
			collector.printerTargetTempBed, prometheus.GaugeValue,
			float64(printer.Temperature.Bed.Target),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		targetTempNozzle := prometheus.MustNewConstMetric(
			collector.printerTargetTempNozzle, prometheus.GaugeValue,
			float64(printer.Temperature.Tool0.Target),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		filamentLoaded := 0
		if printer.Telemetry.Material != "---" {
			filamentLoaded = 1
		}

		material := prometheus.MustNewConstMetric(
			collector.printerMaterial, prometheus.GaugeValue,
			float64(filamentLoaded),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.Telemetry.Material)

		printerVersion := prometheus.MustNewConstMetric(
			collector.printerVersion, prometheus.GaugeValue,
			1,
			s.Address, s.Type, s.Name, version.API, version.Server, version.Text)

		printerFiles := prometheus.MustNewConstMetric(
			collector.printerFiles, prometheus.GaugeValue,
			float64(len(files.Files[0].Children)),
			s.Address, s.Type, s.Name, files.Files[0].Display)

		zHeight := prometheus.MustNewConstMetric(
			collector.printerZHeight, prometheus.GaugeValue,
			printer.Telemetry.ZHeight,
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		ch <- bedTemp
		ch <- nozzleTemp
		ch <- printProgress
		ch <- printSpeed
		ch <- printTimeRemaining
		ch <- printing
		ch <- printTime
		ch <- targetTempBed
		ch <- targetTempNozzle
		ch <- material
		ch <- printerVersion
		ch <- printerFiles
		ch <- zHeight

	}
	for _, s := range cfg.Printers.APIKey {
		log.Println("Scraping " + s.Address)
		printer := getPrinter(s.Address, s.Apikey, "", "")
		files := getFiles(s.Address, s.Apikey, "", "")
		version := getVersion(s.Address, s.Apikey, "", "")
		job := getJob(s.Address, s.Apikey, "", "")

		bedTemp := prometheus.MustNewConstMetric(
			collector.printerBedTemp, prometheus.GaugeValue, // collector
			float64(printer.Temperature.Bed.Actual),                         // value
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path) // labels

		nozzleTemp := prometheus.MustNewConstMetric(
			collector.printerNozzleTemp, prometheus.GaugeValue,
			float64(printer.Temperature.Tool0.Actual),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printProgress := prometheus.MustNewConstMetric(
			collector.printerPrintProgress, prometheus.GaugeValue,
			float64(job.Progress.Completion),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printSpeed := prometheus.MustNewConstMetric(
			collector.printerPrintSpeed, prometheus.GaugeValue,
			float64(printer.Telemetry.PrintSpeed),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printTimeRemaining := prometheus.MustNewConstMetric(
			collector.printerPrintTimeRemaining, prometheus.GaugeValue,
			float64(job.Progress.PrintTimeLeft),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printingMetric := 0
		if job.State == "Printing" {
			printingMetric = 1
		}

		printing := prometheus.MustNewConstMetric(
			collector.printerPrinting, prometheus.GaugeValue,
			float64(printingMetric),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		printTime := prometheus.MustNewConstMetric(
			collector.printerPrintTime, prometheus.GaugeValue,
			float64(job.Progress.PrintTime),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		targetTempBed := prometheus.MustNewConstMetric(
			collector.printerTargetTempBed, prometheus.GaugeValue,
			float64(printer.Temperature.Bed.Target),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		targetTempNozzle := prometheus.MustNewConstMetric(
			collector.printerTargetTempNozzle, prometheus.GaugeValue,
			float64(printer.Temperature.Tool0.Target),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		filamentLoaded := 0
		if printer.Telemetry.Material != "---" {
			filamentLoaded = 1
		}

		material := prometheus.MustNewConstMetric(
			collector.printerMaterial, prometheus.GaugeValue,
			float64(filamentLoaded),
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path, printer.Telemetry.Material)

		printerVersion := prometheus.MustNewConstMetric(
			collector.printerVersion, prometheus.GaugeValue,
			1,
			s.Address, s.Type, s.Name, version.API, version.Server, version.Text)

		printerFiles := prometheus.MustNewConstMetric(
			collector.printerFiles, prometheus.GaugeValue,
			float64(len(files.Files[0].Children)),
			s.Address, s.Type, s.Name, files.Files[0].Display)

		zHeight := prometheus.MustNewConstMetric(
			collector.printerZHeight, prometheus.GaugeValue,
			printer.Telemetry.ZHeight,
			s.Address, s.Type, s.Name, job.Job.File.Name, job.Job.File.Path)

		ch <- bedTemp
		ch <- nozzleTemp
		ch <- printProgress
		ch <- printSpeed
		ch <- printTimeRemaining
		ch <- printing
		ch <- printTime
		ch <- targetTempBed
		ch <- targetTempNozzle
		ch <- material
		ch <- printerVersion
		ch <- printerFiles
		ch <- zHeight

	}
}
