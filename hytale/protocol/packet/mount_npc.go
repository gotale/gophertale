package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type MountNPC struct{}

func (*MountNPC) ID() uint32 {
	return IDMountNPC
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       MountNPC obj = new MountNPC();
//       obj.anchorX = buf.getFloatLE(offset + 0);
//       obj.anchorY = buf.getFloatLE(offset + 4);
//       obj.anchorZ = buf.getFloatLE(offset + 8);
//       obj.entityId = buf.getIntLE(offset + 12);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeFloatLE(this.anchorX);
//       buf.writeFloatLE(this.anchorY);
//       buf.writeFloatLE(this.anchorZ);
//       buf.writeIntLE(this.entityId);
//    }

func (pk *MountNPC) Marshal(io protocol.IO) {
	// TODO: Implement
}
