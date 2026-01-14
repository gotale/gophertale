package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetFlyCameraMode struct{}

func (*SetFlyCameraMode) ID() uint32 {
	return IDSetFlyCameraMode
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetFlyCameraMode obj = new SetFlyCameraMode();
//       obj.entering = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.entering ? 1 : 0);
//    }

func (pk *SetFlyCameraMode) Marshal(io protocol.IO) {
	// TODO: Implement
}
