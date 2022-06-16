package utils

import (
	"os"
	"runtime"
)

// Touch 模仿Linux命令touch一个新文件
func Touch(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if err != nil {
		return err
	}

	return f.Close()
}

func Or[v any](assert bool, yes, no v) v {
	if assert {
		return yes
	} else {
		return no
	}
}

// Fn 获取方法名称
func Fn(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
