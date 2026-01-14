package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEqualizerEffects struct{}

func (*UpdateEqualizerEffects) ID() uint32 {
	return IDUpdateEqualizerEffects
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEqualizerEffects obj = new UpdateEqualizerEffects();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int effectsCount = VarInt.peek(buf, pos);
//          if (effectsCount < 0) {
//             throw ProtocolException.negativeLength("Effects", effectsCount);
//          }
//
//          if (effectsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Effects", effectsCount, 4096000);
//          }
//
//          pos += VarInt.size(effectsCount);
//          obj.effects = new HashMap<>(effectsCount);
//
//          for (int i = 0; i < effectsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             EqualizerEffect val = EqualizerEffect.deserialize(buf, pos);
//             pos += EqualizerEffect.computeBytesConsumed(buf, pos);
//             if (obj.effects.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("effects", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.effects != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.effects != null) {
//          if (this.effects.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Effects", this.effects.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.effects.size());
//
//          for (Entry<Integer, EqualizerEffect> e : this.effects.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateEqualizerEffects) Marshal(io protocol.IO) {
	// TODO: Implement
}
