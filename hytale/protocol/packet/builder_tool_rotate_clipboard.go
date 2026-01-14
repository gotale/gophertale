package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolRotateClipboard struct{}

func (*BuilderToolRotateClipboard) ID() uint32 {
	return IDBuilderToolRotateClipboard
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolRotateClipboard obj = new BuilderToolRotateClipboard();
//       obj.angle = buf.getIntLE(offset + 0);
//       obj.axis = Axis.fromValue(buf.getByte(offset + 4));
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.angle);
//       buf.writeByte(this.axis.getValue());
//    }

func (pk *BuilderToolRotateClipboard) Marshal(io protocol.IO) {
	// TODO: Implement
}
