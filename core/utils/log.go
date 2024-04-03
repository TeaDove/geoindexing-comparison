package utils

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
)

func SendInterface(values ...any) {
	arr := zerolog.Arr()
	for _, value := range values {
		arr.Dict(zerolog.Dict().Interface(GetType(value), value))
	}

	log.Info().Array("items", arr).Str("status", "logging.struct").Send()
}

func CloseOrLog(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Error().
			Stack().
			Err(errors.WithStack(err)).
			Str("status", "failed.to.close").
			Send()
	}
}
