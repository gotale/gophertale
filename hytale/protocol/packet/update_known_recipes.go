package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateKnownRecipes struct{}

func (*UpdateKnownRecipes) ID() uint32 {
	return IDUpdateKnownRecipes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateKnownRecipes obj = new UpdateKnownRecipes();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int knownCount = VarInt.peek(buf, pos);
//          if (knownCount < 0) {
//             throw ProtocolException.negativeLength("Known", knownCount);
//          }
//
//          if (knownCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Known", knownCount, 4096000);
//          }
//
//          pos += VarInt.size(knownCount);
//          obj.known = new HashMap<>(knownCount);
//
//          for (int i = 0; i < knownCount; i++) {
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
//             CraftingRecipe val = CraftingRecipe.deserialize(buf, pos);
//             pos += CraftingRecipe.computeBytesConsumed(buf, pos);
//             if (obj.known.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("known", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.known != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.known != null) {
//          if (this.known.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Known", this.known.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.known.size());
//
//          for (Entry<String, CraftingRecipe> e : this.known.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateKnownRecipes) Marshal(io protocol.IO) {
	// TODO: Implement
}
