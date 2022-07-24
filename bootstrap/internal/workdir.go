package internal

import (
	"github.com/soonio/pupil/app"
	"os"
)

// Workdir 初始化工作目录
func Workdir() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	app.Home = dir
}
