package models

import (
	"os"
)

type Clip struct {
	Filedescriptor *os.File
	ClipBuffer     []string
}

func (c *Clip) Init(config Configuration) {
	//if file does not exist, create it
	_, err := os.Stat(config.DumpFilePath)
	defer c.Filedescriptor.Close()
	if err != nil {
		c.Filedescriptor, err = os.Create(config.DumpFilePath)
		if err != nil {
			panic(err)
		}
	} else {
		//open file
		c.Filedescriptor, err = os.OpenFile(config.DumpFilePath, os.O_RDWR, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		//if file exists, clear content of file
		err := c.Filedescriptor.Truncate(0)
		if err != nil {
			panic(err)
		}
	}
}
