package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerTags struct{}

func (*ServerTags) ID() uint32 {
	return IDServerTags
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerTags obj = new ServerTags();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int tagsCount = VarInt.peek(buf, pos);
//          if (tagsCount < 0) {
//             throw ProtocolException.negativeLength("Tags", tagsCount);
//          }
//
//          if (tagsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Tags", tagsCount, 4096000);
//          }
//
//          pos += VarInt.size(tagsCount);
//          obj.tags = new HashMap<>(tagsCount);
//
//          for (int i = 0; i < tagsCount; i++) {
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
//             int val = buf.getIntLE(pos);
//             pos += 4;
//             if (obj.tags.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("tags", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.tags != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.tags != null) {
//          if (this.tags.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Tags", this.tags.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.tags.size());
//
//          for (Entry<String, Integer> e : this.tags.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             buf.writeIntLE(e.getValue());
//          }
//       }
//    }

func (pk *ServerTags) Marshal(io protocol.IO) {
	// TODO: Implement
}
