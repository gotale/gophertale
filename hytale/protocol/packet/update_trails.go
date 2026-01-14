package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateTrails struct {
	UpdateType uint8
	Trails     map[string]Trail
}

func (*UpdateTrails) ID() uint32 {
	return IDUpdateTrails
}

func (pk *UpdateTrails) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasTrails := opts.Has(len(pk.Trails) > 0)
	opts.Write()

	io.Uint8(&pk.UpdateType)
	if hasTrails {
		protocol.Map(io, &pk.Trails, 4096000, func(s *string) {
			io.VarString(s, 4096000)
		})
	}
}

type Trail struct{}

func (x *Trail) Marshal(io protocol.IO) {
	// TODO: Implementation
}
