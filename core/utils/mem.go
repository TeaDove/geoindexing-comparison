package utils

import (
	"github.com/rs/zerolog/log"
	"runtime"
)

func LogMemUsage() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	log.Info().
		Str("status", "perfstats").
		Float64("stop.the.world.ms", ToFixed(float64(m.PauseTotalNs)/1024/1024, 2)).
		Float64("heap.alloc.mb", ToMega(m.HeapAlloc)).
		Float64("cum.heap.alloc.mb", ToMega(m.TotalAlloc)).
		Float64("heap.alloc.count.k", ToKilo(m.HeapObjects)).
		Float64("stack.in.use.mb", ToMega(m.StackInuse)).
		Float64("total.sys.mb", ToMega(m.Sys)).
		Float64("gc.cpu.percent", ToFixed(m.GCCPUFraction*100, 4)).
		Uint32("gc.cycles", m.NumGC).
		Int("gorutine.count", runtime.NumGoroutine()).
		Send()
}