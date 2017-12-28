package ul

import (
	"errors"
	"time"
)

const (
	statusLen = 10
	statusPL  = 0x50

	sLoc = "IST"
)

var (
	errInvalidLen = errors.New("status: invalid packet length")
	errInvalidPL  = errors.New("status: invalid payload")
)

// Status is status of light
type Status struct {
	DeviceID            byte      `json:"deviceID"`
	InputPower          byte      `json:"inputPower"`
	Dim                 byte      `json:"dim"`
	OutputVoltage       byte      `json:"outputVoltage"`
	OutputCurrent       float32   `json:"outputCurrent"`
	InternalTemperature byte      `json:"internalTemperatue"`
	Timestamp           time.Time `json:"timestamp"`
}

// NewStatus makes Status and return its pointer
func NewStatus(data []byte) (*Status, error) {
	if len(data) != statusLen {
		return nil, errInvalidLen
	}

	if data[0] != statusPL {
		return nil, errInvalidPL
	}

	return &Status{
		DeviceID:            data[0],
		InputPower:          data[2],
		OutputVoltage:       data[3],
		OutputCurrent:       float32(data[4]) / 10,
		InternalTemperature: data[7],
		Timestamp:           getDeviceTime(data[8], data[9]),
	}, nil
}

func getDeviceTime(hour, minute byte) time.Time {
	loc, _ := time.LoadLocation(sLoc)
	ts := time.Now().In(loc)
	return time.Date(ts.Year(), ts.Month(), ts.Day(),
		int(hour), int(minute), ts.Second(), ts.Nanosecond(), loc)
}
