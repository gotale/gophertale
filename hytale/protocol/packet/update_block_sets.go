package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockSets struct{}

func (*UpdateBlockSets) ID() uint32 {
	return IDUpdateBlockSets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockSets obj = new UpdateBlockSets();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int blockSetsCount = VarInt.peek(buf, pos);
//          if (blockSetsCount < 0) {
//             throw ProtocolException.negativeLength("BlockSets", blockSetsCount);
//          }
//
//          if (blockSetsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockSets", blockSetsCount, 4096000);
//          }
//
//          pos += VarInt.size(blockSetsCount);
//          obj.blockSets = new HashMap<>(blockSetsCount);
//
//          for (int i = 0; i < blockSetsCount; i++) {
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
//             BlockSet val = BlockSet.deserialize(buf, pos);
//             pos += BlockSet.computeBytesConsumed(buf, pos);
//             if (obj.blockSets.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("blockSets", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.blockSets != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.blockSets != null) {
//          if (this.blockSets.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockSets", this.blockSets.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.blockSets.size());
//
//          for (Entry<String, BlockSet> e : this.blockSets.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateBlockSets) Marshal(io protocol.IO) {
	// TODO: Implement
}
