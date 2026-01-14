package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type RequestServerAccess struct{}

func (*RequestServerAccess) ID() uint32 {
	return IDRequestServerAccess
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       RequestServerAccess obj = new RequestServerAccess();
//       obj.access = Access.fromValue(buf.getByte(offset + 0));
//       obj.externalPort = buf.getShortLE(offset + 1);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.access.getValue());
//       buf.writeShortLE(this.externalPort);
//    }

func (pk *RequestServerAccess) Marshal(io protocol.IO) {
	// TODO: Implement
}
