package initialize

import (
	"github.com/dinhdc/go-ecommerce/global"
	"github.com/dinhdc/go-ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
