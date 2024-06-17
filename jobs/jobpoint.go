package main

import (
	"flag"
	. "main/models/jobModels"
	"os"
)

func main() {
	var jobFlag = flag.String("j", "DefaultJob", "This is the job type, to show current list please use ... -h")
	helpFlag := flag.Bool("h", false, "print helper message")

	flag.Parse()

	if *helpFlag {
		println("this is the helper flag")
		os.Exit(0)
	}

	// handle jobs type passed by the flag
	println(*jobFlag)
	var job GenericJobInterface
	switch *jobFlag {
	case "RestoreFromFilesystem":
		job = new(RestoreBackupFileSystem)
	default:
		println("orococoasod")
		return
	}
	// execute job
	job.RunJob()

}
