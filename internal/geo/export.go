package geo

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"geoindexing_comparison/utils"
	"github.com/google/uuid"
	"os"
	"os/exec"
)

func (r Points) MustExport(filename string) {
	csvFile, err := os.Create(filename)
	defer csvFile.Close()
	utils.Check(err)

	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.Write([]string{"id", "lat", "lon", "color"})
	utils.Check(err)

	for _, point := range r {
		err = csvwriter.Write([]string{point.ID.String(), fmt.Sprintf("%f", point.Lat), fmt.Sprintf("%f", point.Lon), string(point.Color)})
		utils.Check(err)
	}
	csvwriter.Flush()
}

func (r Points) Draw() {
	filename := fmt.Sprintf("/tmp/%s.csv", uuid.New().String())
	r.MustExport(filename)
	fmt.Println(filename)
	cmd := exec.Command("./draw.py", filename, "points.png")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()

	fmt.Println(outb.String())
	fmt.Println(errb.String())
	err = os.Remove(filename)
	utils.Check(err)
}
