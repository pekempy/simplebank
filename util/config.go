package util

import "github.com/spf13/viper"

// Config stores all configuration of the application
// Values read by Viper from a config file/environmental variables
type Config struct {
	DBDriver 		string  	`mapstructure:"DB_DRIVER"`
	DBSource 		string  	`mapstructure:"DB_SOURCE"`
	ServerAddress 	string  	`mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig reads config from file or environmental variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml
	
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}