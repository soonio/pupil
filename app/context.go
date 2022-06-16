package app

import (
	"github.com/soonio/pupil/config"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Home   string         // 项目根目录
	Config *config.Config // 配置
	DB     *gorm.DB       // DB对象
	Redis  *redis.Client  // reid
)
