package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateTime struct {
	GameTime protocol.InstantData
}

func (*UpdateTime) ID() uint32 {
	return IDUpdateTime
}

func (pk *UpdateTime) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	opts.Has(pk.GameTime != protocol.InstantData{})
	opts.Write()

	protocol.Single(io, &pk.GameTime)
}
