package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateSoundSets struct{}

func (*UpdateSoundSets) ID() uint32 {
	return IDUpdateSoundSets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateSoundSets obj = new UpdateSoundSets();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int soundSetsCount = VarInt.peek(buf, pos);
//          if (soundSetsCount < 0) {
//             throw ProtocolException.negativeLength("SoundSets", soundSetsCount);
//          }
//
//          if (soundSetsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("SoundSets", soundSetsCount, 4096000);
//          }
//
//          pos += VarInt.size(soundSetsCount);
//          obj.soundSets = new HashMap<>(soundSetsCount);
//
//          for (int i = 0; i < soundSetsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             SoundSet val = SoundSet.deserialize(buf, pos);
//             pos += SoundSet.computeBytesConsumed(buf, pos);
//             if (obj.soundSets.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("soundSets", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.soundSets != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.soundSets != null) {
//          if (this.soundSets.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("SoundSets", this.soundSets.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.soundSets.size());
//
//          for (Entry<Integer, SoundSet> e : this.soundSets.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateSoundSets) Marshal(io protocol.IO) {
	// TODO: Implement
}
