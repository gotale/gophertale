package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateSoundEvents struct{}

func (*UpdateSoundEvents) ID() uint32 {
	return IDUpdateSoundEvents
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateSoundEvents obj = new UpdateSoundEvents();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int soundEventsCount = VarInt.peek(buf, pos);
//          if (soundEventsCount < 0) {
//             throw ProtocolException.negativeLength("SoundEvents", soundEventsCount);
//          }
//
//          if (soundEventsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("SoundEvents", soundEventsCount, 4096000);
//          }
//
//          pos += VarInt.size(soundEventsCount);
//          obj.soundEvents = new HashMap<>(soundEventsCount);
//
//          for (int i = 0; i < soundEventsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             SoundEvent val = SoundEvent.deserialize(buf, pos);
//             pos += SoundEvent.computeBytesConsumed(buf, pos);
//             if (obj.soundEvents.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("soundEvents", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.soundEvents != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.soundEvents != null) {
//          if (this.soundEvents.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("SoundEvents", this.soundEvents.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.soundEvents.size());
//
//          for (Entry<Integer, SoundEvent> e : this.soundEvents.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateSoundEvents) Marshal(io protocol.IO) {
	// TODO: Implement
}
