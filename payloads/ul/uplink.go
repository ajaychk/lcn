package ul

import (
	"encoding/json"
	"log"

	"github.com/lcn/comm/mqtt"
	"github.com/lcn/serial"
)

func init() {
	go processUL()
}

func processUL() {
	for {
		data := <-serial.ChanUplink
		switch data[1] {
		case statusPL:
			su, err := NewStatus(data)
			if err != nil {
				log.Println(err)
				continue
			}
			bData, err := json.Marshal(su)
			if err != nil {
				log.Println(err)
				continue
			}
			mqtt.Send("status", bData)
		}
	}
}
