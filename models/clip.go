package models

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"golang.design/x/clipboard"
)

type Clip struct {
	Filedescriptor    *os.File
	bufferHash        map[string]bool
	bufferLenght      int
	cancellingContext context.Context
}

var bufferLimit int

func (c *Clip) Init(config Configuration) {
	//if file does not exist, create it
	_, err := os.Stat(config.DumpFilePath)
	//defer c.Filedescriptor.Close()
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

	//init clipboard

	cerr := clipboard.Init()
	if cerr != nil {
		println("ERROR: error init clipboard")
		panic(err)
	}

	//init map
	c.bufferHash = make(map[string]bool)

	//load bufferLimit
	bufferLimit = config.BufferLenght
	c.bufferLenght = 0
}

func (c *Clip) Watching(ctx context.Context, kill context.CancelFunc) {
	changed := clipboard.Watch(ctx, clipboard.FmtText)
	go func() {
		//open file
		for i := range changed {
			//handle file writing
			dump := base64.StdEncoding.EncodeToString(i)
			_, ok := c.bufferHash[string(i)]
			if !ok {
				if c.bufferLenght < bufferLimit {
					//logging
					fmt.Println(("value: " + string(i)))
					//add sha to map
					c.bufferHash[dump] = true
					_, err := c.Filedescriptor.WriteString(string(i) + "\n")
					if err != nil {
						panic(err)
					}
					//add content to file
				}
			}
		}
	}()
	<-ctx.Done()
	kill()
}
