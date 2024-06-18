package jobmodels

import (
	"main/models"
	"os"
)

type RestoreBackupFileSystem struct{}

func initConfig() *models.Configuration {
	var config = new(models.Configuration)
	config.LoadConfiguration()
	return config
}

func (r *RestoreBackupFileSystem) RunJob() {
	// first find the latest backup
	backupPath := initConfig().BackUpStoragePath
	files, err := os.ReadDir(backupPath)
	if err != nil {
		println("ERROR: error reading content of" + backupPath + " please read the error following: " + err.Error())
		return
	}
	println(os.Getwd())
	// l'ultimo file è quello più recente
	if len(files) > 1 {
		println(files[len(files)-1].Name())
	}
}
