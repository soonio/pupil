package api

import "log"

func ding(err error) error {

	// 写日志
	// 钉钉、微信、邮件等告警

	log.Fatalf(err.Error())

	return err
}
