package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEntityEffects struct{}

func (*UpdateEntityEffects) ID() uint32 {
	return IDUpdateEntityEffects
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEntityEffects obj = new UpdateEntityEffects();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int entityEffectsCount = VarInt.peek(buf, pos);
//          if (entityEffectsCount < 0) {
//             throw ProtocolException.negativeLength("EntityEffects", entityEffectsCount);
//          }
//
//          if (entityEffectsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("EntityEffects", entityEffectsCount, 4096000);
//          }
//
//          pos += VarInt.size(entityEffectsCount);
//          obj.entityEffects = new HashMap<>(entityEffectsCount);
//
//          for (int i = 0; i < entityEffectsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             EntityEffect val = EntityEffect.deserialize(buf, pos);
//             pos += EntityEffect.computeBytesConsumed(buf, pos);
//             if (obj.entityEffects.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("entityEffects", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.entityEffects != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.entityEffects != null) {
//          if (this.entityEffects.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("EntityEffects", this.entityEffects.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.entityEffects.size());
//
//          for (Entry<Integer, EntityEffect> e : this.entityEffects.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateEntityEffects) Marshal(io protocol.IO) {
	// TODO: Implement
}
