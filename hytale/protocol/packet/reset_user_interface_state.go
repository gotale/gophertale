package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ResetUserInterfaceState struct{}

func (*ResetUserInterfaceState) ID() uint32 {
	return IDResetUserInterfaceState
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new ResetUserInterfaceState();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *ResetUserInterfaceState) Marshal(io protocol.IO) {
	// TODO: Implement
}
