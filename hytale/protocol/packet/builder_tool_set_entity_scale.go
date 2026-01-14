package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSetEntityScale struct{}

func (*BuilderToolSetEntityScale) ID() uint32 {
	return IDBuilderToolSetEntityScale
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSetEntityScale obj = new BuilderToolSetEntityScale();
//       obj.entityId = buf.getIntLE(offset + 0);
//       obj.scale = buf.getFloatLE(offset + 4);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.entityId);
//       buf.writeFloatLE(this.scale);
//    }

func (pk *BuilderToolSetEntityScale) Marshal(io protocol.IO) {
	// TODO: Implement
}
