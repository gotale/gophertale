package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorAssetListSetup struct{}

func (*AssetEditorAssetListSetup) ID() uint32 {
	return IDAssetEditorAssetListSetup
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorAssetListSetup obj = new AssetEditorAssetListSetup();
//       byte nullBits = buf.getByte(offset);
//       obj.isReadOnly = buf.getByte(offset + 1) != 0;
//       obj.canBeDeleted = buf.getByte(offset + 2) != 0;
//       obj.tree = AssetEditorFileTree.fromValue(buf.getByte(offset + 3));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 12 + buf.getIntLE(offset + 4);
//          int packLen = VarInt.peek(buf, varPos0);
//          if (packLen < 0) {
//             throw ProtocolException.negativeLength("Pack", packLen);
//          }
//
//          if (packLen > 4096000) {
//             throw ProtocolException.stringTooLong("Pack", packLen, 4096000);
//          }
//
//          obj.pack = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 12 + buf.getIntLE(offset + 8);
//          int pathsCount = VarInt.peek(buf, varPos1);
//          if (pathsCount < 0) {
//             throw ProtocolException.negativeLength("Paths", pathsCount);
//          }
//
//          if (pathsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Paths", pathsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + pathsCount * 2L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Paths", varPos1 + varIntLen + pathsCount * 2, buf.readableBytes());
//          }
//
//          obj.paths = new AssetEditorFileEntry[pathsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < pathsCount; i++) {
//             obj.paths[i] = AssetEditorFileEntry.deserialize(buf, elemPos);
//             elemPos += AssetEditorFileEntry.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.pack != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.paths != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.isReadOnly ? 1 : 0);
//       buf.writeByte(this.canBeDeleted ? 1 : 0);
//       buf.writeByte(this.tree.getValue());
//       int packOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int pathsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.pack != null) {
//          buf.setIntLE(packOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.pack, 4096000);
//       } else {
//          buf.setIntLE(packOffsetSlot, -1);
//       }
//
//       if (this.paths != null) {
//          buf.setIntLE(pathsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.paths.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Paths", this.paths.length, 4096000);
//          }
//
//          VarInt.write(buf, this.paths.length);
//
//          for (AssetEditorFileEntry item : this.paths) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(pathsOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorAssetListSetup) Marshal(io protocol.IO) {
	// TODO: Implement
}
