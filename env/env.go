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
	ClientLog         bool
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

	Settings = &Config{
		Frequency:       viper.GetDuration("FREQUENCY"),
		PayloadFilePath: viper.GetString("PAYLOAD_FILE_PATH"),

		Mqtt: Mqtt{
			ClientLog:         viper.GetBool("MQTT_CLIENT_LOG"),
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
