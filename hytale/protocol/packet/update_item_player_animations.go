package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateItemPlayerAnimations struct{}

func (*UpdateItemPlayerAnimations) ID() uint32 {
	return IDUpdateItemPlayerAnimations
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateItemPlayerAnimations obj = new UpdateItemPlayerAnimations();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int itemPlayerAnimationsCount = VarInt.peek(buf, pos);
//          if (itemPlayerAnimationsCount < 0) {
//             throw ProtocolException.negativeLength("ItemPlayerAnimations", itemPlayerAnimationsCount);
//          }
//
//          if (itemPlayerAnimationsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemPlayerAnimations", itemPlayerAnimationsCount, 4096000);
//          }
//
//          pos += VarInt.size(itemPlayerAnimationsCount);
//          obj.itemPlayerAnimations = new HashMap<>(itemPlayerAnimationsCount);
//
//          for (int i = 0; i < itemPlayerAnimationsCount; i++) {
//             int keyLen = VarInt.peek(buf, pos);
//             if (keyLen < 0) {
//                throw ProtocolException.negativeLength("key", keyLen);
//             }
//
//             if (keyLen > 4096000) {
//                throw ProtocolException.stringTooLong("key", keyLen, 4096000);
//             }
//
//             int keyVarLen = VarInt.length(buf, pos);
//             String key = PacketIO.readVarString(buf, pos);
//             pos += keyVarLen + keyLen;
//             ItemPlayerAnimations val = ItemPlayerAnimations.deserialize(buf, pos);
//             pos += ItemPlayerAnimations.computeBytesConsumed(buf, pos);
//             if (obj.itemPlayerAnimations.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("itemPlayerAnimations", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.itemPlayerAnimations != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.itemPlayerAnimations != null) {
//          if (this.itemPlayerAnimations.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemPlayerAnimations", this.itemPlayerAnimations.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.itemPlayerAnimations.size());
//
//          for (Entry<String, ItemPlayerAnimations> e : this.itemPlayerAnimations.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateItemPlayerAnimations) Marshal(io protocol.IO) {
	// TODO: Implement
}
