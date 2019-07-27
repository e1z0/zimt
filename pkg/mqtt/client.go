package mqtt

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/radiohive/zimt/pkg/config"
)

// NewClient return new mqtt client
func NewClient() mqtt.Client {
	opts := newClientOptions()
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}

func newClientOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	cfg := config.NewMqttConfig()

	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", cfg.Broker, cfg.Port))
	opts.SetClientID(cfg.ClientID)

	if cfg.User != "" {
		opts.SetUsername(cfg.User)
	}
	if cfg.Password != "" {
		opts.SetPassword(cfg.Password)
	}

	return opts
}
