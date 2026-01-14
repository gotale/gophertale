package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdatePortal struct{}

func (*UpdatePortal) ID() uint32 {
	return IDUpdatePortal
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdatePortal obj = new UpdatePortal();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.state = PortalState.deserialize(buf, offset + 1);
//       }
//
//       int pos = offset + 6;
//       if ((nullBits & 2) != 0) {
//          obj.definition = PortalDef.deserialize(buf, pos);
//          pos += PortalDef.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.state != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.definition != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.state != null) {
//          this.state.serialize(buf);
//       } else {
//          buf.writeZero(5);
//       }
//
//       if (this.definition != null) {
//          this.definition.serialize(buf);
//       }
//    }

func (pk *UpdatePortal) Marshal(io protocol.IO) {
	// TODO: Implement
}
