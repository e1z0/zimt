package api

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// BridgeConfig represents message of `#/bridge/config` topic
type BridgeConfig struct {
	Version     string `json:"version"`
	Commit      string `json:"commit"`
	Coordinator int    `json:"coordinator"`
	LogLevel    string `json:"log_level"`
	PermitJoin  bool   `json:"permit_join"`
}

// GetBridgeConfig subscribes and returns message  from `#/bridge/config` topic
func GetBridgeConfig(client mqtt.Client) BridgeConfig {
	topic := topic("bridge/config")
	msg := getSubscribedOnce(client, topic)
	config := BridgeConfig{}
	if err := json.Unmarshal(msg.Payload(), &config); err != nil {
		panic(err)
	}
	return config
}
