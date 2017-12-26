package main

import (
	"log"
	"sync"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/lcn/serial"
)

var wg sync.WaitGroup
var cl mqtt.Client

func main() {
	initClient()
}

func initClient() {
	opt := mqtt.NewClientOptions()
	opt = opt.AddBroker("tcp://localhost:1883")
	opt = opt.SetUsername("senra")
	opt = opt.SetPassword("sc2havellssenra")
	cl = mqtt.NewClient(opt)

	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	wg.Add(1)
	//handle downlink
	if token := cl.Subscribe("/lcn/payloads/downlink", 0, handleDownlink); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
	wg.Wait()

	go sendStatus()
}

func handleDownlink(cl mqtt.Client, msg mqtt.Message) {
	log.Printf("message received % x\n", msg.Payload())
	serial.Send(msg.Payload())
	//	wg.Done()
}

func sendStatus() {
	for {
		data := <-serial.ChanStatus

		if token := cl.Publish("/lcn/payloads/uplink", 0, true, data); token.Wait() && token.Error() != nil {
			log.Fatalf("error in send: %s", token.Error())
		}
	}
}
