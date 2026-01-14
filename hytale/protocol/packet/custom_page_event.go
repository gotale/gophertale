package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type CustomPageEvent struct{}

func (*CustomPageEvent) ID() uint32 {
	return IDCustomPageEvent
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       CustomPageEvent obj = new CustomPageEvent();
//       byte nullBits = buf.getByte(offset);
//       obj.type = CustomPageEventType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int dataLen = VarInt.peek(buf, pos);
//          if (dataLen < 0) {
//             throw ProtocolException.negativeLength("Data", dataLen);
//          }
//
//          if (dataLen > 4096000) {
//             throw ProtocolException.stringTooLong("Data", dataLen, 4096000);
//          }
//
//          int dataVarLen = VarInt.length(buf, pos);
//          obj.data = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += dataVarLen + dataLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.data != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.data != null) {
//          PacketIO.writeVarString(buf, this.data, 4096000);
//       }
//    }

func (pk *CustomPageEvent) Marshal(io protocol.IO) {
	// TODO: Implement
}
