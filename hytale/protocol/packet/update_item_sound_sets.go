package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateItemSoundSets struct{}

func (*UpdateItemSoundSets) ID() uint32 {
	return IDUpdateItemSoundSets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateItemSoundSets obj = new UpdateItemSoundSets();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int itemSoundSetsCount = VarInt.peek(buf, pos);
//          if (itemSoundSetsCount < 0) {
//             throw ProtocolException.negativeLength("ItemSoundSets", itemSoundSetsCount);
//          }
//
//          if (itemSoundSetsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemSoundSets", itemSoundSetsCount, 4096000);
//          }
//
//          pos += VarInt.size(itemSoundSetsCount);
//          obj.itemSoundSets = new HashMap<>(itemSoundSetsCount);
//
//          for (int i = 0; i < itemSoundSetsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             ItemSoundSet val = ItemSoundSet.deserialize(buf, pos);
//             pos += ItemSoundSet.computeBytesConsumed(buf, pos);
//             if (obj.itemSoundSets.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("itemSoundSets", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.itemSoundSets != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.itemSoundSets != null) {
//          if (this.itemSoundSets.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemSoundSets", this.itemSoundSets.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.itemSoundSets.size());
//
//          for (Entry<Integer, ItemSoundSet> e : this.itemSoundSets.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateItemSoundSets) Marshal(io protocol.IO) {
	// TODO: Implement
}
