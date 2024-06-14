package models

import (
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
		return
	}

	dest, err := os.Create(time.Now().Format("20060102150405"))
	if err != nil {
		println("ERROR: creating file for backup " + err.Error())
		return
	}
	defer dest.Close()

	//open file to read content

	toWrite, err := os.ReadFile(c.backupJob.filepath)
	if err != nil {
		println("ERROR: reading file for backup " + err.Error())
		return
	}

	//compress the backup
	exitdump, err := CreateCompressedLog(toWrite)
	if err != nil {
		println("ERROR: " + "error during compression")
		return
	}

	// write to file
	_, err = dest.Write(exitdump)

	if err != nil {
		println("ERROR: error writing logs to file" + err.Error())
		return
	}

	println("SUCCESS: successfully backupped my clipboard history, hooray")
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
