package mqtt

import (
	"log"
	"sync"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/lcn/serial"
)

const (
	brokerAddress = "tcp://localhost:1883"
	username      = "senra"
	password      = "sc2havellssenra"

	ulMsgChan = "/payloads/uplink/"
	dlMsgChan = "/payloads/downlink"
)

var wg sync.WaitGroup
var cl mqtt.Client

func init() {
	initClient()
}

func initClient() {
	opt := mqtt.NewClientOptions()
	opt = opt.AddBroker(brokerAddress)
	opt = opt.SetUsername(username)
	opt = opt.SetPassword(password)
	cl = mqtt.NewClient(opt)

	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
}

// StartSubscription starts receiving data
func StartSubscription() {
	wg.Add(1)
	//handle downlink
	if token := cl.Subscribe(dlMsgChan, 0, handleDownlink); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
	wg.Wait()
}

func handleDownlink(cl mqtt.Client, msg mqtt.Message) {
	log.Printf("message received % x\n", msg.Payload())
	serial.Send(msg.Payload())
	//	wg.Done()
}

// Send publishes data
func Send(msgChan string, data []byte) {
	if token := cl.Publish(ulMsgChan+msgChan, 0, true, data); token.Wait() && token.Error() != nil {
		log.Printf("error in send: %s \n", token.Error())
		return
	}
	log.Println("DATA SENT SUCCESSFULLY")
}
