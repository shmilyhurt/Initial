package task

import (
	"Initial/dao"
	"github.com/robfig/cron/v3"
)

func InitTask() {
	c := cron.New()
	_, _ = c.AddFunc("@every 10s", FayForPro)
	c.Start()
}

func FayForPro() {
	dao.CreateDashProStatus()
	dao.CreateDashProType()
	dao.CreateDashAll()
}
