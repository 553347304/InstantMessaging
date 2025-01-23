package config

import (
	"fim_server/utils/stores/logs"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
)

// Yaml 读取yaml文件的配置
func Yaml(configFile string) *Config {
	c := &Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		logs.Fatal("配置文件加载失败", err.Error())
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		logs.Fatal("配置文件解析失败", err.Error())
	}
	return c
}
func SetYaml(configFile string, data Config) {
	byteData, err := yaml.Marshal(data)
	if err != nil {
		logs.Fatal("配置修改失败", err.Error())
	}
	err = ioutil.WriteFile(configFile, byteData, fs.ModePerm)
	if err != nil {
		logs.Fatal("配置修改失败", err.Error())
	}
}
