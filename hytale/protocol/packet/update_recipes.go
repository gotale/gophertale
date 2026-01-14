package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateRecipes struct{}

func (*UpdateRecipes) ID() uint32 {
	return IDUpdateRecipes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateRecipes obj = new UpdateRecipes();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 10 + buf.getIntLE(offset + 2);
//          int recipesCount = VarInt.peek(buf, varPos0);
//          if (recipesCount < 0) {
//             throw ProtocolException.negativeLength("Recipes", recipesCount);
//          }
//
//          if (recipesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Recipes", recipesCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          obj.recipes = new HashMap<>(recipesCount);
//          int dictPos = varPos0 + varIntLen;
//
//          for (int i = 0; i < recipesCount; i++) {
//             int keyLen = VarInt.peek(buf, dictPos);
//             if (keyLen < 0) {
//                throw ProtocolException.negativeLength("key", keyLen);
//             }
//
//             if (keyLen > 4096000) {
//                throw ProtocolException.stringTooLong("key", keyLen, 4096000);
//             }
//
//             int keyVarLen = VarInt.length(buf, dictPos);
//             String key = PacketIO.readVarString(buf, dictPos);
//             dictPos += keyVarLen + keyLen;
//             CraftingRecipe val = CraftingRecipe.deserialize(buf, dictPos);
//             dictPos += CraftingRecipe.computeBytesConsumed(buf, dictPos);
//             if (obj.recipes.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("recipes", key);
//             }
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 10 + buf.getIntLE(offset + 6);
//          int removedRecipesCount = VarInt.peek(buf, varPos1);
//          if (removedRecipesCount < 0) {
//             throw ProtocolException.negativeLength("RemovedRecipes", removedRecipesCount);
//          }
//
//          if (removedRecipesCount > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedRecipes", removedRecipesCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + removedRecipesCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("RemovedRecipes", varPos1 + varIntLen + removedRecipesCount * 1, buf.readableBytes());
//          }
//
//          obj.removedRecipes = new String[removedRecipesCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < removedRecipesCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("removedRecipes[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("removedRecipes[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.removedRecipes[i] = PacketIO.readVarString(buf, elemPos);
//             elemPos += strVarLen + strLen;
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.recipes != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.removedRecipes != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       int recipesOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int removedRecipesOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.recipes != null) {
//          buf.setIntLE(recipesOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.recipes.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Recipes", this.recipes.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.recipes.size());
//
//          for (Entry<String, CraftingRecipe> e : this.recipes.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       } else {
//          buf.setIntLE(recipesOffsetSlot, -1);
//       }
//
//       if (this.removedRecipes != null) {
//          buf.setIntLE(removedRecipesOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.removedRecipes.length > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedRecipes", this.removedRecipes.length, 4096000);
//          }
//
//          VarInt.write(buf, this.removedRecipes.length);
//
//          for (String item : this.removedRecipes) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(removedRecipesOffsetSlot, -1);
//       }
//    }

func (pk *UpdateRecipes) Marshal(io protocol.IO) {
	// TODO: Implement
}
