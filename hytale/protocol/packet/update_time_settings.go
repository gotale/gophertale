package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateTimeSettings struct {
	DayTimeDurationSeconds   int32
	NightTimeDurationSeconds int32
	TotalMoonPhases          uint8
	TimePaused               bool
}

func (*UpdateTimeSettings) ID() uint32 {
	return IDUpdateTimeSettings
}

func (pk *UpdateTimeSettings) Marshal(io protocol.IO) {
	io.Int32(&pk.DayTimeDurationSeconds)
	io.Int32(&pk.NightTimeDurationSeconds)
	io.Uint8(&pk.TotalMoonPhases)
	io.Bool(&pk.TimePaused)
}
