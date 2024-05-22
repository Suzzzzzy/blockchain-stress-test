package config

import (
	"github.com/spf13/viper"
)

// Config 설정 정보 관리, 쉽게 로드하기 위함
type Config struct {
	DatabaseDSN string `yaml:"database_dsn"`
	RPCURL      string `yaml:"rpc_url"`
	PrivateKey  string `yaml:"private_key"`
	FromAddress string `yaml:"from_address"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
