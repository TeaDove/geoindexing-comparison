package presentation

import (
	"context"
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog"
	"time"
)

type DrawDTO struct {
	PlotName string `json:"plotName"`
}

func (r *Presentation) wsHandle(c *websocket.Conn) {
	ctx := r.mustGetLogContext(c)

	// TODO move to settings
	ctx, cancel := context.WithTimeout(ctx, time.Hour)
	defer cancel()

	zerolog.Ctx(ctx).
		Info().
		Msg("new.ws.stream")

	var (
		msg []byte
		err error
	)
	for {
		drawDTO := DrawDTO{"Test"}
		msg, err = json.Marshal(drawDTO)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Stack().Err(err).
				Interface("drawDTO", drawDTO).
				Msg("failed.to.marshal")
			break
		}

		err = c.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Stack().Err(err).
				Str("msg", string(msg)).
				Msg("failed.to.write")
			break
		}
		time.Sleep(10 * time.Second)
	}

	zerolog.Ctx(ctx).
		Info().
		Msg("stream.closed")
}
