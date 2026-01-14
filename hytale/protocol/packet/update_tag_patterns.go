package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateTagPatterns struct{}

func (*UpdateTagPatterns) ID() uint32 {
	return IDUpdateTagPatterns
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateTagPatterns obj = new UpdateTagPatterns();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int patternsCount = VarInt.peek(buf, pos);
//          if (patternsCount < 0) {
//             throw ProtocolException.negativeLength("Patterns", patternsCount);
//          }
//
//          if (patternsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Patterns", patternsCount, 4096000);
//          }
//
//          pos += VarInt.size(patternsCount);
//          obj.patterns = new HashMap<>(patternsCount);
//
//          for (int i = 0; i < patternsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             TagPattern val = TagPattern.deserialize(buf, pos);
//             pos += TagPattern.computeBytesConsumed(buf, pos);
//             if (obj.patterns.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("patterns", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.patterns != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.patterns != null) {
//          if (this.patterns.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Patterns", this.patterns.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.patterns.size());
//
//          for (Entry<Integer, TagPattern> e : this.patterns.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateTagPatterns) Marshal(io protocol.IO) {
	// TODO: Implement
}
