package models

import (
	"os"
	"time"
)

type FileSystemBackup struct {
	cronJobber *Cronjobber
}

func (f *FileSystemBackup) Run() {
	//execute backup strategy copying the dump file and renaming it with the date of the backup
	_, err := os.Stat(f.cronJobber.GetFilePth())
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

	toWrite, err := os.ReadFile(f.cronJobber.GetFilePth())
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

func (f *FileSystemBackup) InitBackup(filepath string) {
	f.cronJobber = new(Cronjobber)
	f.cronJobber.InitCronJobber(filepath)
}

func (f *FileSystemBackup) Do(schedule string) {
	err := f.cronJobber.ScheduleJob(schedule, f)
	if err != nil {
		println("ERROR: " + "error parsing the schedule")
		panic(err)
	}
}
