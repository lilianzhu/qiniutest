package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type EnvParam struct {
	Environment struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func (e *EnvParam) GetHost() {
	filename, _ := filepath.Abs("../../../qiniutest.com/configs/env.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &e)
	if err != nil {
		panic(err)
	}
}
