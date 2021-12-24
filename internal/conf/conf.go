package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Http     HttpConfig     `yaml:"http"`
	Grpc     GrpcConfig     `yaml:"grpc"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}
type HttpConfig struct {
	Addr    string `yaml:"addr"`
	Timeout string `yaml:"timeout"`
}
type GrpcConfig struct {
	Addr    string `yaml:"addr"`
	Timeout string `yaml:"timeout"`
}
type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}
type RedisConfig struct {
	Addr         string `yaml:"addr"`
	ReadTimeout  string `yaml:"read_timeout"`
	WriteTimeout string `yaml:"write_timeout"`
}

func GetConf(path string) *Config {
	conf := &Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
