package internal

import (
	"github.com/soonio/pupil/app"
	"os"
)

func Workdir() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	app.Home = dir
}
