package utils

import "github.com/rs/zerolog/log"

func Check(err error) {
	if err != nil {
		log.Panic().Str("status", "check.failed").Stack().Err(err).Send()
	}
}
