package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetBlockPlacementOverride struct{}

func (*SetBlockPlacementOverride) ID() uint32 {
	return IDSetBlockPlacementOverride
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetBlockPlacementOverride obj = new SetBlockPlacementOverride();
//       obj.enabled = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.enabled ? 1 : 0);
//    }

func (pk *SetBlockPlacementOverride) Marshal(io protocol.IO) {
	// TODO: Implement
}
