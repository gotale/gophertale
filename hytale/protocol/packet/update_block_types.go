package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockTypes struct{}

func (*UpdateBlockTypes) ID() uint32 {
	return IDUpdateBlockTypes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockTypes obj = new UpdateBlockTypes();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       obj.updateBlockTextures = buf.getByte(offset + 6) != 0;
//       obj.updateModelTextures = buf.getByte(offset + 7) != 0;
//       obj.updateModels = buf.getByte(offset + 8) != 0;
//       obj.updateMapGeometry = buf.getByte(offset + 9) != 0;
//       int pos = offset + 10;
//       if ((nullBits & 1) != 0) {
//          int blockTypesCount = VarInt.peek(buf, pos);
//          if (blockTypesCount < 0) {
//             throw ProtocolException.negativeLength("BlockTypes", blockTypesCount);
//          }
//
//          if (blockTypesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockTypes", blockTypesCount, 4096000);
//          }
//
//          pos += VarInt.size(blockTypesCount);
//          obj.blockTypes = new HashMap<>(blockTypesCount);
//
//          for (int i = 0; i < blockTypesCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             BlockType val = BlockType.deserialize(buf, pos);
//             pos += BlockType.computeBytesConsumed(buf, pos);
//             if (obj.blockTypes.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("blockTypes", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.blockTypes != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       buf.writeByte(this.updateBlockTextures ? 1 : 0);
//       buf.writeByte(this.updateModelTextures ? 1 : 0);
//       buf.writeByte(this.updateModels ? 1 : 0);
//       buf.writeByte(this.updateMapGeometry ? 1 : 0);
//       if (this.blockTypes != null) {
//          if (this.blockTypes.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockTypes", this.blockTypes.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.blockTypes.size());
//
//          for (Entry<Integer, BlockType> e : this.blockTypes.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateBlockTypes) Marshal(io protocol.IO) {
	// TODO: Implement
}
