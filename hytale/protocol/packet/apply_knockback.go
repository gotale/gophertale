package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ApplyKnockback struct{}

func (*ApplyKnockback) ID() uint32 {
	return IDApplyKnockback
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ApplyKnockback obj = new ApplyKnockback();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.hitPosition = Position.deserialize(buf, offset + 1);
//       }
//
//       obj.x = buf.getFloatLE(offset + 25);
//       obj.y = buf.getFloatLE(offset + 29);
//       obj.z = buf.getFloatLE(offset + 33);
//       obj.changeType = ChangeVelocityType.fromValue(buf.getByte(offset + 37));
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.hitPosition != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.hitPosition != null) {
//          this.hitPosition.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//
//       buf.writeFloatLE(this.x);
//       buf.writeFloatLE(this.y);
//       buf.writeFloatLE(this.z);
//       buf.writeByte(this.changeType.getValue());
//    }

func (pk *ApplyKnockback) Marshal(io protocol.IO) {
	// TODO: Implement
}
