package core

import (
	"fim_server/core/config"
	"fim_server/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
)

// Yaml 读取yaml文件的配置
func Yaml() *config.Config {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println("配置文件加载失败", err.Error())
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		fmt.Println("config Init Unmarshal: ", err.Error())
	}
	return c
}

// SetYaml 修改yaml文件
func SetYaml() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		fmt.Println("配置修改失败: " + err.Error())
	}
	err = ioutil.WriteFile(configFile, byteData, fs.ModePerm)
	if err != nil {
		fmt.Println("配置修改失败: " + err.Error())
	}
}
