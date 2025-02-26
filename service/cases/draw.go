package cases

//func drawResultsForTask(task tasks.Task, caseName string, results []Result) {
//	values := make([][]any, 0, 10)
//
//	for _, result := range results {
//		durWithSE := result.Durs.AvgWithSE()
//		values = append(values, []any{result.CollectionName, result.Amount, durWithSE[0]})
//		values = append(values, []any{result.CollectionName, result.Amount, durWithSE[1]})
//		values = append(values, []any{result.CollectionName, result.Amount, durWithSE[2]})
//	}
//
//	ctx := context.Background()
//	lineplotImg, err := supplier.DrawLinePlot(ctx, &ds_supplier.DrawLinePlotInput{
//		DrawInput: ds_supplier.DrawInput{
//			Title:  fmt.Sprintf("%s\n%s", task.Name(), task.Description()),
//			XLabel: "Кол-во точек в структуре",
//			YLabel: "Время выполнение задачи в наносекундах",
//		},
//		Values: values,
//	})
//	utils.Check(err)
//
//	saveImageLocally(fmt.Sprintf("extra/time_%s_%s.jpeg", caseName, task.Filename()), lineplotImg)
//
//	values = make([][]any, 0, 10)
//	for _, result := range results {
//		values = append(values, []any{result.CollectionName, result.Amount, result.Mems.Avg()})
//	}
//
//	ctx = context.Background()
//	lineplotImg, err = supplier.DrawLinePlot(ctx, &ds_supplier.DrawLinePlotInput{
//		DrawInput: ds_supplier.DrawInput{
//			Title:  fmt.Sprintf("%s\n%s", task.Name(), task.Description()),
//			XLabel: "Кол-во точек в структуре",
//			YLabel: "Затраты памяти на выполнение задачи в байтах",
//		},
//		Values: values,
//	})
//	utils.Check(err)
//
//	saveImageLocally(fmt.Sprintf("extra/mem_%s_%s.jpeg", caseName, task.Filename()), lineplotImg)
//}
