package env

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type config struct {
	data map[string]string
}

var env config

func Init(path string) {
	var fileData []byte = make([]byte, 1024)
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
	}

	out := make(map[string]interface{})
	if err := json.Unmarshal(fileData, &out); err != nil {
	}

	configs := make(map[string]string)
	for k, v := range out {
		configs[k] = fmt.Sprintf("%v", v)
	}

	env = config{data: configs}
}

func Get(name string) string {
	if value := os.Getenv(strings.ToUpper(name)); value != "" {
		return value
	}

	if value := env.data[name]; value != "" {
		return value
	}
	return ""
}
