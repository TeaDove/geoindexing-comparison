package utils

import (
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/rs/zerolog/log"
)

type MemCatcher struct {
	startedMemUsage uint64
}

func StartMemCatching() *MemCatcher {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	return &MemCatcher{startedMemUsage: m.TotalAlloc}
}

func (r *MemCatcher) StopMemCatching() uint64 {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	return m.TotalAlloc - r.startedMemUsage
}

func LogMemUsage() {
	samples := make([]metrics.Sample, 1)
	samples[0].Name = "/memory/classes/total:bytes"

	metrics.Read(samples)
	totalBytes := samples[0].Value.Uint64()

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
		Float64("total.mem.mb", ToMega(totalBytes)).
		Int("gorutine.count", runtime.NumGoroutine()).
		Send()
}

func SpamLogMemUsage() {
	for {
		LogMemUsage()
		time.Sleep(1 * time.Second)
	}
}
