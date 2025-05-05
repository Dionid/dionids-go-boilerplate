package main

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env string `mapstructure:"ENV"`

	Host  string `mapstructure:"HOST"`
	Port  int    `mapstructure:"PORT"`
	Https bool   `mapstructure:"HTTPS"`

	MainDbConnectionStr string `mapstructure:"MAIN_DB_CONNECTION_STR"`

	FirstUserEmail    string `mapstructure:"FIRST_USER_EMAIL"`
	FirstUserPassword string `mapstructure:"FIRST_USER_PASSWORD"`

	SwaggerPathPrefix string `mapstructure:"SWAGGER_PATH_PREFIX"`

	JwtSecret          string `mapstructure:"JWT_SECRET"`
	JwtExpireInSeconds int64  `mapstructure:"JWT_EXPIRE_IN_SECONDS"`
}

// Call to load the variables from env
func initConfig() (*Config, error) {
	// # Read os env
	viper.AutomaticEnv()

	// # Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(".")

	viper.SetDefault("PORT", 8080)

	// # Tell viper the name of your file
	viper.SetConfigName("app")

	// # Tell viper the type of your file
	viper.SetConfigType("env")

	// # Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower("Not Found in")) {
			return nil, err
		}
	}

	config := &Config{}

	// # Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
