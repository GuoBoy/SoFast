package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Config   *CModel
	filename string
	isCreate bool
)

func init() {
	filename = *flag.String("c", "config.yaml", "yaml格式配置文件")
	isCreate = *flag.Bool("mc", false, "生成配置文件模板")
}

// Init 初始化配置
func Init() *CModel {
	if isCreate {
		MakeDefaultConfig("example.config.yaml")
		log.Println("生成完成")
		os.Exit(0)
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Println("打开配置文件错误，将使用默认设置！")
		MakeDefaultConfig(filename)
	} else if err = yaml.Unmarshal(file, &Config); err != nil {
		log.Println("解析配置文件错误，将使用默认设置")
		MakeDefaultConfig(filename)
	}
	if _, err = os.Stat(Config.UploadPath); err != nil {
		if err = os.Mkdir(Config.UploadPath, 0644); err != nil {
			log.Fatal("创建上传文件夹失败", err)
		}
	}
	return Config
}

// MakeDefaultConfig 生成默认配置文件
func MakeDefaultConfig(filename string) {
	Config = &CModel{
		Port:       23456,
		UploadPath: "upload",
		Dev:        false,
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("生成默认配置文件失败", err)
	}
	defer file.Close()
	if err = yaml.NewEncoder(file).Encode(Config); err != nil {
		log.Fatal("保存默认配置文件失败", err)
	}
}
