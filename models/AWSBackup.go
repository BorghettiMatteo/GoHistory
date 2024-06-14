package models

type AWSBackupper struct {
	AWSPointer string
	backupJob  InternJob
}

func (a *AWSBackupper) initBackup(filepath string) {
	println("per ora ci faccio niente,  ma grazie per il pensiero")
	a.backupJob = InternJob{filepath: filepath}
}

func (a *AWSBackupper) do(schedule string) {
	println("sto backuppando su aws!")
}
