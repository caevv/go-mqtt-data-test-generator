package main

import (
	"io/ioutil"
	"time"

	"github.com/pkg/errors"

	"github.com/caevv/go-mqtt-data-test-generator/env"
	"github.com/caevv/go-mqtt-data-test-generator/mqtt"
)

func main() {
	payload, err := ioutil.ReadFile(env.Settings.PayloadFilePath)
	if err != nil {
		panic(errors.Wrapf(err, "failed to read file from path %q", env.Settings.PayloadFilePath))
	}

	client, err := mqtt.New()
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to MQTT broker"))
	}
	defer client.Disconnect()

	for {
		if err := client.Pub(env.Settings.Mqtt.Topic, env.Settings.Mqtt.Qos, env.Settings.Mqtt.Retained, payload); err != nil {
			panic(errors.Wrap(err, "failed to publish to MQTT"))
		}

		time.Sleep(env.Settings.Frequency)
	}
}
