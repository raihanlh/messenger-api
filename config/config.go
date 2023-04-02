package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName    string `mapstructure:"APP_NAME"`
	AppVersion string `mapstructure:"APP_VERSION"`
	Port       string `mapstructure:"PORT"`
	Secret     string `mapstructure:"SECRET"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBTimezone string `mapstructure:"DB_TIMEZONE"`
	Env        string `mapstructure:"ENV"`
	Debug      bool   `mapstructure:"DEBUG"`
}

func Setup() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func New() *Config {
	Setup()
	var conf Config
	err := viper.Unmarshal(&conf)

	if err != nil {
		panic(err)
	}

	return &conf
}
