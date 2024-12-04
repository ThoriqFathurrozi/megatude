package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

// Config is a struct that holds the configuration values
type Config struct {
	App struct {
		Port    int    `yaml:"port" env:"PORT" env-default:"8080"`
		Host    string `yaml:"host" env:"HOST" env-default:"localhost"`
		Version string `yaml:"version" env:"VERSION" env-default:"v1"`
	} `yaml:"app"`
	Database struct {
		Hostname string `yaml:"hostname" env:"DB_HOST" env-default:"hostname"`
		Port     int    `yaml:"port" env:"DB_PORT" env-default:"5432"`
		Name     string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
		User     string `yaml:"user" env:"DB_USER" env-default:"megatude"`
		Password string `yaml:"password" env:"DB_PASSWORD" env-default:"password"`
	} `yaml:"database"`
}

var cfg *Config

func GetConfig() *Config {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	if cfg == nil {
		sugar.Fatal("Config not loaded")
	}
	return cfg
}

func LoadConfig(configPath string) error {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err == nil {
		sugar.Fatal("Failed to load config", zap.Error(err))
	}
	return nil
}
