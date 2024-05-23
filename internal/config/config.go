package config

import (
	"blockchain/internal/util"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

// Config 설정 정보 관리, 쉽게 로드하기 위함
type Config struct {
	DatabaseDSN string `yaml:"database_dsn"`
	PrivateKey  string `yaml:"private_key"`
	FromAddress string `yaml:"from_address"`
}

func LoadConfig() (*Config, error) {
	logger := util.LogManager{}
	cfg := Config{}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(filepath.Join(basepath, "../..")) // 루트 디렉토리로 경로 설정

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("failed read config")
		return nil, err
	}
	data := viper.AllSettings()
	if value, ok := data["database_dsn"].(string); ok {
		cfg.DatabaseDSN = value
	}
	if value, ok := data["private_key"].(string); ok {
		cfg.PrivateKey = value
	}
	if value, ok := data["from_address"].(string); ok {
		cfg.FromAddress = value
	}

	return &cfg, nil
}
