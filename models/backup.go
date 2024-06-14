//this file will define everything needed to run backup job on the clipboard file

package models

type ConcreteBackUpStrategy interface {
	do(schedule string)
	initBackup(filepath string)
}

type Backup struct {
	Backup   ConcreteBackUpStrategy
	Schedule string
	Keys     string
}

func (c *Backup) ExecuteBackup() {
	c.Backup.do(c.Schedule)
}

func (c *Backup) SetupBackup(config *Configuration) {
	c.Backup = new(Cronjobber)
	c.Schedule = config.BackUpFrequency
	c.Backup.initBackup(config.DumpFilePath)
}
