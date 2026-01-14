package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEnvironments struct{}

func (*UpdateEnvironments) ID() uint32 {
	return IDUpdateEnvironments
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEnvironments obj = new UpdateEnvironments();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       obj.rebuildMapGeometry = buf.getByte(offset + 6) != 0;
//       int pos = offset + 7;
//       if ((nullBits & 1) != 0) {
//          int environmentsCount = VarInt.peek(buf, pos);
//          if (environmentsCount < 0) {
//             throw ProtocolException.negativeLength("Environments", environmentsCount);
//          }
//
//          if (environmentsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Environments", environmentsCount, 4096000);
//          }
//
//          pos += VarInt.size(environmentsCount);
//          obj.environments = new HashMap<>(environmentsCount);
//
//          for (int i = 0; i < environmentsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             WorldEnvironment val = WorldEnvironment.deserialize(buf, pos);
//             pos += WorldEnvironment.computeBytesConsumed(buf, pos);
//             if (obj.environments.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("environments", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.environments != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       buf.writeByte(this.rebuildMapGeometry ? 1 : 0);
//       if (this.environments != null) {
//          if (this.environments.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Environments", this.environments.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.environments.size());
//
//          for (Entry<Integer, WorldEnvironment> e : this.environments.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateEnvironments) Marshal(io protocol.IO) {
	// TODO: Implement
}
