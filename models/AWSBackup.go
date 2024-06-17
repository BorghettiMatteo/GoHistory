package models

type AWSBackupper struct {
	awsPointer string
	cronJobber *Cronjobber
}

func (a *AWSBackupper) InitBackup(filepath string) {
	a.cronJobber = new(Cronjobber)
	println("per ora ci faccio niente,  ma grazie per il pensiero")
	a.cronJobber.InitCronJobber(filepath)
}

func (a *AWSBackupper) Run() {
	println("eeeeeeeeeeee")
}

func (a *AWSBackupper) Do(schedule string) {
	err := a.cronJobber.ScheduleJob(schedule, a)
	if err != nil {
		println("ERROR: " + "error parsing the schedule")
		panic(err)
	}
}
