package bootstrap

import (
	"github.com/soonio/pupil/bootstrap/internal"
)

// Bootstrap 项目初始化
func Bootstrap(config string) {
	// TIPS 注意顺序
	internal.Workdir()
	internal.Viper(config)
	internal.Redis()
	internal.Cron()
}
