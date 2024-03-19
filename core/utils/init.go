package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime/debug"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(512 * 1024 * 1024)
}
