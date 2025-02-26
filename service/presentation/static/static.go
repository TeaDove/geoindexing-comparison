package static

import (
	"embed"
	"github.com/pkg/errors"
)

//go:embed *
var FS embed.FS

func init() {
	files, err := FS.ReadDir(".")
	if err != nil {
		panic(errors.Wrap(err, "failed to read static directory"))
	}

	if len(files) == 0 {
		panic(errors.New("no static files found"))
	}
}
