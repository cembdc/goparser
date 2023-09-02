package main

type Mqttconfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	ClientId string `json:"clientId"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Direction      string `json:"direction"`
	Source         string `json:"source"`
	SourceSettings interface{}
}
