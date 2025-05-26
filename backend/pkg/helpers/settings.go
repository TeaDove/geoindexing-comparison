package helpers

import (
	"github.com/teadove/teasutils/utils/settings_utils"
)

type settings struct {
	ManagerURL string `env:"MANAGER_URL" envDefault:"http://127.0.0.1:8000"`
	SqlitePath string `env:"SQLITE_PATH" envDefault:"../.data/db.sqlite"`
}

var Settings = settings_utils.MustGetSetting[settings]("GEO_") // nolint: gochecknoglobals // it's ok
