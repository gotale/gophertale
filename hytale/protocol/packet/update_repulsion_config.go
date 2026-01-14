package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateRepulsionConfig struct{}

func (*UpdateRepulsionConfig) ID() uint32 {
	return IDUpdateRepulsionConfig
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateRepulsionConfig obj = new UpdateRepulsionConfig();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int repulsionConfigsCount = VarInt.peek(buf, pos);
//          if (repulsionConfigsCount < 0) {
//             throw ProtocolException.negativeLength("RepulsionConfigs", repulsionConfigsCount);
//          }
//
//          if (repulsionConfigsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("RepulsionConfigs", repulsionConfigsCount, 4096000);
//          }
//
//          pos += VarInt.size(repulsionConfigsCount);
//          obj.repulsionConfigs = new HashMap<>(repulsionConfigsCount);
//
//          for (int i = 0; i < repulsionConfigsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             RepulsionConfig val = RepulsionConfig.deserialize(buf, pos);
//             pos += RepulsionConfig.computeBytesConsumed(buf, pos);
//             if (obj.repulsionConfigs.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("repulsionConfigs", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.repulsionConfigs != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.repulsionConfigs != null) {
//          if (this.repulsionConfigs.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("RepulsionConfigs", this.repulsionConfigs.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.repulsionConfigs.size());
//
//          for (Entry<Integer, RepulsionConfig> e : this.repulsionConfigs.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateRepulsionConfig) Marshal(io protocol.IO) {
	// TODO: Implement
}
