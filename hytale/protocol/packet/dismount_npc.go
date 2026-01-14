package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type DismountNPC struct{}

func (*DismountNPC) ID() uint32 {
	return IDDismountNPC
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new DismountNPC();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *DismountNPC) Marshal(io protocol.IO) {
	// TODO: Implement
}
