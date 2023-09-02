package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("content/config.json")
	if err != nil {
		log.Println(err)
	}

	dynamicJson(f)

	// var data map[string]interface{}
	// json.Unmarshal([]byte(f), &data)

	// log.Println(data)
	// for k, v := range data {
	// 	log.Println(k, ":", v)
	// }

	// pingJSON := make(map[string][]mqttconfig)
	// err = json.Unmarshal([]byte(f), &pingJSON)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("\n\n json object:::: %+v", pingJSON)
}

func dynamicJson(input []byte) {

	var msg json.RawMessage
	conf := Config{
		SourceSettings: &msg,
	}

	if err := json.Unmarshal([]byte(input), &conf); err != nil {
		log.Fatal(err)
	}
	switch conf.Source {
	case "mqtt":
		var mqttconf Mqttconfig
		if err := json.Unmarshal(msg, &mqttconf); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("\n\n json object:::: %+v", mqttconf)
		fmt.Println(mqttconf.Host)
		fmt.Println(mqttconf.Port)
		fmt.Println(mqttconf.ClientId)
		fmt.Println(mqttconf.User)
		fmt.Println(mqttconf.Password)
	default:
		log.Fatalf("unknown message type: %q", conf.Source)
	}
}
