package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerSetPaused struct{}

func (*ServerSetPaused) ID() uint32 {
	return IDServerSetPaused
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerSetPaused obj = new ServerSetPaused();
//       obj.paused = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.paused ? 1 : 0);
//    }

func (pk *ServerSetPaused) Marshal(io protocol.IO) {
	// TODO: Implement
}
