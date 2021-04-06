package jsonxml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/DustiasTheGuy/jsonxml/utils"
)

func (j *JSONXML) writeJson(v interface{}) {
	bytes, err := json.Marshal(v)

	if err != nil {
		errorHandler(err)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.%s", j.OutputPath, utils.RandSeq(20), "json"), bytes, 0644); err != nil {
		errorHandler(err)
	}
}
