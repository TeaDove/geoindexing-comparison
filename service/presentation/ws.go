package presentation

import (
	"context"
	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog"
	"time"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type DrawDTO struct {
	LegendToPoints map[string][]Point `json:"legendToPoints"`
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
		idx = 0
		err error
	)
	for {
		drawDTO := DrawDTO{LegendToPoints: map[string][]Point{"First Dataset": {{X: float64(idx), Y: float64(idx)}}}}

		err = c.WriteJSON(drawDTO)
		if err != nil {
			zerolog.Ctx(ctx).
				Error().
				Stack().Err(err).
				Msg("failed.to.write")
			break
		}
		time.Sleep(10 * time.Second)
		idx += 1
	}

	zerolog.Ctx(ctx).
		Info().
		Msg("stream.closed")
}
