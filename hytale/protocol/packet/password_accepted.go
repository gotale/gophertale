package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PasswordAccepted struct{}

func (*PasswordAccepted) ID() uint32 {
	return IDPasswordAccepted
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new PasswordAccepted();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *PasswordAccepted) Marshal(io protocol.IO) {
	// TODO: Implement
}
