package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateCameraShake struct{}

func (*UpdateCameraShake) ID() uint32 {
	return IDUpdateCameraShake
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateCameraShake obj = new UpdateCameraShake();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int profilesCount = VarInt.peek(buf, pos);
//          if (profilesCount < 0) {
//             throw ProtocolException.negativeLength("Profiles", profilesCount);
//          }
//
//          if (profilesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Profiles", profilesCount, 4096000);
//          }
//
//          pos += VarInt.size(profilesCount);
//          obj.profiles = new HashMap<>(profilesCount);
//
//          for (int i = 0; i < profilesCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             CameraShake val = CameraShake.deserialize(buf, pos);
//             pos += CameraShake.computeBytesConsumed(buf, pos);
//             if (obj.profiles.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("profiles", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.profiles != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.profiles != null) {
//          if (this.profiles.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Profiles", this.profiles.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.profiles.size());
//
//          for (Entry<Integer, CameraShake> e : this.profiles.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateCameraShake) Marshal(io protocol.IO) {
	// TODO: Implement
}
