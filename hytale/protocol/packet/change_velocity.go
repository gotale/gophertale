package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ChangeVelocity struct{}

func (*ChangeVelocity) ID() uint32 {
	return IDChangeVelocity
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ChangeVelocity obj = new ChangeVelocity();
//       byte nullBits = buf.getByte(offset);
//       obj.x = buf.getFloatLE(offset + 1);
//       obj.y = buf.getFloatLE(offset + 5);
//       obj.z = buf.getFloatLE(offset + 9);
//       obj.changeType = ChangeVelocityType.fromValue(buf.getByte(offset + 13));
//       if ((nullBits & 1) != 0) {
//          obj.config = VelocityConfig.deserialize(buf, offset + 14);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.config != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeFloatLE(this.x);
//       buf.writeFloatLE(this.y);
//       buf.writeFloatLE(this.z);
//       buf.writeByte(this.changeType.getValue());
//       if (this.config != null) {
//          this.config.serialize(buf);
//       } else {
//          buf.writeZero(21);
//       }
//    }

func (pk *ChangeVelocity) Marshal(io protocol.IO) {
	// TODO: Implement
}
