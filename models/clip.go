package models

import (
	"os"
)

type clip struct {
	FilePath os.File
}

func (c clip) Init(config Configuration) {
	//if file does not exist, create it
	_, err := os.Stat(config.DumpFilePath)
	if err != nil {
		os.WriteFile(config.DumpFilePath, []byte{}, 0666)
	}
}
