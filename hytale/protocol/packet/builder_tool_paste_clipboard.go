package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolPasteClipboard struct{}

func (*BuilderToolPasteClipboard) ID() uint32 {
	return IDBuilderToolPasteClipboard
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolPasteClipboard obj = new BuilderToolPasteClipboard();
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

func (pk *BuilderToolPasteClipboard) Marshal(io protocol.IO) {
	// TODO: Implement
}
