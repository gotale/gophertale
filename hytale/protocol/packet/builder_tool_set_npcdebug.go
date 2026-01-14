package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSetNPCDebug struct{}

func (*BuilderToolSetNPCDebug) ID() uint32 {
	return IDBuilderToolSetNPCDebug
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSetNPCDebug obj = new BuilderToolSetNPCDebug();
//       obj.entityId = buf.getIntLE(offset + 0);
//       obj.enabled = buf.getByte(offset + 4) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.entityId);
//       buf.writeByte(this.enabled ? 1 : 0);
//    }

func (pk *BuilderToolSetNPCDebug) Marshal(io protocol.IO) {
	// TODO: Implement
}
