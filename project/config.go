package project

import (
	"encoding/json"
	"io/ioutil"
)

const (
	CONFIG = "cmplr.conf"
)

type ConfigFile struct {
	ProjectName string
	Jobs        []*Job
}

// Reads 'cmplr.conf' file with the following format:
// Example:
//   {
//     "Jobs": [
//       {
//         "Outputfile": "dist/css/main.css",
//         "InputFiles": [
//           "src/css/layout.less",
//           "src/css/menu.less"
//         ]
//       }
//     ]
//   }
func ReadConfig() (cfg *ConfigFile, err error) {
	file, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &cfg)
	return
}
