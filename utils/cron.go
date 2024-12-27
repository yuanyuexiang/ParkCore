package utils

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

// ExecCornFunc ExecCornFunc
func ExecCornFunc(spec string, function func()) (cron *cron.Cron, err error) {
	cron = newWithSeconds()
	if _, err = cron.AddFunc(spec, function); err != nil {
		fmt.Println(err)
	} else {
		cron.Start()
	}
	return
}
