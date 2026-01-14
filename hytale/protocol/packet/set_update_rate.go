package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetUpdateRate struct {
	UpdatesPersecond int32
}

func (*SetUpdateRate) ID() uint32 {
	return IDSetUpdateRate
}

func (pk *SetUpdateRate) Marshal(io protocol.IO) {
	io.Int32(&pk.UpdatesPersecond)
}
