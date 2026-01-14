package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type HideEventTitle struct{}

func (*HideEventTitle) ID() uint32 {
	return IDHideEventTitle
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       HideEventTitle obj = new HideEventTitle();
//       obj.fadeOutDuration = buf.getFloatLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeFloatLE(this.fadeOutDuration);
//    }

func (pk *HideEventTitle) Marshal(io protocol.IO) {
	// TODO: Implement
}
