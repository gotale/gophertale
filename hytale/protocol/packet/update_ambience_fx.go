package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateAmbienceFX struct{}

func (*UpdateAmbienceFX) ID() uint32 {
	return IDUpdateAmbienceFX
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateAmbienceFX obj = new UpdateAmbienceFX();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int ambienceFXCount = VarInt.peek(buf, pos);
//          if (ambienceFXCount < 0) {
//             throw ProtocolException.negativeLength("AmbienceFX", ambienceFXCount);
//          }
//
//          if (ambienceFXCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("AmbienceFX", ambienceFXCount, 4096000);
//          }
//
//          pos += VarInt.size(ambienceFXCount);
//          obj.ambienceFX = new HashMap<>(ambienceFXCount);
//
//          for (int i = 0; i < ambienceFXCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             AmbienceFX val = AmbienceFX.deserialize(buf, pos);
//             pos += AmbienceFX.computeBytesConsumed(buf, pos);
//             if (obj.ambienceFX.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("ambienceFX", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.ambienceFX != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.ambienceFX != null) {
//          if (this.ambienceFX.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("AmbienceFX", this.ambienceFX.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.ambienceFX.size());
//
//          for (Entry<Integer, AmbienceFX> e : this.ambienceFX.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateAmbienceFX) Marshal(io protocol.IO) {
	// TODO: Implement
}
