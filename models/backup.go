//this file will define everything needed to run backup job on the clipboard file

package models

import (
	"bytes"
	"compress/gzip"
)

type ConcreteBackUpStrategy interface {
	Do(schedule string)
	InitBackup(filepath string)
}

type InternJob struct {
	filepath string
}

type Backup struct {
	Backup   ConcreteBackUpStrategy
	Schedule string
	Keys     string
}

func (c *Backup) ExecuteBackup() {
	c.Backup.Do(c.Schedule)
}

func (c *Backup) SetupBackup(config *Configuration) {
	switch config.BackUSptrategy {
	case "cron":
		c.Backup = new(FileSystemBackup)
	case "aws":
		c.Backup = new(AWSBackupper)
	}
	c.Schedule = config.BackUpFrequency
	c.Backup.InitBackup(config.DumpFilePath)
}

func CreateCompressedLog(dump []byte) ([]byte, error) {
	var bytesBuff bytes.Buffer
	writer := gzip.NewWriter(&bytesBuff)
	_, err := writer.Write(dump)
	defer writer.Close()
	if err != nil {
		return nil, err
	}
	return bytesBuff.Bytes(), nil
}
