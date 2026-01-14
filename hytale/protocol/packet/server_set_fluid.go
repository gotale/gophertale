package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerSetFluid struct{}

func (*ServerSetFluid) ID() uint32 {
	return IDServerSetFluid
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerSetFluid obj = new ServerSetFluid();
//       obj.x = buf.getIntLE(offset + 0);
//       obj.y = buf.getIntLE(offset + 4);
//       obj.z = buf.getIntLE(offset + 8);
//       obj.fluidId = buf.getIntLE(offset + 12);
//       obj.fluidLevel = buf.getByte(offset + 16);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//       buf.writeIntLE(this.fluidId);
//       buf.writeByte(this.fluidLevel);
//    }

func (pk *ServerSetFluid) Marshal(io protocol.IO) {
	// TODO: Implement
}
