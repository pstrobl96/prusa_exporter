package cmd

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/rs/zerolog/log"
)

var (
	configFile = kingpin.Flag("config.file", "Configuration file for prusa_exporter.").Default("config.yml").ExistingFile()
)

func Run() {
	log.Info().Msg("Prusa exporter starting")
}
