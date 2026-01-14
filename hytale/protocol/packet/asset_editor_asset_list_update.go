package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorAssetListUpdate struct{}

func (*AssetEditorAssetListUpdate) ID() uint32 {
	return IDAssetEditorAssetListUpdate
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorAssetListUpdate obj = new AssetEditorAssetListUpdate();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 13 + buf.getIntLE(offset + 1);
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
//          int varPos1 = offset + 13 + buf.getIntLE(offset + 5);
//          int additionsCount = VarInt.peek(buf, varPos1);
//          if (additionsCount < 0) {
//             throw ProtocolException.negativeLength("Additions", additionsCount);
//          }
//
//          if (additionsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Additions", additionsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + additionsCount * 2L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Additions", varPos1 + varIntLen + additionsCount * 2, buf.readableBytes());
//          }
//
//          obj.additions = new AssetEditorFileEntry[additionsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < additionsCount; i++) {
//             obj.additions[i] = AssetEditorFileEntry.deserialize(buf, elemPos);
//             elemPos += AssetEditorFileEntry.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 13 + buf.getIntLE(offset + 9);
//          int deletionsCount = VarInt.peek(buf, varPos2);
//          if (deletionsCount < 0) {
//             throw ProtocolException.negativeLength("Deletions", deletionsCount);
//          }
//
//          if (deletionsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Deletions", deletionsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos2);
//          if (varPos2 + varIntLen + deletionsCount * 2L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Deletions", varPos2 + varIntLen + deletionsCount * 2, buf.readableBytes());
//          }
//
//          obj.deletions = new AssetEditorFileEntry[deletionsCount];
//          int elemPos = varPos2 + varIntLen;
//
//          for (int i = 0; i < deletionsCount; i++) {
//             obj.deletions[i] = AssetEditorFileEntry.deserialize(buf, elemPos);
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
//       if (this.additions != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.deletions != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       int packOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int additionsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int deletionsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.pack != null) {
//          buf.setIntLE(packOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.pack, 4096000);
//       } else {
//          buf.setIntLE(packOffsetSlot, -1);
//       }
//
//       if (this.additions != null) {
//          buf.setIntLE(additionsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.additions.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Additions", this.additions.length, 4096000);
//          }
//
//          VarInt.write(buf, this.additions.length);
//
//          for (AssetEditorFileEntry item : this.additions) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(additionsOffsetSlot, -1);
//       }
//
//       if (this.deletions != null) {
//          buf.setIntLE(deletionsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.deletions.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Deletions", this.deletions.length, 4096000);
//          }
//
//          VarInt.write(buf, this.deletions.length);
//
//          for (AssetEditorFileEntry item : this.deletions) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(deletionsOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorAssetListUpdate) Marshal(io protocol.IO) {
	// TODO: Implement
}
