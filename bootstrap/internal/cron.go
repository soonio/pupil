package internal

import (
	"github.com/soonio/pupil/app/logic"

	"github.com/robfig/cron/v3"
)

// Cron 初始化定时任务
func Cron() {

	var jobs = []struct {
		Spec string
		Fn   func()
	}{
		{Spec: "1 0 * * *", Fn: logic.Clock.ZeroTime},
	}

	c := cron.New()

	for _, job := range jobs {
		_, err := c.AddFunc(job.Spec, job.Fn)
		if err != nil {
			panic(err)
		}
	}

	go c.Start()
}
