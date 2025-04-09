package presentation

//func (r *Presentation) wsHandle(c *websocket.Conn) {
//	ctx := fiber_utils.GetLogCtx(c)
//
//	// TODO move to settings
//	ctx, cancel := context.WithTimeout(ctx, time.Hour)
//	defer cancel()
//
//	zerolog.Ctx(ctx).
//		Info().
//		Msg("new.ws.stream")
//
//	var (
//		idx = 0
//		err error
//	)
//	for {
//		resultsLen := len(r.results)
//		if idx >= resultsLen {
//			time.Sleep(200 * time.Millisecond)
//			continue
//		}
//
//		points := make([]Point, 0, 10)
//		for _, res := range r.results[idx:resultsLen] {
//			points = append(points, Point{Chart: fmt.Sprintf("%s %s", "", res.Task), Dataset: res.Index, X: float64(res.Amount), Y: float64(res.Durs.Avg())})
//		}
//
//		err = c.WriteJSON(points)
//		if err != nil {
//			zerolog.Ctx(ctx).
//				Error().
//				Stack().Err(err).
//				Msg("failed.to.write")
//			break
//		}
//		idx = resultsLen
//	}
//
//	zerolog.Ctx(ctx).
//		Info().
//		Msg("stream.closed")
//}
