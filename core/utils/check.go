package utils

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func Check(err error) {
	if err != nil {
		FancyPanic(err)
	}
}

func FancyPanic(err error) {
	log.Panic().Str("status", "check.failed").Stack().Err(errors.WithStack(err)).Send()
}
