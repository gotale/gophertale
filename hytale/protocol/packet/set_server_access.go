package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetServerAccess struct{}

func (*SetServerAccess) ID() uint32 {
	return IDSetServerAccess
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetServerAccess obj = new SetServerAccess();
//       byte nullBits = buf.getByte(offset);
//       obj.access = Access.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int passwordLen = VarInt.peek(buf, pos);
//          if (passwordLen < 0) {
//             throw ProtocolException.negativeLength("Password", passwordLen);
//          }
//
//          if (passwordLen > 4096000) {
//             throw ProtocolException.stringTooLong("Password", passwordLen, 4096000);
//          }
//
//          int passwordVarLen = VarInt.length(buf, pos);
//          obj.password = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += passwordVarLen + passwordLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.password != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.access.getValue());
//       if (this.password != null) {
//          PacketIO.writeVarString(buf, this.password, 4096000);
//       }
//    }

func (pk *SetServerAccess) Marshal(io protocol.IO) {
	// TODO: Implement
}
