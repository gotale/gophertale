package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorAssetUpdated struct{}

func (*AssetEditorAssetUpdated) ID() uint32 {
	return IDAssetEditorAssetUpdated
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorAssetUpdated obj = new AssetEditorAssetUpdated();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          obj.path = AssetPath.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int dataCount = VarInt.peek(buf, varPos1);
//          if (dataCount < 0) {
//             throw ProtocolException.negativeLength("Data", dataCount);
//          }
//
//          if (dataCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Data", dataCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + dataCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Data", varPos1 + varIntLen + dataCount * 1, buf.readableBytes());
//          }
//
//          obj.data = new byte[dataCount];
//
//          for (int i = 0; i < dataCount; i++) {
//             obj.data[i] = buf.getByte(varPos1 + varIntLen + i * 1);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.path != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.data != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int pathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int dataOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
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

func (pk *AssetEditorAssetUpdated) Marshal(io protocol.IO) {
	// TODO: Implement
}
