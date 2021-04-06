package jsonxml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/DustiasTheGuy/jsonxml/utils"
)

type JSONXML struct {
	DataPath string
}

type Output struct {
	Post []string `xml:"post"`
	Date []string `xml:"date"`
}

func (j *JSONXML) Init() {
	files, err := j.readDir()

	if err != nil {
		errorHandler(err)
	}

	consoleOutput(fmt.Sprintf("Found %d files", len(files)))
	for i := 0; i < 10; i++ {
		j.parseFile(files[i])
	}
}

func (j *JSONXML) readDir() ([]string, error) {
	var filesNames []string

	files, err := ioutil.ReadDir(j.DataPath)

	if err != nil {
		return nil, err
	}

	for _, val := range files {
		filesNames = append(filesNames, fmt.Sprintf("%s/%s", j.DataPath, val.Name()))
	}

	return filesNames, nil
}

// 5114.male.25.indUnk.Scorpio.xml
func (j *JSONXML) parseFile(filename string) {
	var formatted []map[string]interface{}
	var dest Output

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		errorHandler(err)
	}

	if err := xml.Unmarshal(bytes, &dest); err != nil {
		//fmt.Printf("error in: %s\n", filename)
		errorHandler(err)
	}

	for i := 0; i < len(dest.Post); i++ {
		formatted = append(formatted, map[string]interface{}{
			"post": strings.TrimSpace(dest.Post[i]),
			"date": dest.Date[i],
		})
	}

	writeResult(formatted)
}

func writeResult(v interface{}) {
	bytes, err := json.Marshal(v)
	outputDir := "./output"

	if err != nil {
		errorHandler(err)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.%s", outputDir, utils.RandSeq(20), "json"), bytes, 0644); err != nil {
		errorHandler(err)
	}
}

func consoleOutput(str string) {
	fmt.Printf("----- %s -----\n", str)
}

func errorHandler(err error) {
	//fmt.Println(err)
}
