package tasks

//type Insert struct{}
//
//func (r *Insert) Run(index index.IndexImpl, amount uint64) time.Duration {
//	durs := make([]time.Duration, 0, amount/10)
//
//	for range amount / 10 {
//		runtime.GC()
//
//		durs = append(
//			durs,
//			index.InsertTimed(generator.DefaultGenerator.Point(&generator.DefaultInput)),
//		)
//
//		runtime.GC()
//	}
//
//	return stats.NewArray(durs).Avg()
//}
