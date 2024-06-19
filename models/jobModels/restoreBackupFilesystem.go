package jobmodels

import (
	"bytes"
	"compress/gzip"
	"main/models"
	"os"
	"path"
)

type RestoreBackupFileSystem struct{}

func initConfig() *models.Configuration {
	var config = new(models.Configuration)
	config.LoadConfiguration()
	return config
}

func writeNewDumpFile(gzReder *gzip.Reader) error {
	_, err := os.Stat(initConfig().DumpFilePath)
	//handle gunzip bytes
	var resB bytes.Buffer

	//every time I read the files, i got an EOF, but the file seems ok
	resB.ReadFrom(gzReder)
	if os.IsExist(err) {
		//If the file exists, override it
		fileDescriptor, err := os.OpenFile(initConfig().DumpFilePath, os.O_WRONLY, 0775)
		if err != nil {
			println("ERROR: error opening dump for overriding " + err.Error())
			return err
		}
		fileDescriptor.Close()

		//write to file
		_, err = fileDescriptor.Write(resB.Bytes())
		if err != nil {
			println("ERROR: error writing dump file during backup restore")
			return err
		}
	} else {
		// In the case in which the does not file exist
		dest, err := os.Create(initConfig().DumpFilePath)
		if err != nil {
			println("ERROR: error creating new dump file after restore " + err.Error())
			return err
		}
		defer dest.Close()
		_, err = dest.Write(resB.Bytes())
		if err != nil {
			println("ERROR: generic error, I'm bored " + err.Error())
			return err
		}
	}
	return nil
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
		file := files[len(files)-1]
		gzFileDescriptor, err := os.Open(path.Join(backupPath, file.Name()))
		if err != nil {
			println("ERROR: error opening backup file " + err.Error())
			return
		}
		//defer gzFileDescriptor.Close()

		//create gzipper reader
		gzReader, err := gzip.NewReader(gzFileDescriptor)
		if err != nil {
			println("ERROR creating reader ", err.Error())
			return
		}
		restoreError := writeNewDumpFile(gzReader)
		if restoreError != nil {
			println("RESTORE ERROR: restore error: " + restoreError.Error())
			return
		}
		defer gzReader.Close()

	}
}
