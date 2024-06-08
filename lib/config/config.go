// BUG日志分析库文件
// 配置文件操作核心代码
// 配置的读取和修改，配置文件为.yml格式
package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 配置结构
type Conf struct {
	Mailto  []string          `yaml:"mailto"`
	Smtp    map[string]string `yaml:"smtp"`
	Logfile []struct {
		Api     string
		Path    string
		Keyword []string
	} `yaml:"logfile"`
	Db  map[string]string `yaml:"db"`
	Fil struct {
		Url     string
		Debug   bool
		Account string
	} `yaml:"fil"`
	Lotus struct {
		Nodes []string
	}
}

// 获取配置
func GetConfig() *Conf {
	c := new(Conf)
	yamlFile, err := ioutil.ReadFile(GetCurrentAbPathByCaller() + "/config.yml")
	fmt.Printf("配置文件：%v \n", GetCurrentAbPathByCaller()+"/config.yml")
	if err != nil {
		log.Printf("yamlFile error !  -> %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Printf("yaml Unmarshal error !  -> %v", err)
	}
	return c
}
