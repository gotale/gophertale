package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolShowAnchor struct{}

func (*BuilderToolShowAnchor) ID() uint32 {
	return IDBuilderToolShowAnchor
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolShowAnchor obj = new BuilderToolShowAnchor();
//       obj.x = buf.getIntLE(offset + 0);
//       obj.y = buf.getIntLE(offset + 4);
//       obj.z = buf.getIntLE(offset + 8);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//    }

func (pk *BuilderToolShowAnchor) Marshal(io protocol.IO) {
	// TODO: Implement
}
