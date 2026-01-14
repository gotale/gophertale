package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockGroups struct{}

func (*UpdateBlockGroups) ID() uint32 {
	return IDUpdateBlockGroups
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockGroups obj = new UpdateBlockGroups();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int groupsCount = VarInt.peek(buf, pos);
//          if (groupsCount < 0) {
//             throw ProtocolException.negativeLength("Groups", groupsCount);
//          }
//
//          if (groupsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Groups", groupsCount, 4096000);
//          }
//
//          pos += VarInt.size(groupsCount);
//          obj.groups = new HashMap<>(groupsCount);
//
//          for (int i = 0; i < groupsCount; i++) {
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
//             BlockGroup val = BlockGroup.deserialize(buf, pos);
//             pos += BlockGroup.computeBytesConsumed(buf, pos);
//             if (obj.groups.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("groups", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.groups != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.groups != null) {
//          if (this.groups.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Groups", this.groups.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.groups.size());
//
//          for (Entry<String, BlockGroup> e : this.groups.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateBlockGroups) Marshal(io protocol.IO) {
	// TODO: Implement
}
