package main

import (
	"io/ioutil"
	"time"

	"github.com/caevv/go-mqtt-data-test-generator/env"
	"github.com/caevv/go-mqtt-data-test-generator/mqtt"
)

func main() {
	payload, err := ioutil.ReadFile(env.Settings.PayloadFilePath)

	client, err := mqtt.New()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

	for {
		if err := client.Pub(env.Settings.Mqtt.Topic, env.Settings.Mqtt.Qos, env.Settings.Mqtt.Retained, payload); err != nil {
			panic(err)
		}

		time.Sleep(env.Settings.Frequency)
	}
}
