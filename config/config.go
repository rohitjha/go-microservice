package config

import "github.com/spf13/viper"

type Config struct {
	Port             string `mapstructure:"PORT"`
	CertFileLocation string `mapstructure:"CRT_FILE"`
	KeyFileLocation  string `mapstructure:"KEY_FILE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
