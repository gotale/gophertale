package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PlayInteractionFor struct{}

func (*PlayInteractionFor) ID() uint32 {
	return IDPlayInteractionFor
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PlayInteractionFor obj = new PlayInteractionFor();
//       byte nullBits = buf.getByte(offset);
//       obj.entityId = buf.getIntLE(offset + 1);
//       obj.chainId = buf.getIntLE(offset + 5);
//       obj.operationIndex = buf.getIntLE(offset + 9);
//       obj.interactionId = buf.getIntLE(offset + 13);
//       obj.interactionType = InteractionType.fromValue(buf.getByte(offset + 17));
//       obj.cancel = buf.getByte(offset + 18) != 0;
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 27 + buf.getIntLE(offset + 19);
//          obj.forkedId = ForkedChainId.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 27 + buf.getIntLE(offset + 23);
//          int interactedItemIdLen = VarInt.peek(buf, varPos1);
//          if (interactedItemIdLen < 0) {
//             throw ProtocolException.negativeLength("InteractedItemId", interactedItemIdLen);
//          }
//
//          if (interactedItemIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("InteractedItemId", interactedItemIdLen, 4096000);
//          }
//
//          obj.interactedItemId = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.forkedId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.interactedItemId != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.entityId);
//       buf.writeIntLE(this.chainId);
//       buf.writeIntLE(this.operationIndex);
//       buf.writeIntLE(this.interactionId);
//       buf.writeByte(this.interactionType.getValue());
//       buf.writeByte(this.cancel ? 1 : 0);
//       int forkedIdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int interactedItemIdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.forkedId != null) {
//          buf.setIntLE(forkedIdOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.forkedId.serialize(buf);
//       } else {
//          buf.setIntLE(forkedIdOffsetSlot, -1);
//       }
//
//       if (this.interactedItemId != null) {
//          buf.setIntLE(interactedItemIdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.interactedItemId, 4096000);
//       } else {
//          buf.setIntLE(interactedItemIdOffsetSlot, -1);
//       }
//    }

func (pk *PlayInteractionFor) Marshal(io protocol.IO) {
	// TODO: Implement
}
