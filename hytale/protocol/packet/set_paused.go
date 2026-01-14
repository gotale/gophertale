package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetPaused struct{}

func (*SetPaused) ID() uint32 {
	return IDSetPaused
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetPaused obj = new SetPaused();
//       obj.paused = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.paused ? 1 : 0);
//    }

func (pk *SetPaused) Marshal(io protocol.IO) {
	// TODO: Implement
}
