package config

import (
	"github.com/spf13/viper"
)

// MqttConfig represents mqtt connection options
type MqttConfig struct {
	Broker    string
	Port      int
	BaseTopic string
	User      string
	Password  string
}

// NewMqttConfig returns new value of mqtt config
func NewMqttConfig() MqttConfig {
	return MqttConfig{
		Broker:    viper.GetString("mqtt.broker"),
		Port:      viper.GetInt("mqtt.port"),
		BaseTopic: viper.GetString("mqtt.base-topic"),
		User:      viper.GetString("mqtt.user"),
		Password:  viper.GetString("mqtt.password"),
	}
}

func init() {
	viper.SetDefault("mqtt.broker", "localhost")
	viper.SetDefault("mqtt.port", 1883)
	viper.SetDefault("mqtt.base-topic", "zigbee2mqtt")
}
