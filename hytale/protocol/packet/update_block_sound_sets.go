package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockSoundSets struct{}

func (*UpdateBlockSoundSets) ID() uint32 {
	return IDUpdateBlockSoundSets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockSoundSets obj = new UpdateBlockSoundSets();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int blockSoundSetsCount = VarInt.peek(buf, pos);
//          if (blockSoundSetsCount < 0) {
//             throw ProtocolException.negativeLength("BlockSoundSets", blockSoundSetsCount);
//          }
//
//          if (blockSoundSetsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockSoundSets", blockSoundSetsCount, 4096000);
//          }
//
//          pos += VarInt.size(blockSoundSetsCount);
//          obj.blockSoundSets = new HashMap<>(blockSoundSetsCount);
//
//          for (int i = 0; i < blockSoundSetsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             BlockSoundSet val = BlockSoundSet.deserialize(buf, pos);
//             pos += BlockSoundSet.computeBytesConsumed(buf, pos);
//             if (obj.blockSoundSets.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("blockSoundSets", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.blockSoundSets != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.blockSoundSets != null) {
//          if (this.blockSoundSets.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockSoundSets", this.blockSoundSets.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.blockSoundSets.size());
//
//          for (Entry<Integer, BlockSoundSet> e : this.blockSoundSets.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateBlockSoundSets) Marshal(io protocol.IO) {
	// TODO: Implement
}
