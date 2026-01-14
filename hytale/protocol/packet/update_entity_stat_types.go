package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEntityStatTypes struct{}

func (*UpdateEntityStatTypes) ID() uint32 {
	return IDUpdateEntityStatTypes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEntityStatTypes obj = new UpdateEntityStatTypes();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int typesCount = VarInt.peek(buf, pos);
//          if (typesCount < 0) {
//             throw ProtocolException.negativeLength("Types", typesCount);
//          }
//
//          if (typesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Types", typesCount, 4096000);
//          }
//
//          pos += VarInt.size(typesCount);
//          obj.types = new HashMap<>(typesCount);
//
//          for (int i = 0; i < typesCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             EntityStatType val = EntityStatType.deserialize(buf, pos);
//             pos += EntityStatType.computeBytesConsumed(buf, pos);
//             if (obj.types.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("types", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.types != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.types != null) {
//          if (this.types.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Types", this.types.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.types.size());
//
//          for (Entry<Integer, EntityStatType> e : this.types.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateEntityStatTypes) Marshal(io protocol.IO) {
	// TODO: Implement
}
