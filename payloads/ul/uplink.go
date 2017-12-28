package ul

import (
	"log"

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
			mqtt.Send(su)
		}
	}
}
