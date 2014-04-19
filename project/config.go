package project

import (
	"io/ioutil"
	"encoding/json"
)

const (
	CONFIG = "cmplr.conf"
)


type ConfigFile struct {
	ProjectName string
	Jobs []*Job
}

func ReadConfig() (cfg *ConfigFile, err error) {
	file, err := ioutil.ReadFile(CONFIG)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &cfg)
	return
}
