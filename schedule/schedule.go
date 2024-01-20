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
	exp1 := viper.GetString("schedule.tasks.expression.tx.update.status")

	entryId, err := c.AddFunc(exp1, PeriodicalUpdateStatusOfUserSendingTransaction)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("init schedule! func PeriodicalUpdateStatusOfUserSendingTransaction entryId: ", entryId)

	c.Start()
}
