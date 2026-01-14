package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateItemQualities struct{}

func (*UpdateItemQualities) ID() uint32 {
	return IDUpdateItemQualities
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateItemQualities obj = new UpdateItemQualities();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int itemQualitiesCount = VarInt.peek(buf, pos);
//          if (itemQualitiesCount < 0) {
//             throw ProtocolException.negativeLength("ItemQualities", itemQualitiesCount);
//          }
//
//          if (itemQualitiesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemQualities", itemQualitiesCount, 4096000);
//          }
//
//          pos += VarInt.size(itemQualitiesCount);
//          obj.itemQualities = new HashMap<>(itemQualitiesCount);
//
//          for (int i = 0; i < itemQualitiesCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             ItemQuality val = ItemQuality.deserialize(buf, pos);
//             pos += ItemQuality.computeBytesConsumed(buf, pos);
//             if (obj.itemQualities.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("itemQualities", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.itemQualities != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.itemQualities != null) {
//          if (this.itemQualities.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemQualities", this.itemQualities.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.itemQualities.size());
//
//          for (Entry<Integer, ItemQuality> e : this.itemQualities.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateItemQualities) Marshal(io protocol.IO) {
	// TODO: Implement
}
