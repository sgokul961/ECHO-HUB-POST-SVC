package config

import "github.com/spf13/viper"

type Config struct {
	Port               string `mapstructure:"PORT"`
	DBUrl              string `mapstructure:"DB_URL"`
	AuthHubUrl         string `mapstructure:"auth_hub_url"`
	NotificationHubUrl string `mapstructure:"notification_hub_url"`

	//Kafka KafkaConfig
}

// type KafkaConfig struct {
// 	BrokerURL string `mapstructure:"KAFKA_BROKER_URL"`
// 	// Add more Kafka configuration fields as needed
// }

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./pkg/config/envs")

	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Add Kafka-specific environment variable prefix
	viper.SetEnvPrefix("KAFKA")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	// // Unmarshal Kafka configuration into KafkaConfig struct
	// var kafkaConfig KafkaConfig
	// err = viper.Unmarshal(&kafkaConfig)
	// if err != nil {
	// 	return
	// }
	// config.Kafka = kafkaConfig

	err = viper.Unmarshal(&config)
	return

}
