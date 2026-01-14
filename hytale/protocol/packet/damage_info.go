package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type DamageInfo struct{}

func (*DamageInfo) ID() uint32 {
	return IDDamageInfo
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       DamageInfo obj = new DamageInfo();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.damageSourcePosition = Vector3d.deserialize(buf, offset + 1);
//       }
//
//       obj.damageAmount = buf.getFloatLE(offset + 25);
//       int pos = offset + 29;
//       if ((nullBits & 2) != 0) {
//          obj.damageCause = DamageCause.deserialize(buf, pos);
//          pos += DamageCause.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.damageSourcePosition != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.damageCause != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.damageSourcePosition != null) {
//          this.damageSourcePosition.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//
//       buf.writeFloatLE(this.damageAmount);
//       if (this.damageCause != null) {
//          this.damageCause.serialize(buf);
//       }
//    }

func (pk *DamageInfo) Marshal(io protocol.IO) {
	// TODO: Implement
}
