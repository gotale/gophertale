package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerSetBlock struct{}

func (*ServerSetBlock) ID() uint32 {
	return IDServerSetBlock
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerSetBlock obj = new ServerSetBlock();
//       obj.x = buf.getIntLE(offset + 0);
//       obj.y = buf.getIntLE(offset + 4);
//       obj.z = buf.getIntLE(offset + 8);
//       obj.blockId = buf.getIntLE(offset + 12);
//       obj.filler = buf.getShortLE(offset + 16);
//       obj.rotation = buf.getByte(offset + 18);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//       buf.writeIntLE(this.blockId);
//       buf.writeShortLE(this.filler);
//       buf.writeByte(this.rotation);
//    }

func (pk *ServerSetBlock) Marshal(io protocol.IO) {
	// TODO: Implement
}
