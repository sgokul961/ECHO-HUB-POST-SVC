package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBUrl      string `mapstructure:"DB_URL"`
	AuthHubUrl string `mapstructure:"auth_hub_url"`
}

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./pkg/config/envs")

	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}