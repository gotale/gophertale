package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ViewRadius struct {
	Value int32
}

func (*ViewRadius) ID() uint32 {
	return IDViewRadius
}

func (pk *ViewRadius) Marshal(io protocol.IO) {
	io.Int32(&pk.Value)
}
