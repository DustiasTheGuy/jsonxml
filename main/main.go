package main

import jsonxml "github.com/DustiasTheGuy/jsonxml/jtx"

type Output struct {
	Post []string `xml:"post"`
	Date []string `xml:"date"`
}

func main() {
	config := jsonxml.JSONXML{
		DataPath:   "./blogs",
		OutputPath: "./output",
		MaxFiles:   10,
	}

	config.Init()
}
