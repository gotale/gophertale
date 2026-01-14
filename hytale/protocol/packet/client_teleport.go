package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClientTeleport struct{}

func (*ClientTeleport) ID() uint32 {
	return IDClientTeleport
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ClientTeleport obj = new ClientTeleport();
//       byte nullBits = buf.getByte(offset);
//       obj.teleportId = buf.getByte(offset + 1);
//       if ((nullBits & 1) != 0) {
//          obj.modelTransform = ModelTransform.deserialize(buf, offset + 2);
//       }
//
//       obj.resetVelocity = buf.getByte(offset + 51) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.modelTransform != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.teleportId);
//       if (this.modelTransform != null) {
//          this.modelTransform.serialize(buf);
//       } else {
//          buf.writeZero(49);
//       }
//
//       buf.writeByte(this.resetVelocity ? 1 : 0);
//    }

func (pk *ClientTeleport) Marshal(io protocol.IO) {
	// TODO: Implement
}
