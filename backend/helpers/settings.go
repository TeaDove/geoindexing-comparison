package helpers

import (
	"github.com/teadove/teasutils/utils/logger_utils"
	"github.com/teadove/teasutils/utils/settings_utils"
)

type settings struct {
	ManagerURL string `env:"MANAGER_URL" envDefault:"http://127.0.0.1:8000"`
}

var Settings = settings_utils.MustGetSetting[settings]( //nolint: gochecknoglobals // it's ok
	logger_utils.NewLoggedCtx(),
	"GEO_",
)
