package inttests

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type TestConfig struct {
	MainDbConnection    string `mapstructure:"MAIN_DB_CONNECTION"`
	RmqConnectionString string `mapstructure:"RMQ_CONNECTION_STRING"`
}

func NewTestConfig() (*TestConfig, error) {
	// # Read os env
	viper.AutomaticEnv()

	pwd := os.Getenv("TEST_CONFIG_PWD")

	// # Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath(pwd)

	// # Tell viper the name of your file
	viper.SetConfigName("test")

	// # Tell viper the type of your file
	viper.SetConfigType("env")

	// # Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower("Not Found in")) {
			return nil, err
		}
	}

	config := &TestConfig{}

	// # Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
