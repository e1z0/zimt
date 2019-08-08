package api

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

const (
	oneMinInSeconds  = 60
	oneHourInSeconds = oneMinInSeconds * 60
	oneDayInSeconds  = oneHourInSeconds * 24
)

// Broker represents broker details from multiple `$SYS/broker/*` topics
type Broker struct {
	Version          string
	Uptime           string
	MessagesSent     string
	MessagesReceived string
	ClientsConnected string
	ClientsTotal     string
}

// FormatUptime formats uptime string
func (b Broker) FormatUptime() string {
	var seconds int
	_, err := fmt.Sscanf(b.Uptime, "%d seconds", &seconds)
	if err != nil {
		if viper.GetBool("verbose") {
			fmt.Println("err while parsing uptime: ", err.Error())
		}
		return b.Uptime
	}

	days := seconds / oneDayInSeconds
	if days > 0 {
		seconds %= days * oneDayInSeconds
	}

	hours := seconds / oneHourInSeconds
	if hours > 0 {
		seconds %= hours * oneHourInSeconds
	}

	minutes := seconds / oneMinInSeconds
	if minutes > 0 {
		seconds %= minutes * oneMinInSeconds
	}

	return fmt.Sprintf("%dd %dh %dm %ds (%s)", days, hours, minutes, seconds, b.Uptime)
}

// GetBrokerDetails subscribes and takes messages from multiple `$SYS/broker/*` topics
func GetBrokerDetails(c mqtt.Client) Broker {
	broker := Broker{}

	broker.Version = string(
		getSubscribedOnce(c, systopic("version")).Payload(),
	)

	broker.Uptime = string(
		getSubscribedOnce(c, systopic("uptime")).Payload(),
	)

	broker.MessagesSent = string(
		getSubscribedOnce(c, systopic("messages/sent")).Payload(),
	)

	broker.MessagesReceived = string(
		getSubscribedOnce(c, systopic("messages/received")).Payload(),
	)

	broker.ClientsConnected = string(
		getSubscribedOnce(c, systopic("clients/connected")).Payload(),
	)

	broker.ClientsTotal = string(
		getSubscribedOnce(c, systopic("clients/total")).Payload(),
	)

	return broker
}
