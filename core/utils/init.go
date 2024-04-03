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
	gcMemLimit := int64(3 * 1024 * 1024 * 1024)

	debug.SetMemoryLimit(gcMemLimit)
	log.Info().
		Str("status", "gc.set").
		Float64("mem.limit.gb", ToFixed(ToGiga(gcMemLimit), 2)).
		Str("gc", "disabled").
		Send()
}
