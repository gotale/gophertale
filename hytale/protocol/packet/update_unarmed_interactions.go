package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateUnarmedInteractions struct {
	UpdateType   uint8
	Interactions map[uint8]int32
}

func (*UpdateUnarmedInteractions) ID() uint32 {
	return IDUpdateUnarmedInteractions
}

func (pk *UpdateUnarmedInteractions) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasInteractions := opts.Has(len(pk.Interactions) > 0)
	opts.Write()

	io.Uint8(&pk.UpdateType)
	if hasInteractions {
		protocol.FuncMap(io, &pk.Interactions, 4096000, io.Uint8, io.Int32)
	}
}
