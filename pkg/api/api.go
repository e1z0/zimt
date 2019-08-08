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

func getSubscribedOnPublishOnce(client mqtt.Client, subTopic string, pubTopic string, pubPayload string) mqtt.Message {
	resp := make(chan mqtt.Message, 1)
	select {
	case m := <-subscribeAndPublish(client, subTopic, pubTopic, pubPayload):
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

func publish(c mqtt.Client, pt string, pp string) {
	token := c.Publish(pt, 0, false, pp)
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		panic(err)
	}
}

func subscribeAndPublish(c mqtt.Client, st string, pt string, pp string) chan mqtt.Message {
	ch := make(chan mqtt.Message, 1)
	c.Subscribe(st, 0, func(client mqtt.Client, msg mqtt.Message) {
		ch <- msg
	})
	publish(c, pt, pp)
	return ch
}

func topic(path string) string {
	c := config.NewMqttConfig()
	return fmt.Sprintf("%s/%s", c.BaseTopic, path)
}

func systopic(path string) string {
	return fmt.Sprintf("$SYS/broker/%s", path)
}
