package utils

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// Config 应用配置
type Config struct {
	Env  string     `yaml:"env"`  // 开发模式
	HTTP HTTPConfig `yaml:"http"` // HTTP 配置
	DB   DBConfig   `yaml:"db"`   // 数据库配置
	APIs APIs       `yaml:"apis"` //接口地址配置
}

// HTTPConfig 服务器配置
type HTTPConfig struct {
	Port    string        `yaml:"port"` // 监听端口
	JWT     JWTConfig     `yaml:"jwt"`
	Session SessionConfig `yaml:"session"` // Session 配置
}

// JWTConfig JWT 配置
type JWTConfig struct {
	Key     string `yaml:"key"`     // signer 名字
	Expires string `yaml:"expires"` // 过期天数
}

// APIs 配置
type APIs struct {
	UsersAPI string `yaml:"usersAPI"` // UsersAPI
	ToolsAPI string `yaml:"toolsAPI"` // toolsAPI
}

// SessionConfig Session 配置
type SessionConfig struct {
	Key     string `yaml:"key"`     // Cookies 名字
	Expires int    `yaml:"expires"` // 过期天数
}

// DBConfig 数据库配置
type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var config *Config

// LoadConf 从文件读取配置信息
func (c *Config) LoadConf(path string) *Config {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic().Err(err).Msg("Can't read config file")
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Panic().Err(err).Msg("Can't marshal config file")
	}
	log.Info().Msg("Read config from " + path)
	config = c
	return c
}

// GetConf 获取全局配置
func GetConf() *Config {
	return config
}
