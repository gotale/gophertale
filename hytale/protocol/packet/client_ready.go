package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClientReady struct{}

func (*ClientReady) ID() uint32 {
	return IDClientReady
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ClientReady obj = new ClientReady();
//       obj.readyForChunks = buf.getByte(offset + 0) != 0;
//       obj.readyForGameplay = buf.getByte(offset + 1) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.readyForChunks ? 1 : 0);
//       buf.writeByte(this.readyForGameplay ? 1 : 0);
//    }

func (pk *ClientReady) Marshal(io protocol.IO) {
	// TODO: Implement
}
