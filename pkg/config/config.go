package config

import (
	"fmt"
	"ginLearn/pkg/file"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type App struct {
	JwtSecret string `yaml:"jwtSecret"`
	LogPath   string `yaml:"logPath"`
}
type Server struct {
	RunMode  string `yaml:"runMode"`
	HttpPort int    `yaml:"httpPort"`
}
type DataSource struct {
	Host        string `yaml:"host"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Port        string `yaml:"port"`
	Name        string `yaml:"name"`
	TablePrefix string
}
type Redis struct {
	Host        string        `yaml:"host"`
	Password    string        `yaml:"password"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxActive   int           `yaml:"maxActive"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

var yamlConfig = Yaml{}
var AppConfig App
var ServerConfig Server
var RedisConfig Redis
var DefaultConfigFile = "configs/config.yaml"
var DataSourceConfig DataSource

type Yaml struct {
	App    App    `yaml:"app"`
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
}

func Setup(configPath string) {
	path := DefaultConfigFile
	if len(configPath) > 0 {
		path = configPath
	}
	configExist, _ := file.PathExists(path)
	if true != configExist {
		log.Println("配置文件 " + path + "没有找到!")
		os.Exit(-1)
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse "+DefaultConfigFile+": %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	AppConfig = yamlConfig.App
	ServerConfig = yamlConfig.Server
	RedisConfig = yamlConfig.Redis
}
func InitRedisTest() {
	var data = `redis:
  host: 192.168.137.128:6379
  password:
  maxIdle: 30
  maxActive: 30
  idleTimeout: 200
`
	err := yaml.Unmarshal([]byte(data), &yamlConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	RedisConfig = yamlConfig.Redis
}
