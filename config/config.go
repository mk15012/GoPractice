package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Datastore struct {
		MySQL struct {
			UserName string `yaml:"user"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			DbName   string `yaml:"db_name"`
		} `yaml:"mysql"`
		// Add other datastores like Redis here in the future
	} `yaml:"datastore"`

	Port struct {
		HTTPPort int `yaml:"http.port"`
		GRPCPort int `yaml:"grpc.port"`
	} `yaml:"port"`
}

func LoadConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
