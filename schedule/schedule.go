package schedule

import (
	log "github.com/cihub/seelog"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func init() {
	InitSchedule()
}

func InitSchedule() {
	c := cron.New()
	exp1 := viper.GetString("schedule.tasks.expression.1")
	//exp2 := viper.GetString("schedule.tasks.expression.2")

	entryId, err := c.AddFunc(exp1, func1)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("init schedule! func1 entryId: ", entryId)

	c.Start()
	select {}
}

func func1() {
	PeriodicalUpdateStatusOfUserSendingTransaction()
}
