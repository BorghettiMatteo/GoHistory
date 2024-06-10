package models

import (
	"context"
	"encoding/base64"
	"os"
	"regexp"

	"golang.design/x/clipboard"
)

type Clip struct {
	Filedescriptor *os.File
	bufferHash     map[string]bool
	bufferLenght   int
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
		defer c.Filedescriptor.Close()
		for i := range changed {
			println("STRINGVALUE: " + string(i))
			//handle file writing
			dump := base64.StdEncoding.EncodeToString(i)
			_, ok := c.bufferHash[string(i)]
			if !ok {
				if c.bufferLenght > bufferLimit {
					//clear history
					c.bufferHash = make(map[string]bool)
				}
				//add sha to map
				c.bufferHash[dump] = true
				//write to file
				c.Filedescriptor.WriteString("#\n")
				_, err := c.Filedescriptor.WriteString(dump + "\n")
				if err != nil {
					panic(err)
				}
				c.bufferLenght++
			}
		}
	}()
	<-ctx.Done()
	println("GRACEFULLY SHUTDOWN")
	kill()
}

func (c *Clip) InitializeBashScript(config Configuration) {
	// write dumpfileconfig to disc
	re := regexp.MustCompile(`filepath=.*\n`)
	_, err := os.Stat(config.ClipGui)
	if err == nil {
		// read file
		filed, err := os.ReadFile(config.ClipGui)
		if err != nil {
			panic(err)
		}
		data := re.FindAll(filed, -1)
		println(data)
		newstring := re.ReplaceAllString(string(filed), "filepath=\""+config.DumpFilePath+"\"\n")
		err = os.WriteFile(config.ClipGui, ([]byte(newstring)), 0666)
		if err != nil {
			panic(err)
		}
	}
}
