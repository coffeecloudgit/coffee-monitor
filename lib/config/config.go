// BUG日志分析库文件
// 配置文件操作核心代码
// 配置的读取和修改，配置文件为.yml格式
package config

import (
	"coffee-monitor/lib/log"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var CONF *Conf

func init() {
	CONF = GetConfig()
}

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
		Url       string `yaml:"url"`
		Debug     bool   `yaml:"debug"`
		Account   string `yaml:"account"`
		MsgServer string `yaml:"msg_server"`
		Sectors   string `yaml:"sectors"`
	} `yaml:"fil"`
	Lotus struct {
		Nodes []string
	}
	Cron struct {
		SectorsExpire string `yaml:"sectors_expire"` // 扇区过期信息检查时间
	} `yaml:"cron"`
}

// 获取配置
func GetConfig() *Conf {
	c := new(Conf)
	yamlFile, err := ioutil.ReadFile(GetCurrentAbPathByCaller() + "/config.yml")
	fmt.Printf("配置文件：%v \n", GetCurrentAbPathByCaller()+"/config.yml")
	if err != nil {
		log.Logger.Info("yamlFile error !  -> %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Logger.Info("yaml Unmarshal error !  -> %v", err)
	}
	return c
}
