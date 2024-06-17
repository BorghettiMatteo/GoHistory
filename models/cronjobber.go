/* questo file contiene la definizione dell'oggetto CronJobber, che sarà la base di ogni backupJob del progetto */

package models

import (
	"github.com/robfig/cron"
)

type Cronjobber struct {
	cronEngine *cron.Cron
	backupJob  InternJob
}

func (c *Cronjobber) GetFilePth() string {
	return c.backupJob.filepath
}

func (c *Cronjobber) InitCronJobber(filepath string) error {
	c.cronEngine = cron.New()
	c.backupJob = InternJob{filepath: filepath}
	// in the future move the checks form runtime to definition of the job, thats why the error is here
	return nil

}

func (c *Cronjobber) ScheduleJob(schedule string, job cron.Job) error {
	csched, err := cron.Parse(schedule)
	if err != nil {
		return err
	}
	c.cronEngine.Schedule(csched, job)
	c.cronEngine.Start()
	return nil
}
