package models

import (
	"encoding/xml"
	"os"
)

type Configuration struct {
	XMLName         xml.Name `xml:"Configuration"`
	ClipGui         string   `xml:"clipGui"`
	DumpFilePath    string   `xml:"DumpFilePath"`
	BufferLenght    int      `xml:"BufferLenght"`
	BackUpFrequency string   `xml:"BackUpFrequency"`
}

func (c *Configuration) LoadConfiguration() {
	//load path of configuration

	//check if file exist in root directory
	//args[1] contains the path to the config file
	_, err := os.Stat("/home/matteo/programmazione/GoClipboard/GoHistory/config.xml" /* os.Args[1] */)
	if err != nil {
		println("ERROR: File does not exist goddamit")
		return
	}
	//if there is the configfile, dump and unmarshall
	dump, err := os.ReadFile("/home/matteo/programmazione/GoClipboard/GoHistory/config.xml" /* os.Args[1] */)
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
