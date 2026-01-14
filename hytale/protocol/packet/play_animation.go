package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PlayAnimation struct{}

func (*PlayAnimation) ID() uint32 {
	return IDPlayAnimation
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PlayAnimation obj = new PlayAnimation();
//       byte nullBits = buf.getByte(offset);
//       obj.entityId = buf.getIntLE(offset + 1);
//       obj.slot = AnimationSlot.fromValue(buf.getByte(offset + 5));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 14 + buf.getIntLE(offset + 6);
//          int itemAnimationsIdLen = VarInt.peek(buf, varPos0);
//          if (itemAnimationsIdLen < 0) {
//             throw ProtocolException.negativeLength("ItemAnimationsId", itemAnimationsIdLen);
//          }
//
//          if (itemAnimationsIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("ItemAnimationsId", itemAnimationsIdLen, 4096000);
//          }
//
//          obj.itemAnimationsId = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 14 + buf.getIntLE(offset + 10);
//          int animationIdLen = VarInt.peek(buf, varPos1);
//          if (animationIdLen < 0) {
//             throw ProtocolException.negativeLength("AnimationId", animationIdLen);
//          }
//
//          if (animationIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("AnimationId", animationIdLen, 4096000);
//          }
//
//          obj.animationId = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.itemAnimationsId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.animationId != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.entityId);
//       buf.writeByte(this.slot.getValue());
//       int itemAnimationsIdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int animationIdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.itemAnimationsId != null) {
//          buf.setIntLE(itemAnimationsIdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.itemAnimationsId, 4096000);
//       } else {
//          buf.setIntLE(itemAnimationsIdOffsetSlot, -1);
//       }
//
//       if (this.animationId != null) {
//          buf.setIntLE(animationIdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.animationId, 4096000);
//       } else {
//          buf.setIntLE(animationIdOffsetSlot, -1);
//       }
//    }

func (pk *PlayAnimation) Marshal(io protocol.IO) {
	// TODO: Implement
}
