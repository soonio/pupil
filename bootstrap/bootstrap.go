package bootstrap

import (
	"github.com/soonio/pupil/bootstrap/internal"
)

// Bootstrap 项目初始化
func Bootstrap(config string) {
	// TIPS 注意顺序
	// 未使用recover()，异常信息会直接抛出，便于在启动时即使发现问题
	internal.Workdir()
	internal.Viper(config)
	internal.Redis()
	internal.Cron()
}
