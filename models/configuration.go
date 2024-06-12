package models

import (
	"encoding/xml"
	"os"
)

type Configuration struct {
	XMLName      xml.Name `xml:"Configuration"`
	DaemonPath   string   `xml:"DaemonPath"`
	DumpFilePath string   `xml:"DumpFilePath"`
}

func (c *Configuration) LoadConfiguration() {
	//load path of configuration

	//check if file exist in root directory
	_, err := os.Stat("/home/matteo/programmazione/GoClipboard/GoHistory/config.xml")
	if err != nil {
		println("ERROR: File does not exist goddamit ")
		return
	}
	//if there is the configfile, dump and unmarshal
	dump, err := os.ReadFile("config.xml")
	if err != nil {
		println("ERROR: Cannot load configuration")
		panic(err)
	}

	//unmarshal struct
	err = xml.Unmarshal(dump, c)
	if err != nil {
		panic(err)
	}
}
