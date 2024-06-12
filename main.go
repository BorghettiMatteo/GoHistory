package main

import (
	"context"
	"fmt"
	. "main/models"
	"time"

	"golang.design/x/clipboard"
)

func init() {
	// check and load configuration
	var config = new(Configuration)

	//actually load configuration
	config.LoadConfiguration()
}

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	//var forever chan struct{}

	fmt.Println(string(clipboard.Read(clipboard.FmtText)))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	//defer cancel()
	changed := clipboard.Watch(ctx, clipboard.FmtText)
	go func() {
		for i := range changed {
			println(string(i))
		}
	}()
	<-ctx.Done()

}
