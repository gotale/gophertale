package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClearWorldMap struct{}

func (*ClearWorldMap) ID() uint32 {
	return IDClearWorldMap
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new ClearWorldMap();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *ClearWorldMap) Marshal(io protocol.IO) {
	// TODO: Implement
}
