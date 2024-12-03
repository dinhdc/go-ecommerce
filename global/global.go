package global

import (
	"github.com/dinhdc/go-ecommerce/pkg/logger"
	"github.com/dinhdc/go-ecommerce/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
