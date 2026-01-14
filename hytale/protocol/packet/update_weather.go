package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateWeather struct {
	WeatherIndex      int32
	TransitionSeconds float32
}

func (*UpdateWeather) ID() uint32 {
	return IDUpdateWeather
}

func (pk *UpdateWeather) Marshal(io protocol.IO) {
	io.Int32(&pk.WeatherIndex)
	io.Float32(&pk.TransitionSeconds)
}
