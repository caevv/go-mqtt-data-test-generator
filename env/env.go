package env

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Frequency       time.Duration
	PayloadFilePath string
	Mqtt            Mqtt
}

type Mqtt struct {
	Broker            string
	CleanSession      bool
	ClientID          string
	OrderMatters      bool
	WriteTimeout      time.Duration
	ConnectionTimeout time.Duration
	KeepAlive         time.Duration
	AutoReconnect     bool
	Topic             string
	Qos               int
	Retained          bool
}

var Settings *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetDefault("FREQUENCY", "1s")
	viper.SetDefault("MQTT_CLEAN_SESSION", "true")
	viper.SetDefault("MQTT_ORDER_MATTERS", "true")
	viper.SetDefault("MQTT_WRITE_TIMEOUT", "1s")
	viper.SetDefault("MQTT_CONNECTION_TIMEOUT", "1s")
	viper.SetDefault("MQTT_KEEP_ALIVE", "2s")
	viper.SetDefault("MQTT_AUTO_RECONNECT", "false")
	viper.SetDefault("MQTT_QOS", "1")
	viper.SetDefault("MQTT_RETAINED_MESSAGE", "false")

	Settings = &Config{
		Frequency:       viper.GetDuration("FREQUENCY"),
		PayloadFilePath: viper.GetString("PAYLOAD_FILE_PATH"),

		Mqtt: Mqtt{
			Broker:            viper.GetString("MQTT_BROKER"),
			CleanSession:      viper.GetBool("MQTT_CLEAN_SESSION"),
			ClientID:          viper.GetString("MQTT_CLIENT_ID"),
			OrderMatters:      viper.GetBool("MQTT_ORDER_MATTERS"),
			WriteTimeout:      viper.GetDuration("MQTT_WRITE_TIMEOUT"),
			ConnectionTimeout: viper.GetDuration("MQTT_CONNECTION_TIMEOUT"),
			KeepAlive:         viper.GetDuration("MQTT_KEEP_ALIVE"),
			AutoReconnect:     viper.GetBool("MQTT_AUTO_RECONNECT"),
			Topic:             viper.GetString("MQTT_TOPIC"),
			Qos:               viper.GetInt("MQTT_QOS"),
			Retained:          viper.GetBool("MQTT_RETAINED_MESSAGE"),
		},
	}
}
