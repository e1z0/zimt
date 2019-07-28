package structs

import (
	"reflect"
	"strings"
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

func TestPrint(t *testing.T) {
	viper.SetConfigFile(ConfigFile)
	viper.ReadInConfig()

	type mycfg struct {
		Broker    string `viper:"mqtt.broker"`
		Port      int    `viper:"mqtt.port"`
		User      string `viper:"mqtt.user"`
		Password  string `viper:"mqtt.password" print:"mask"`
		ClientID  string `viper:"mqtt.client-id" print:"-"`
		AccountID string
	}

	cfg := mycfg{
		Broker:    "1.2.3.4",
		Port:      1234,
		User:      "my-user",
		Password:  "my-password",
		ClientID:  "my-client-id",
		AccountID: "my-account-id",
	}

	actual := strings.Builder{}
	Print(&cfg, "viper", &actual)

	expected := `mqtt.broker="1.2.3.4"
mqtt.port=1234
mqtt.user="my-user"
mqtt.password="my**********rd"
AccountID="my-account-id"
`

	if !reflect.DeepEqual(expected, actual.String()) {
		t.Errorf("Expected %+v, got %+v", expected, actual.String())
	}
}
