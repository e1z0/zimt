package config

import (
	"os"

	"github.com/spf13/viper"

	"github.com/radiohive/zimt/pkg/structs"
)

// MqttConfig represents mqtt connection options
type MqttConfig struct {
	Broker    string `viper:"mqtt.broker"`
	Port      int    `viper:"mqtt.port"`
	BaseTopic string `viper:"mqtt.base-topic"`
	User      string `viper:"mqtt.user"`
	Password  string `viper:"mqtt.password" print:"mask"`
	ClientID  string `viper:"mqtt.client-id"`
}

// Print prints mqtt config to standard output
func (mc MqttConfig) Print() {
	structs.Print(&mc, "viper", os.Stdout)
}

// NewMqttConfig returns new value of mqtt config
func NewMqttConfig() MqttConfig {
	var mqtt MqttConfig
	structs.UnmarshalViper(&mqtt)
	return mqtt
}

func init() {
	viper.SetDefault("mqtt.broker", "localhost")
	viper.SetDefault("mqtt.port", 1883)
	viper.SetDefault("mqtt.base-topic", "zigbee2mqtt")
	viper.SetDefault("mqtt.client-id", "zimt")
}
