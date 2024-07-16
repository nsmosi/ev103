package config

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

	return nil, nil
}
