//this file will define everything needed to run backup job on the clipboard file

package models

import (
	"bytes"
	"compress/gzip"
)

type ConcreteBackUpStrategy interface {
	Do(schedule string)
	InitBackup(c *Configuration)
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
	case "filesystem":
		c.Backup = new(FileSystemBackup)
	case "aws":
		c.Backup = new(AWSBackupper)
	}
	c.Schedule = config.BackUpFrequency
	c.Backup.InitBackup(config)
}

func CreateCompressedLog(dump []byte) ([]byte, error) {
	var bytesBuff bytes.Buffer
	writer := gzip.NewWriter(&bytesBuff)
	_, err := writer.Write(dump)
	writer.Flush()
	defer writer.Close()
	if err != nil {
		return nil, err
	}
	return bytesBuff.Bytes(), nil
}
