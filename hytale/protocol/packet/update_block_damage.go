package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockDamage struct{}

func (*UpdateBlockDamage) ID() uint32 {
	return IDUpdateBlockDamage
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockDamage obj = new UpdateBlockDamage();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.blockPosition = BlockPosition.deserialize(buf, offset + 1);
//       }
//
//       obj.damage = buf.getFloatLE(offset + 13);
//       obj.delta = buf.getFloatLE(offset + 17);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.blockPosition != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.blockPosition != null) {
//          this.blockPosition.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       buf.writeFloatLE(this.damage);
//       buf.writeFloatLE(this.delta);
//    }

func (pk *UpdateBlockDamage) Marshal(io protocol.IO) {
	// TODO: Implement
}
