package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type DataSource struct {
	Host        string `yaml:"host"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Port        string `yaml:"port"`
	Name        string `yaml:"name"`
	TablePrefix string
}

var DataSourceConfig = DataSource{}

type Yaml struct {
	Redis Redis `yaml:"redis"`
}
type Redis struct {
	Host        string        `yaml:"host"`
	Password    string        `yaml:"password"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxActive   int           `yaml:"maxActive"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
}

var yamlConfig = Yaml{}
var RedisConfig = Redis{}

func Init() {
	yamlFile, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/config.yaml': %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	RedisConfig = yamlConfig.Redis
}
func InitRedisTest() {
	var data = `redis:
  host: 192.168.137.133:6379
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
