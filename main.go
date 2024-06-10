package main

import (
	"context"
	. "main/models"
	"os"
	"os/signal"
)

var clip Clip
var backup Backup

func init() {
	// check and load configuration
	var config = new(Configuration)

	//actually load configuration
	config.LoadConfiguration()

	//ini the files and clipboard object
	clip.Init(*config)

	clip.InitializeBashScript(*config)
	backup.SetupBackup(config)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	backup.ExecuteBackup()
	clip.Watching(ctx, cancel)
}
