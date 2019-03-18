package schedule

import (
	"time"

	"github.com/lexkong/log"
	"github.com/robfig/cron"
)

func Setup() {

	c := cron.New()
	c.AddFunc("*/20 * * * * *", func() {
		log.Infof("test cron, time: %d ", time.Now().Unix())
	})

	c.Start()
}
