package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"strconv"
)

type AppConfig struct {
	Server  *ServerConfig
	Logging *LogConfig
}

type LogConfig struct {
	Filename         string
	Size             int64
	MaxNumberOfFiles int
}

type ServerConfig struct {
	Port int64
}

func (a *AppConfig) ReadConfig(configFile string) {
	fileContent, e := ioutil.ReadFile(configFile)

	if e != nil {
		fmt.Println("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(fileContent, &a)

}