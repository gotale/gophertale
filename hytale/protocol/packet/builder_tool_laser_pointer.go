package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolLaserPointer struct{}

func (*BuilderToolLaserPointer) ID() uint32 {
	return IDBuilderToolLaserPointer
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolLaserPointer obj = new BuilderToolLaserPointer();
//       obj.playerNetworkId = buf.getIntLE(offset + 0);
//       obj.startX = buf.getFloatLE(offset + 4);
//       obj.startY = buf.getFloatLE(offset + 8);
//       obj.startZ = buf.getFloatLE(offset + 12);
//       obj.endX = buf.getFloatLE(offset + 16);
//       obj.endY = buf.getFloatLE(offset + 20);
//       obj.endZ = buf.getFloatLE(offset + 24);
//       obj.color = buf.getIntLE(offset + 28);
//       obj.durationMs = buf.getIntLE(offset + 32);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.playerNetworkId);
//       buf.writeFloatLE(this.startX);
//       buf.writeFloatLE(this.startY);
//       buf.writeFloatLE(this.startZ);
//       buf.writeFloatLE(this.endX);
//       buf.writeFloatLE(this.endY);
//       buf.writeFloatLE(this.endZ);
//       buf.writeIntLE(this.color);
//       buf.writeIntLE(this.durationMs);
//    }

func (pk *BuilderToolLaserPointer) Marshal(io protocol.IO) {
	// TODO: Implement
}
