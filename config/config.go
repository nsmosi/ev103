package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Version string `yaml:"version"`
	} `yaml:"app"`
	Rdbs struct {
		Address string `yaml:"addr"`
	} `yaml:"rdbs"`
	CrdbAddress string `yaml:"crdb_addr"`
	ApiServer   string `yaml:"api_server"`
	EnableLoad  bool   `yaml:"enable_load"`
	DataFile    string `yaml:"data_file"`
	BundleFile  string `yaml:"bundle_file"`
}

func LoadConfigData(configFilePath string) (*Config, error) {

	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML file: %v", err)
	}

	return &config, nil
}
