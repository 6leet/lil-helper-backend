package scheduler

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

var Cron *gocron.Scheduler

func init() {
	Cron = gocron.NewScheduler()
	fmt.Println("init cron")
}
