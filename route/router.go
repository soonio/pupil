package route

import (
	"github.com/soonio/pupil/app/api"
	"github.com/soonio/pupil/pkg/http"
)

// Initialize 初始化路由
func Initialize() {
	e := http.Server()

	e.GET("/version", api.Home.Version)
}
