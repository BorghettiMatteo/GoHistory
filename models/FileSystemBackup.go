package models

import (
	"os"
	"path"
	"time"
)

type FileSystemBackup struct {
	cronJobber       *Cronjobber
	backUpFolderPath string
}

func (f *FileSystemBackup) Run() {
	//execute backup strategy copying the dump file and renaming it with the date of the backup
	_, err := os.Stat(f.cronJobber.GetFilePth())
	if err != nil {
		println("ERROR: error reading filedump for backup")
		return
	}
	_, err = os.Stat(f.backUpFolderPath)
	if err != nil {
		println("Error searching for backupFolder: " + err.Error())
	}
	if os.IsNotExist(err) {
		//create
		mkdirErr := os.Mkdir(f.backUpFolderPath, 0755)
		if mkdirErr != nil {
			println("ERROR: creating backup folder: " + mkdirErr.Error())
		}
	}
	dest, err := os.Create(path.Join(f.backUpFolderPath, time.Now().Format("20060102150405")))
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

func (f *FileSystemBackup) InitBackup(config *Configuration) {
	f.cronJobber = new(Cronjobber)
	f.cronJobber.InitCronJobber(config.DumpFilePath)
	f.backUpFolderPath = config.BackUpStoragePath

}

func (f *FileSystemBackup) Do(schedule string) {
	err := f.cronJobber.ScheduleJob(schedule, f)
	if err != nil {
		println("ERROR: " + "error parsing the schedule")
		panic(err)
	}
}
