package models

import (
	"io"
	"os"
	"time"

	"github.com/robfig/cron"
)

type Cronjobber struct {
	cronEngine *cron.Cron
	backupJob  InternJob
}

func (c *Cronjobber) Run() {
	//execute backup strategy copying the dump file and renaming it with the date of the backup
	_, err := os.Stat(c.backupJob.filepath)
	if err != nil {
		println("ERROR: error reading filedump for backup")
		panic(err)
	}

	dest, err := os.Create(time.Now().String())
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	//open file to read content

	source, err := os.Open(c.backupJob.filepath)
	if err != nil {
		panic(err)
	}
	defer source.Close()
	_, err = io.Copy(dest, source)
	if err != nil {
		panic(err)
	}
}

func (c *Cronjobber) initBackup(filepath string) {
	c.cronEngine = cron.New()
	c.backupJob = InternJob{filepath: filepath}
}

func (c *Cronjobber) do(schedule string) {
	csched, err := cron.Parse(schedule)
	if err != nil {
		println("ERROR: " + "error parsing the schedule")
		panic(err)
	}
	c.cronEngine.Schedule(csched, c)
	c.cronEngine.Start()
}
