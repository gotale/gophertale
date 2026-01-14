package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetTimeDilation struct {
	TimeDilation float32
}

func (*SetTimeDilation) ID() uint32 {
	return IDSetTimeDilation
}

func (pk *SetTimeDilation) Marshal(io protocol.IO) {
	io.Float32(&pk.TimeDilation)
}
