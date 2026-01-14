package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSelectionUpdate struct{}

func (*BuilderToolSelectionUpdate) ID() uint32 {
	return IDBuilderToolSelectionUpdate
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSelectionUpdate obj = new BuilderToolSelectionUpdate();
//       obj.xMin = buf.getIntLE(offset + 0);
//       obj.yMin = buf.getIntLE(offset + 4);
//       obj.zMin = buf.getIntLE(offset + 8);
//       obj.xMax = buf.getIntLE(offset + 12);
//       obj.yMax = buf.getIntLE(offset + 16);
//       obj.zMax = buf.getIntLE(offset + 20);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.xMin);
//       buf.writeIntLE(this.yMin);
//       buf.writeIntLE(this.zMin);
//       buf.writeIntLE(this.xMax);
//       buf.writeIntLE(this.yMax);
//       buf.writeIntLE(this.zMax);
//    }

func (pk *BuilderToolSelectionUpdate) Marshal(io protocol.IO) {
	// TODO: Implement
}
