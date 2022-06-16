package app

import (
	"github.com/soonio/pupil/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Home = "" // 项目根目录

	Config *config.Config // 项目配置

	DB    *gorm.DB // 存储多个DB对象
	Redis *redis.Client
)
