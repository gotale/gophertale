package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ShowEventTitle struct{}

func (*ShowEventTitle) ID() uint32 {
	return IDShowEventTitle
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ShowEventTitle obj = new ShowEventTitle();
//       byte nullBits = buf.getByte(offset);
//       obj.fadeInDuration = buf.getFloatLE(offset + 1);
//       obj.fadeOutDuration = buf.getFloatLE(offset + 5);
//       obj.duration = buf.getFloatLE(offset + 9);
//       obj.isMajor = buf.getByte(offset + 13) != 0;
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 26 + buf.getIntLE(offset + 14);
//          int iconLen = VarInt.peek(buf, varPos0);
//          if (iconLen < 0) {
//             throw ProtocolException.negativeLength("Icon", iconLen);
//          }
//
//          if (iconLen > 4096000) {
//             throw ProtocolException.stringTooLong("Icon", iconLen, 4096000);
//          }
//
//          obj.icon = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 26 + buf.getIntLE(offset + 18);
//          obj.primaryTitle = FormattedMessage.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 26 + buf.getIntLE(offset + 22);
//          obj.secondaryTitle = FormattedMessage.deserialize(buf, varPos2);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.icon != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.primaryTitle != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.secondaryTitle != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeFloatLE(this.fadeInDuration);
//       buf.writeFloatLE(this.fadeOutDuration);
//       buf.writeFloatLE(this.duration);
//       buf.writeByte(this.isMajor ? 1 : 0);
//       int iconOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int primaryTitleOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int secondaryTitleOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.icon != null) {
//          buf.setIntLE(iconOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.icon, 4096000);
//       } else {
//          buf.setIntLE(iconOffsetSlot, -1);
//       }
//
//       if (this.primaryTitle != null) {
//          buf.setIntLE(primaryTitleOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.primaryTitle.serialize(buf);
//       } else {
//          buf.setIntLE(primaryTitleOffsetSlot, -1);
//       }
//
//       if (this.secondaryTitle != null) {
//          buf.setIntLE(secondaryTitleOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.secondaryTitle.serialize(buf);
//       } else {
//          buf.setIntLE(secondaryTitleOffsetSlot, -1);
//       }
//    }

func (pk *ShowEventTitle) Marshal(io protocol.IO) {
	// TODO: Implement
}
