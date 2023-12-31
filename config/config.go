package config

import "github.com/spf13/viper"

type Config struct {
	Driver string `mapstructure:"DRIVER"`
	Source string `mapstructure:"DATABASE_URL"`
	Port   string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
