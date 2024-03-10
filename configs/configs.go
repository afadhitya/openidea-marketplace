package configs

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	AppName    string `mapstructure:"app_name"`
	AppVersion string `mapstructure:"app_version"`
	DBHost     string `mapstructure:"db_host"`
	DBName     string `mapstructure:"db_name"`
	DBPassword string `mapstructure:"db_password"`
	DBPort     string `mapstructure:"db_port"`
	DBUsername string `mapstructure:"db_username"`
	ENV        string `mapstructure:"env"`
	Port       string `mapstructure:"port"`
}

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	if cfg == nil {
		cfg = &Config{}
	}

	return cfg
}

func Load() {
	once.Do(func() {
		v := viper.New()
		v.AutomaticEnv()

		v.AddConfigPath(".")
		v.SetConfigType("env")
		v.SetConfigName(".env")

		err := v.ReadInConfig()
		if err != nil {
			fmt.Println("config file not found: ", err)
		}

		config := new(Config)
		err = v.Unmarshal(config)
		if err != nil {
			panic(err)
		}

		cfg = config
	})
}
