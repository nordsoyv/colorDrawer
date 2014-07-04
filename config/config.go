package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ColorCubeSide int
	OutputFilename string
}

func Read(path string) (Config) {
	var config Config
	configFile , err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	json.Unmarshal( configFile , &config)
	return config;

}
