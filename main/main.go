package main

import jsonxml "github.com/DustiasTheGuy/jsonxml/jtx"

func main() {
	config := jsonxml.JSONXML{
		DataPath: "./blogs",
	}

	config.Init()
}
