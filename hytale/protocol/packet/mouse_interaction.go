package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type MouseInteraction struct{}

func (*MouseInteraction) ID() uint32 {
	return IDMouseInteraction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       MouseInteraction obj = new MouseInteraction();
//       byte nullBits = buf.getByte(offset);
//       obj.clientTimestamp = buf.getLongLE(offset + 1);
//       obj.activeSlot = buf.getIntLE(offset + 9);
//       if ((nullBits & 2) != 0) {
//          obj.screenPoint = Vector2f.deserialize(buf, offset + 13);
//       }
//
//       if ((nullBits & 4) != 0) {
//          obj.mouseButton = MouseButtonEvent.deserialize(buf, offset + 21);
//       }
//
//       if ((nullBits & 16) != 0) {
//          obj.worldInteraction = WorldInteraction.deserialize(buf, offset + 24);
//       }
//
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 52 + buf.getIntLE(offset + 44);
//          int itemInHandIdLen = VarInt.peek(buf, varPos0);
//          if (itemInHandIdLen < 0) {
//             throw ProtocolException.negativeLength("ItemInHandId", itemInHandIdLen);
//          }
//
//          if (itemInHandIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("ItemInHandId", itemInHandIdLen, 4096000);
//          }
//
//          obj.itemInHandId = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 8) != 0) {
//          int varPos1 = offset + 52 + buf.getIntLE(offset + 48);
//          obj.mouseMotion = MouseMotionEvent.deserialize(buf, varPos1);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.itemInHandId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.screenPoint != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.mouseButton != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       if (this.mouseMotion != null) {
//          nullBits = (byte)(nullBits | 8);
//       }
//
//       if (this.worldInteraction != null) {
//          nullBits = (byte)(nullBits | 16);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeLongLE(this.clientTimestamp);
//       buf.writeIntLE(this.activeSlot);
//       if (this.screenPoint != null) {
//          this.screenPoint.serialize(buf);
//       } else {
//          buf.writeZero(8);
//       }
//
//       if (this.mouseButton != null) {
//          this.mouseButton.serialize(buf);
//       } else {
//          buf.writeZero(3);
//       }
//
//       if (this.worldInteraction != null) {
//          this.worldInteraction.serialize(buf);
//       } else {
//          buf.writeZero(20);
//       }
//
//       int itemInHandIdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int mouseMotionOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.itemInHandId != null) {
//          buf.setIntLE(itemInHandIdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.itemInHandId, 4096000);
//       } else {
//          buf.setIntLE(itemInHandIdOffsetSlot, -1);
//       }
//
//       if (this.mouseMotion != null) {
//          buf.setIntLE(mouseMotionOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.mouseMotion.serialize(buf);
//       } else {
//          buf.setIntLE(mouseMotionOffsetSlot, -1);
//       }
//    }

func (pk *MouseInteraction) Marshal(io protocol.IO) {
	// TODO: Implement
}
