package structs

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

const ConfigFile = "../../test/config.yaml"

func TestUnmarshalViper(t *testing.T) {
	viper.SetConfigFile(ConfigFile)
	viper.ReadInConfig()

	type mycfg struct {
		Broker    string `viper:"mqtt.broker"`
		Port      int    `viper:"mqtt.port"`
		User      string `viper:"mqtt.user"`
		AccountID string
	}

	var actual mycfg
	UnmarshalViper(&actual)

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
