package config

import "github.com/spf13/viper"

//CONFIG for configs
type CONFIG struct {
	DBType     string `mapstructure:"DB_TYPE"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassWord string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`

	PORT string `mapstructure:"PORT"`

	DBConnectionString string `mapstructure:"DB_CONNECTION_STRING"`
}

//LoadConfig reads configurations from app.env file
func LoadConfig(path string) (CONFIG, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	var config CONFIG
	//	viper.AutomaticEnv() //reads environmental variables and overrides those in .env
	err := viper.ReadInConfig()
	if err != nil {
		return CONFIG{}, err
	}

	err = viper.Unmarshal(&config)
	return config, nil
}
