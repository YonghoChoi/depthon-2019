package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type Config struct {
	Mongo string `yaml:"mongo"`
}

var gConfig Config
var once sync.Once
var gFilePath string

func SetConfigFilePath(path string) {
	gFilePath = path
}

func GetConfig() Config {
	once.Do(func() {
		if gFilePath == "" {
			gFilePath = "/Users/yongho/Works/Projects/depthon-2019/cmd/pengha-api/conf/config.yml"
		}

		fmt.Println("filename : " + gFilePath)
		yamlFile, err := ioutil.ReadFile(gFilePath)
		err = yaml.Unmarshal(yamlFile, &gConfig)
		if err != nil {
			panic(err)
		}

		fmt.Printf("load config. %v\n", gConfig)
	})
	return gConfig
}
