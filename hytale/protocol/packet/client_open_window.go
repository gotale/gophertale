package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClientOpenWindow struct{}

func (*ClientOpenWindow) ID() uint32 {
	return IDClientOpenWindow
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ClientOpenWindow obj = new ClientOpenWindow();
//       obj.type = WindowType.fromValue(buf.getByte(offset + 0));
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.type.getValue());
//    }

func (pk *ClientOpenWindow) Marshal(io protocol.IO) {
	// TODO: Implement
}
