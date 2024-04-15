package cases

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"geoindexing_comparison/core/cases/tasks"
	"geoindexing_comparison/core/ds_supplier"
	"geoindexing_comparison/core/utils"
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

func drawResultsForTask(task tasks.Task, caseName string, results []Result) {
	values := make([][]any, 0, 10)
	for _, result := range results {
		durWithSE := result.Durs.AvgWithSE()
		values = append(values, []any{result.CollectionName, result.Amount, durWithSE[0]})
		values = append(values, []any{result.CollectionName, result.Amount, durWithSE[1]})
		values = append(values, []any{result.CollectionName, result.Amount, durWithSE[2]})
	}

	ctx := context.Background()
	lineplotImg, err := supplier.DrawLinePlot(ctx, &ds_supplier.DrawLinePlotInput{
		DrawInput: ds_supplier.DrawInput{
			Title:  fmt.Sprintf("%s\n%s", task.Name(), task.Description()),
			XLabel: "Кол-во точек в структуре",
			YLabel: "Время выполнение задачи в наносекундах",
		},
		Values: values,
	})
	utils.Check(err)

	saveImageLocally(fmt.Sprintf("extra/%s_%s.jpeg", caseName, task.Filename()), lineplotImg)
}
