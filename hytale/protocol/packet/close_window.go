package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type CloseWindow struct{}

func (*CloseWindow) ID() uint32 {
	return IDCloseWindow
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       CloseWindow obj = new CloseWindow();
//       obj.id = buf.getIntLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.id);
//    }

func (pk *CloseWindow) Marshal(io protocol.IO) {
	// TODO: Implement
}
