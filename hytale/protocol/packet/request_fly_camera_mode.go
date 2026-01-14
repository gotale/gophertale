package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type RequestFlyCameraMode struct{}

func (*RequestFlyCameraMode) ID() uint32 {
	return IDRequestFlyCameraMode
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       RequestFlyCameraMode obj = new RequestFlyCameraMode();
//       obj.entering = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.entering ? 1 : 0);
//    }

func (pk *RequestFlyCameraMode) Marshal(io protocol.IO) {
	// TODO: Implement
}
