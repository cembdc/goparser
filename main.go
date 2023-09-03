package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	f, err := os.ReadFile("content/config.json")
	if err != nil {
		log.Println(err)
	}

	dynamicJson(f)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal
}

func dynamicJson(input []byte) {

	var agent Agent

	if err := json.Unmarshal([]byte(input), &agent); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("\n\n json object:::: %+v", agent)

	for i := 0; i < len(agent.Configs); i++ {
		// var msg json.RawMessage
		// aa := agent.Configs[i]{
		// 	SourceSettings: &msg,
		// }

		conf := agent.Configs[i]
		// conf.SourceSettings = &msg

		switch conf.Source {
		case "mqtt":
			var mqttconf Mqttconfig
			jsonString, _ := json.Marshal(conf.SourceSettings)

			if err := json.Unmarshal(jsonString, &mqttconf); err != nil {
				log.Fatal(err)
			}
			fmt.Println(conf.Direction)
			fmt.Println(mqttconf.Host)
			fmt.Println(mqttconf.Port)
			fmt.Println(mqttconf.ClientId)
			fmt.Println(mqttconf.Topic)
			fmt.Println(mqttconf.User)
			fmt.Println(mqttconf.Password)

			go connect(mqttconf)
		default:
			log.Fatalf("unknown message type: %q", conf.Source)
		}
	}

}
