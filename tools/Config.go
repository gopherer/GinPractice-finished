package tools

import (
	"bufio"
	"encoding/json"
	"github.com/wonderivan/logger"
	"os"
)

type Config struct {
	WebName  string         `json:"web_name"`
	WebMode  string         `json:"web_mode"`
	WebHost  string         `json:"web_host"`
	WebPort  string         `json:"web_port"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}
func ParseConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.Alert("输入路径有误，json文件读取失败")
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Alert("文件关闭失败")
			os.Exit(1)
		}
	}()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader) //NewDecoder创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	err = decoder.Decode(&cfg)         //Decode从输入流读取下一个json编码值并保存在v指向的值里，参见Unmarshal函数的文档获取细节信息。
	if err != nil {
		return nil, err
	} else {
		return cfg, nil
	}

}
