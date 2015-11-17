package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var map_conf = make(map[string]string)

func selectConfigFile() string {
	env := os.Getenv("TEST_ENV")
	if env == "" {
		panic("Please set environment: TEST_ENV")
	}

	zone := os.Getenv("TEST_ZONE")
	if zone == "" {
		panic("Please set environment: TEST_ZONE")
	}
	return env + "_" + zone + ".yaml"
}

func loadConfig(m *map[string]string) {
	filePath := "../../configs/" + selectConfigFile()
	fmt.Printf("CHOOSE CONFIG FILE: %v \n", filePath)
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("Read Config file failed: " + err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		panic("Failed to unmarshall file: " + filePath)
	}
}

/*
 获取配置文件里的属性值
*/
func ENV(key string) string {
	if map_conf[key] == "" {
		loadConfig(&map_conf)
	}

	fmt.Printf("ENVIRONMENTS: %v \n", map_conf)
	return map_conf[key]
}
