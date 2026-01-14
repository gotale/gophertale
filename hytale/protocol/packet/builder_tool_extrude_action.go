package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolExtrudeAction struct{}

func (*BuilderToolExtrudeAction) ID() uint32 {
	return IDBuilderToolExtrudeAction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolExtrudeAction obj = new BuilderToolExtrudeAction();
//       obj.x = buf.getIntLE(offset + 0);
//       obj.y = buf.getIntLE(offset + 4);
//       obj.z = buf.getIntLE(offset + 8);
//       obj.xNormal = buf.getIntLE(offset + 12);
//       obj.yNormal = buf.getIntLE(offset + 16);
//       obj.zNormal = buf.getIntLE(offset + 20);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//       buf.writeIntLE(this.xNormal);
//       buf.writeIntLE(this.yNormal);
//       buf.writeIntLE(this.zNormal);
//    }

func (pk *BuilderToolExtrudeAction) Marshal(io protocol.IO) {
	// TODO: Implement
}
