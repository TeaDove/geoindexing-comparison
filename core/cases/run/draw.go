package main

import (
	"bytes"
	"context"
	"fmt"
	"geoindexing_comparison/core/cases"
	"geoindexing_comparison/core/ds_supplier"
	"geoindexing_comparison/core/utils"
	"image"
	"image/jpeg"
	"os"
)

var supplier *ds_supplier.Supplier

func init() {
	var err error
	supplier, err = ds_supplier.New()
	utils.Check(err)
}

func saveImageLocally(imgName string, imgFile []byte) {
	_ = os.Remove(imgName)

	img, _, err := image.Decode(bytes.NewReader(imgFile))
	utils.Check(err)

	out, err := os.Create(imgName)
	utils.Check(err)

	err = jpeg.Encode(out, img, nil)
	utils.Check(err)

	utils.CloseOrLog(out)
}

func drawResultsForTask(taskName string, results []cases.Result) {
	values := make(map[string]map[int]float64, 10)
	for _, result := range results {
		collectionName := result.RunCase.Collection().Name()
		_, ok := values[collectionName]
		if !ok {
			values[collectionName] = make(map[int]float64, 10)
		}

		values[collectionName][result.Amount] = float64(result.Durs.Avg().Nanoseconds())
	}

	ctx := context.Background()
	lineplotImg, err := supplier.DrawLinePlot(ctx, &ds_supplier.DrawLinePlotInput{
		DrawInput: ds_supplier.DrawInput{
			Title:  taskName,
			XLabel: "amount",
			YLabel: "nanoseconds",
		},
		Values: values,
	})
	utils.Check(err)

	saveImageLocally(fmt.Sprintf(".%s.jpeg", taskName), lineplotImg)
}

func drawResults(results []cases.Result) {
	taskToResults := make(map[string][]cases.Result, 5)

	for _, result := range results {
		taskName := result.RunCase.Task.Name()
		_, ok := taskToResults[taskName]
		if !ok {
			taskToResults[taskName] = make([]cases.Result, 0, 10)
		}

		taskToResults[taskName] = append(taskToResults[taskName], result)
	}

	for taskName, taskResults := range taskToResults {
		drawResultsForTask(taskName, taskResults)
	}
}
