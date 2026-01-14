package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type Notification struct{}

func (*Notification) ID() uint32 {
	return IDNotification
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       Notification obj = new Notification();
//       byte nullBits = buf.getByte(offset);
//       obj.style = NotificationStyle.fromValue(buf.getByte(offset + 1));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 18 + buf.getIntLE(offset + 2);
//          obj.message = FormattedMessage.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 18 + buf.getIntLE(offset + 6);
//          obj.secondaryMessage = FormattedMessage.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 18 + buf.getIntLE(offset + 10);
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
//       if ((nullBits & 8) != 0) {
//          int varPos3 = offset + 18 + buf.getIntLE(offset + 14);
//          obj.item = ItemWithAllMetadata.deserialize(buf, varPos3);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.message != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.secondaryMessage != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.icon != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       if (this.item != null) {
//          nullBits = (byte)(nullBits | 8);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.style.getValue());
//       int messageOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int secondaryMessageOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int iconOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int itemOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.message != null) {
//          buf.setIntLE(messageOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.message.serialize(buf);
//       } else {
//          buf.setIntLE(messageOffsetSlot, -1);
//       }
//
//       if (this.secondaryMessage != null) {
//          buf.setIntLE(secondaryMessageOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.secondaryMessage.serialize(buf);
//       } else {
//          buf.setIntLE(secondaryMessageOffsetSlot, -1);
//       }
//
//       if (this.icon != null) {
//          buf.setIntLE(iconOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.icon, 4096000);
//       } else {
//          buf.setIntLE(iconOffsetSlot, -1);
//       }
//
//       if (this.item != null) {
//          buf.setIntLE(itemOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.item.serialize(buf);
//       } else {
//          buf.setIntLE(itemOffsetSlot, -1);
//       }
//    }

func (pk *Notification) Marshal(io protocol.IO) {
	// TODO: Implement
}
