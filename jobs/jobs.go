package jobs

import (
	"github.com/baleegh-ud-din/hive/utils"
	"github.com/robfig/cron/v3"
)

var c = cron.New()
var logger = utils.NewLogger()

func StartJobs() {
	logger.Info("🚦 Starting Scheduled Jobs...")
	c.AddFunc("@every 1m", func() {})

	c.Start()
}

func StopJobs() {
	logger.Info("🛑 Stopping Background Jobs...")
	c.Stop()
	logger.Info("⛳️ Stopped Background Jobs.")
}
