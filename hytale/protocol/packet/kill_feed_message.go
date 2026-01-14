package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type KillFeedMessage struct{}

func (*KillFeedMessage) ID() uint32 {
	return IDKillFeedMessage
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       KillFeedMessage obj = new KillFeedMessage();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 13 + buf.getIntLE(offset + 1);
//          obj.killer = FormattedMessage.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 13 + buf.getIntLE(offset + 5);
//          obj.decedent = FormattedMessage.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 13 + buf.getIntLE(offset + 9);
//          int iconLen = VarInt.peek(buf, varPos2);
//          if (iconLen < 0) {
//             throw ProtocolException.negativeLength("Icon", iconLen);
//          }
//
//          if (iconLen > 4096000) {
//             throw ProtocolException.stringTooLong("Icon", iconLen, 4096000);
//          }
//
//          obj.icon = PacketIO.readVarString(buf, varPos2, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.killer != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.decedent != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.icon != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       int killerOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int decedentOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int iconOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.killer != null) {
//          buf.setIntLE(killerOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.killer.serialize(buf);
//       } else {
//          buf.setIntLE(killerOffsetSlot, -1);
//       }
//
//       if (this.decedent != null) {
//          buf.setIntLE(decedentOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.decedent.serialize(buf);
//       } else {
//          buf.setIntLE(decedentOffsetSlot, -1);
//       }
//
//       if (this.icon != null) {
//          buf.setIntLE(iconOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.icon, 4096000);
//       } else {
//          buf.setIntLE(iconOffsetSlot, -1);
//       }
//    }

func (pk *KillFeedMessage) Marshal(io protocol.IO) {
	// TODO: Implement
}
