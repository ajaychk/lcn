package main

import (
	"log"

	"github.com/lcn/comm/mqtt"
)

func main() {
	log.Println("LCN SERVER STARTED")
	log.Println("COMMUNICATION TO LIGHT ON LCN")
	mqtt.StartSubscription()
	log.Println("SHUTING DOWN SERVER")
}
