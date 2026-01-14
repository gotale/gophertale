package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateResourceTypes struct{}

func (*UpdateResourceTypes) ID() uint32 {
	return IDUpdateResourceTypes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateResourceTypes obj = new UpdateResourceTypes();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int resourceTypesCount = VarInt.peek(buf, pos);
//          if (resourceTypesCount < 0) {
//             throw ProtocolException.negativeLength("ResourceTypes", resourceTypesCount);
//          }
//
//          if (resourceTypesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ResourceTypes", resourceTypesCount, 4096000);
//          }
//
//          pos += VarInt.size(resourceTypesCount);
//          obj.resourceTypes = new HashMap<>(resourceTypesCount);
//
//          for (int i = 0; i < resourceTypesCount; i++) {
//             int keyLen = VarInt.peek(buf, pos);
//             if (keyLen < 0) {
//                throw ProtocolException.negativeLength("key", keyLen);
//             }
//
//             if (keyLen > 4096000) {
//                throw ProtocolException.stringTooLong("key", keyLen, 4096000);
//             }
//
//             int keyVarLen = VarInt.length(buf, pos);
//             String key = PacketIO.readVarString(buf, pos);
//             pos += keyVarLen + keyLen;
//             ResourceType val = ResourceType.deserialize(buf, pos);
//             pos += ResourceType.computeBytesConsumed(buf, pos);
//             if (obj.resourceTypes.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("resourceTypes", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.resourceTypes != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.resourceTypes != null) {
//          if (this.resourceTypes.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ResourceTypes", this.resourceTypes.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.resourceTypes.size());
//
//          for (Entry<String, ResourceType> e : this.resourceTypes.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateResourceTypes) Marshal(io protocol.IO) {
	// TODO: Implement
}
