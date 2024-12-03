package initialize

import (
	"github.com/dinhdc/go-ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Check init logger", zap.Int("line", 11))
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
