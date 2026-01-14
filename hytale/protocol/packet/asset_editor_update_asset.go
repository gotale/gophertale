package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUpdateAsset struct{}

func (*AssetEditorUpdateAsset) ID() uint32 {
	return IDAssetEditorUpdateAsset
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUpdateAsset obj = new AssetEditorUpdateAsset();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       obj.assetIndex = buf.getIntLE(offset + 5);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 21 + buf.getIntLE(offset + 9);
//          int assetTypeLen = VarInt.peek(buf, varPos0);
//          if (assetTypeLen < 0) {
//             throw ProtocolException.negativeLength("AssetType", assetTypeLen);
//          }
//
//          if (assetTypeLen > 4096000) {
//             throw ProtocolException.stringTooLong("AssetType", assetTypeLen, 4096000);
//          }
//
//          obj.assetType = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 21 + buf.getIntLE(offset + 13);
//          obj.path = AssetPath.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 21 + buf.getIntLE(offset + 17);
//          int dataCount = VarInt.peek(buf, varPos2);
//          if (dataCount < 0) {
//             throw ProtocolException.negativeLength("Data", dataCount);
//          }
//
//          if (dataCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Data", dataCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos2);
//          if (varPos2 + varIntLen + dataCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Data", varPos2 + varIntLen + dataCount * 1, buf.readableBytes());
//          }
//
//          obj.data = new byte[dataCount];
//
//          for (int i = 0; i < dataCount; i++) {
//             obj.data[i] = buf.getByte(varPos2 + varIntLen + i * 1);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.assetType != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.path != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.data != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       buf.writeIntLE(this.assetIndex);
//       int assetTypeOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int pathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int dataOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.assetType != null) {
//          buf.setIntLE(assetTypeOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.assetType, 4096000);
//       } else {
//          buf.setIntLE(assetTypeOffsetSlot, -1);
//       }
//
//       if (this.path != null) {
//          buf.setIntLE(pathOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.path.serialize(buf);
//       } else {
//          buf.setIntLE(pathOffsetSlot, -1);
//       }
//
//       if (this.data != null) {
//          buf.setIntLE(dataOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.data.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Data", this.data.length, 4096000);
//          }
//
//          VarInt.write(buf, this.data.length);
//
//          for (byte item : this.data) {
//             buf.writeByte(item);
//          }
//       } else {
//          buf.setIntLE(dataOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorUpdateAsset) Marshal(io protocol.IO) {
	// TODO: Implement
}
