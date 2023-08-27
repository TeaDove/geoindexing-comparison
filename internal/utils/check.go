package utils

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func Check(err error) {
	if err != nil {
		log.Panic().Str("status", "check.failed").Stack().Err(errors.WithStack(err)).Send()
	}
}
