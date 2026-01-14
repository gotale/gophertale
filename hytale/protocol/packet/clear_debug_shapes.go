package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClearDebugShapes struct{}

func (*ClearDebugShapes) ID() uint32 {
	return IDClearDebugShapes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new ClearDebugShapes();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *ClearDebugShapes) Marshal(io protocol.IO) {
	// TODO: Implement
}
