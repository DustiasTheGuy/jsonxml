package jsonxml

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/DustiasTheGuy/jsonxml/utils"
)

type JSONXML struct {
	DataPath   string
	OutputPath string
	MaxFiles   int
}

type Output struct {
	Post []string `xml:"post"`
}

func Combine() {

}

func (j *JSONXML) Init() {
	files := j.readDir()
	var data []string
	var str string

	consoleOutput(fmt.Sprintf("Found %d files", len(files)))

	for i := 0; i < 100; i++ {
		data = append(data, j.parseFile(files[i]).Post...)
	}

	for i := 0; i < len(data); i++ {
		str += strings.TrimSpace(data[i])
	}

	j.writeTxt(str)
	size := len(str)
	consoleOutput(fmt.Sprintf("Wrote %d bytes to disk", size))
	consoleOutput("Done")
}

func (j *JSONXML) Output(output Output) {
	var s string

	for _, v := range output.Post {
		s += strings.TrimSpace(v)
	}

	j.writeTxt(s)
}

func (j *JSONXML) writeTxt(v string) {
	if len(v) > 0 {
		err := ioutil.WriteFile(fmt.Sprintf("%s/%s.%s", j.OutputPath, utils.RandSeq(20), "txt"), []byte(v), 0644)
		errorHandler(err)
	}
}

func consoleOutput(str string) {
	fmt.Printf("---> %s \n", str)
}

func errorHandler(err error) {
	if err != nil {
		//log.Fatal(err)
		//fmt.Println("An error")
	}

	//fmt.Println(err)
}
