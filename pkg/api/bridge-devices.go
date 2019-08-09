package api

import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// BridgeConfigDevice represents message of `#/bridge/config/devices`
type BridgeConfigDevice struct {
	IEEEAddr     string `json:"ieeeAddr"`
	Type         string `json:"type"`
	Model        string `json:"model"`
	FriendlyName string `json:"friendly_name"`
	NWKAddr      int    `json:"nwkAddr"`
	ManufID      int    `json:"manufId"`
	ManufName    string `json:"manufName"`
	PowerSource  string `json:"powerSource"`
	ModelID      string `json:"modelId"`
	HWVersion    int    `json:"hwVersion"`
	SWBuildID    string `json:"swBuildId"`
	DateCode     string `json:"dateCode"`
}

// GetBridgeConfigDevices subscribes and returns message  from `#/bridge/config/devices` topic
func GetBridgeConfigDevices(client mqtt.Client) []BridgeConfigDevice {
	subTopic := topic("bridge/config/devices")
	pubTopic := topic("bridge/config/devices/get")
	msg := getSubscribedOnPublishOnce(client, subTopic, pubTopic, "")
	devices := []BridgeConfigDevice{}
	if err := json.Unmarshal(msg.Payload(), &devices); err != nil {
		panic(err)
	}
	return devices
}
