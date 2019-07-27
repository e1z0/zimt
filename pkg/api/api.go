package api

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/radiohive/zimt/pkg/config"
)

func getSubscribedOnce(c mqtt.Client, t string) mqtt.Message {
	resp := make(chan mqtt.Message, 1)
	select {
	case m := <-subscribe(c, t):
		resp <- m
	case <-time.After(5 * time.Second):
		panic("timeout: can't get message over 5 seconds")
	}
	return <-resp
}

func subscribe(c mqtt.Client, t string) chan mqtt.Message {
	ch := make(chan mqtt.Message, 1)
	c.Subscribe(t, 0, func(client mqtt.Client, msg mqtt.Message) {
		ch <- msg
	})
	return ch
}

func topic(path string) string {
	c := config.NewMqttConfig()
	return fmt.Sprintf("%s/%s", c.BaseTopic, path)
}
