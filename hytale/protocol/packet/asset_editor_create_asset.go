package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorCreateAsset struct{}

func (*AssetEditorCreateAsset) ID() uint32 {
	return IDAssetEditorCreateAsset
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorCreateAsset obj = new AssetEditorCreateAsset();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       if ((nullBits & 4) != 0) {
//          obj.rebuildCaches = AssetEditorRebuildCaches.deserialize(buf, offset + 5);
//       }
//
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 22 + buf.getIntLE(offset + 10);
//          obj.path = AssetPath.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 22 + buf.getIntLE(offset + 14);
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
//       if ((nullBits & 8) != 0) {
//          int varPos2 = offset + 22 + buf.getIntLE(offset + 18);
//          int buttonIdLen = VarInt.peek(buf, varPos2);
//          if (buttonIdLen < 0) {
//             throw ProtocolException.negativeLength("ButtonId", buttonIdLen);
//          }
//
//          if (buttonIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("ButtonId", buttonIdLen, 4096000);
//          }
//
//          obj.buttonId = PacketIO.readVarString(buf, varPos2, PacketIO.UTF8);
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
//       if (this.rebuildCaches != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       if (this.buttonId != null) {
//          nullBits = (byte)(nullBits | 8);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       if (this.rebuildCaches != null) {
//          this.rebuildCaches.serialize(buf);
//       } else {
//          buf.writeZero(5);
//       }
//
//       int pathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int dataOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int buttonIdOffsetSlot = buf.writerIndex();
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
//
//       if (this.buttonId != null) {
//          buf.setIntLE(buttonIdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.buttonId, 4096000);
//       } else {
//          buf.setIntLE(buttonIdOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorCreateAsset) Marshal(io protocol.IO) {
	// TODO: Implement
}
