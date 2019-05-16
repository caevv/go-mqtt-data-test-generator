package mqtt

import (
	"log"

	pmqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"

	"github.com/caevv/go-mqtt-data-test-generator/env"
)

type Client struct {
	client pmqtt.Client
}

func New() (*Client, error) {
	clientOptions := pmqtt.NewClientOptions().
		AddBroker(env.Settings.Mqtt.Broker).
		SetCleanSession(env.Settings.Mqtt.CleanSession).
		SetClientID(env.Settings.Mqtt.ClientID).
		SetOrderMatters(env.Settings.Mqtt.OrderMatters).
		SetWriteTimeout(env.Settings.Mqtt.WriteTimeout).
		SetConnectTimeout(env.Settings.Mqtt.ConnectionTimeout).
		SetKeepAlive(env.Settings.Mqtt.KeepAlive).
		SetAutoReconnect(env.Settings.Mqtt.AutoReconnect)

	client := pmqtt.NewClient(clientOptions)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, errors.Wrap(token.Error(), "failed to connect to MQTT")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Pub(topic string, qos int, retained bool, payload interface{}) error {
	token := c.client.Publish(topic, byte(qos), retained, payload)

	if token.Wait(); token.Error() != nil {
		return errors.Wrap(token.Error(), "subscriber token failed")
	}

	return nil
}

func (c *Client) Disconnect() {
	log.Print("disconnecting from MQTT")
	c.client.Disconnect(250)
}
