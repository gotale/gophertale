package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetMovementStates struct {
	MovementStates SavedMovementStates
}

type SavedMovementStates struct {
	Flying bool
}

func (x *SavedMovementStates) Marshal(io protocol.IO) {
	io.Bool(&x.Flying)
}

func (*SetMovementStates) ID() uint32 {
	return IDSetMovementStates
}

func (pk *SetMovementStates) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	opts.Has(pk.MovementStates != (SavedMovementStates{}))
	opts.Write()

	protocol.Single(io, &pk.MovementStates)
}
