package core

import (
	"fmt"
	"io/fs"
	"log"             
	"os"               
	"github.com/backunderstar/zew/config" 
	"github.com/backunderstar/zew/global"

	"gopkg.in/yaml.v2"
)

// 配置文件路径常量
const ConfigFile = "settings.yaml"

// initConfig 函数用于初始化应用程序的配置信息
func InitConfig() {
	// 创建一个指向config.Config结构体的指针
	c := &config.Config{}

	// 尝试读取配置文件内容
	yamlConfig, err := os.ReadFile(ConfigFile)
	if err != nil {
		// 如果读取文件出错，则输出错误信息并终止程序执行
		panic(fmt.Errorf("read config file error: %v", err))
	}
	// 使用yaml.Unmarshal函数将读取到的YAML数据解码到c指向的结构体中
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		// 如果解码过程出错，则记录错误日志并终止程序
		log.Fatalf("unmarshal config file error: %v", err)
	}

	// 如果配置初始化成功，打印日志信息
	log.Println("config init success")

	// 将配置信息赋值给全局变量global.Config
	global.Config = c
}

func UpdateYaml() error {
	data, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, data, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("系统配置更新成功")
	return nil
}
