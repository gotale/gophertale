package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateWeathers struct {
	UpdateType uint8
	MaxID      int32
	Weathers   map[int32]protocol.Weather
}

func (*UpdateWeathers) ID() uint32 {
	return IDUpdateWeathers
}

func (pk *UpdateWeathers) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasWeathers := opts.Has(len(pk.Weathers) > 0)
	opts.Write()

	io.Uint8(&pk.UpdateType)
	io.Int32(&pk.MaxID)
	if hasWeathers {
		protocol.Map(io, &pk.Weathers, 4096000, io.Int32)
	}
}
