package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolLineAction struct{}

func (*BuilderToolLineAction) ID() uint32 {
	return IDBuilderToolLineAction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolLineAction obj = new BuilderToolLineAction();
//       obj.xStart = buf.getIntLE(offset + 0);
//       obj.yStart = buf.getIntLE(offset + 4);
//       obj.zStart = buf.getIntLE(offset + 8);
//       obj.xEnd = buf.getIntLE(offset + 12);
//       obj.yEnd = buf.getIntLE(offset + 16);
//       obj.zEnd = buf.getIntLE(offset + 20);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.xStart);
//       buf.writeIntLE(this.yStart);
//       buf.writeIntLE(this.zStart);
//       buf.writeIntLE(this.xEnd);
//       buf.writeIntLE(this.yEnd);
//       buf.writeIntLE(this.zEnd);
//    }

func (pk *BuilderToolLineAction) Marshal(io protocol.IO) {
	// TODO: Implement
}
