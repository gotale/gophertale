package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSetEntityPickupEnabled struct{}

func (*BuilderToolSetEntityPickupEnabled) ID() uint32 {
	return IDBuilderToolSetEntityPickupEnabled
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSetEntityPickupEnabled obj = new BuilderToolSetEntityPickupEnabled();
//       obj.entityId = buf.getIntLE(offset + 0);
//       obj.enabled = buf.getByte(offset + 4) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.entityId);
//       buf.writeByte(this.enabled ? 1 : 0);
//    }

func (pk *BuilderToolSetEntityPickupEnabled) Marshal(io protocol.IO) {
	// TODO: Implement
}
