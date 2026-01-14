package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateAudioCategories struct{}

func (*UpdateAudioCategories) ID() uint32 {
	return IDUpdateAudioCategories
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateAudioCategories obj = new UpdateAudioCategories();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int categoriesCount = VarInt.peek(buf, pos);
//          if (categoriesCount < 0) {
//             throw ProtocolException.negativeLength("Categories", categoriesCount);
//          }
//
//          if (categoriesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Categories", categoriesCount, 4096000);
//          }
//
//          pos += VarInt.size(categoriesCount);
//          obj.categories = new HashMap<>(categoriesCount);
//
//          for (int i = 0; i < categoriesCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             AudioCategory val = AudioCategory.deserialize(buf, pos);
//             pos += AudioCategory.computeBytesConsumed(buf, pos);
//             if (obj.categories.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("categories", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.categories != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.categories != null) {
//          if (this.categories.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Categories", this.categories.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.categories.size());
//
//          for (Entry<Integer, AudioCategory> e : this.categories.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateAudioCategories) Marshal(io protocol.IO) {
	// TODO: Implement
}
