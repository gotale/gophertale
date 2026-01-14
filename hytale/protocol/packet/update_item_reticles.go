package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateItemReticles struct{}

func (*UpdateItemReticles) ID() uint32 {
	return IDUpdateItemReticles
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateItemReticles obj = new UpdateItemReticles();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int itemReticleConfigsCount = VarInt.peek(buf, pos);
//          if (itemReticleConfigsCount < 0) {
//             throw ProtocolException.negativeLength("ItemReticleConfigs", itemReticleConfigsCount);
//          }
//
//          if (itemReticleConfigsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemReticleConfigs", itemReticleConfigsCount, 4096000);
//          }
//
//          pos += VarInt.size(itemReticleConfigsCount);
//          obj.itemReticleConfigs = new HashMap<>(itemReticleConfigsCount);
//
//          for (int i = 0; i < itemReticleConfigsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             ItemReticleConfig val = ItemReticleConfig.deserialize(buf, pos);
//             pos += ItemReticleConfig.computeBytesConsumed(buf, pos);
//             if (obj.itemReticleConfigs.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("itemReticleConfigs", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.itemReticleConfigs != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.itemReticleConfigs != null) {
//          if (this.itemReticleConfigs.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ItemReticleConfigs", this.itemReticleConfigs.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.itemReticleConfigs.size());
//
//          for (Entry<Integer, ItemReticleConfig> e : this.itemReticleConfigs.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateItemReticles) Marshal(io protocol.IO) {
	// TODO: Implement
}
