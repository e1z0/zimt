package config

import (
	"reflect"
	"testing"
)

func TestUnmarshalWithDefaultConfig(t *testing.T) {
	Load("../../test/config.yaml")

	var actual MqttConfig
	unmarshal(&actual)

	expected := MqttConfig{
		Broker:    "1.2.3.4",
		Port:      1234,
		BaseTopic: "my-topic",
		User:      "my-user",
		Password:  "my-password",
		ClientID:  "my-client",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestUnmarshalWithNotTaggedFields(t *testing.T) {
	Load("../../test/config.yaml")

	type mycfg struct {
		Broker    string `viper:"mqtt.broker"`
		Port      int    `viper:"mqtt.port"`
		User      string `viper:"mqtt.user"`
		AccountID string
	}

	var actual mycfg
	unmarshal(&actual)

	expected := mycfg{
		Broker:    "1.2.3.4",
		Port:      1234,
		User:      "my-user",
		AccountID: "",
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
