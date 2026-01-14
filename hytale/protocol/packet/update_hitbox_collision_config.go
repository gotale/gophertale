package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateHitboxCollisionConfig struct{}

func (*UpdateHitboxCollisionConfig) ID() uint32 {
	return IDUpdateHitboxCollisionConfig
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateHitboxCollisionConfig obj = new UpdateHitboxCollisionConfig();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int hitboxCollisionConfigsCount = VarInt.peek(buf, pos);
//          if (hitboxCollisionConfigsCount < 0) {
//             throw ProtocolException.negativeLength("HitboxCollisionConfigs", hitboxCollisionConfigsCount);
//          }
//
//          if (hitboxCollisionConfigsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("HitboxCollisionConfigs", hitboxCollisionConfigsCount, 4096000);
//          }
//
//          pos += VarInt.size(hitboxCollisionConfigsCount);
//          obj.hitboxCollisionConfigs = new HashMap<>(hitboxCollisionConfigsCount);
//
//          for (int i = 0; i < hitboxCollisionConfigsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             HitboxCollisionConfig val = HitboxCollisionConfig.deserialize(buf, pos);
//             pos += HitboxCollisionConfig.computeBytesConsumed(buf, pos);
//             if (obj.hitboxCollisionConfigs.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("hitboxCollisionConfigs", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.hitboxCollisionConfigs != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.hitboxCollisionConfigs != null) {
//          if (this.hitboxCollisionConfigs.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("HitboxCollisionConfigs", this.hitboxCollisionConfigs.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.hitboxCollisionConfigs.size());
//
//          for (Entry<Integer, HitboxCollisionConfig> e : this.hitboxCollisionConfigs.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateHitboxCollisionConfig) Marshal(io protocol.IO) {
	// TODO: Implement
}
