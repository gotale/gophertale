package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockBreakingDecals struct{}

func (*UpdateBlockBreakingDecals) ID() uint32 {
	return IDUpdateBlockBreakingDecals
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockBreakingDecals obj = new UpdateBlockBreakingDecals();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int blockBreakingDecalsCount = VarInt.peek(buf, pos);
//          if (blockBreakingDecalsCount < 0) {
//             throw ProtocolException.negativeLength("BlockBreakingDecals", blockBreakingDecalsCount);
//          }
//
//          if (blockBreakingDecalsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockBreakingDecals", blockBreakingDecalsCount, 4096000);
//          }
//
//          pos += VarInt.size(blockBreakingDecalsCount);
//          obj.blockBreakingDecals = new HashMap<>(blockBreakingDecalsCount);
//
//          for (int i = 0; i < blockBreakingDecalsCount; i++) {
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
//             BlockBreakingDecal val = BlockBreakingDecal.deserialize(buf, pos);
//             pos += BlockBreakingDecal.computeBytesConsumed(buf, pos);
//             if (obj.blockBreakingDecals.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("blockBreakingDecals", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.blockBreakingDecals != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.blockBreakingDecals != null) {
//          if (this.blockBreakingDecals.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockBreakingDecals", this.blockBreakingDecals.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.blockBreakingDecals.size());
//
//          for (Entry<String, BlockBreakingDecal> e : this.blockBreakingDecals.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateBlockBreakingDecals) Marshal(io protocol.IO) {
	// TODO: Implement
}
