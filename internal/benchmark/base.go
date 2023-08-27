package benchmark

import (
	"geoindexing_comparison/addapter"
	"geoindexing_comparison/generator"
)

type Benchmark struct {
	Generator  *generator.Generator
	Collection addapter.Collection
}
