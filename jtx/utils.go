package jsonxml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func (j *JSONXML) readDir() []string {
	var filesNames []string

	files, err := ioutil.ReadDir(j.DataPath)
	errorHandler(err)

	for _, val := range files {
		filesNames = append(filesNames, fmt.Sprintf("%s/%s", j.DataPath, val.Name()))
	}

	return filesNames
}

// 5114.male.25.indUnk.Scorpio.xml
func (j *JSONXML) parseFile(filename string) Output {
	var output Output

	bytes, err := ioutil.ReadFile(filename)
	errorHandler(err)

	err = xml.Unmarshal(bytes, &output)
	errorHandler(err)

	return output
}
