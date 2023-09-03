package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Topic: %s | %s\n", msg.Topic(), msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %+v", err)
}

func connect(mqttConfig Mqttconfig) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%d", mqttConfig.Host, mqttConfig.Port))
	opts.SetClientID(mqttConfig.ClientId)
	opts.SetUsername(mqttConfig.User)
	opts.SetPassword(mqttConfig.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.SetAutoReconnect(true)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client, mqttConfig.Topic)
}

func sub(client mqtt.Client, topic string) {
	// Subscribe to the LWT connection status

	token := client.Subscribe(topic, 0, nil)

	if token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error()).Msg("")
		os.Exit(1)
	}

	log.Info().Msgf("Subscribed to %s", topic)
}

func publish(client mqtt.Client) {
	// Go to PTZ preset 2 and return to 1 after 15s
	nums := []int{2, 1}
	for _, num := range nums {
		value := fmt.Sprintf("%d", num)
		token := client.Publish("cameras/77/features/ptz/preset/raw", 0, false, value)
		token.Wait()
		time.Sleep(15 * time.Second)
	}
}
