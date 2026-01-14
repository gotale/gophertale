package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorExportAssets struct{}

func (*AssetEditorExportAssets) ID() uint32 {
	return IDAssetEditorExportAssets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorExportAssets obj = new AssetEditorExportAssets();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int pathsCount = VarInt.peek(buf, pos);
//          if (pathsCount < 0) {
//             throw ProtocolException.negativeLength("Paths", pathsCount);
//          }
//
//          if (pathsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Paths", pathsCount, 4096000);
//          }
//
//          int pathsVarLen = VarInt.size(pathsCount);
//          if (pos + pathsVarLen + pathsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Paths", pos + pathsVarLen + pathsCount * 1, buf.readableBytes());
//          }
//
//          pos += pathsVarLen;
//          obj.paths = new AssetPath[pathsCount];
//
//          for (int i = 0; i < pathsCount; i++) {
//             obj.paths[i] = AssetPath.deserialize(buf, pos);
//             pos += AssetPath.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.paths != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.paths != null) {
//          if (this.paths.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Paths", this.paths.length, 4096000);
//          }
//
//          VarInt.write(buf, this.paths.length);
//
//          for (AssetPath item : this.paths) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *AssetEditorExportAssets) Marshal(io protocol.IO) {
	// TODO: Implement
}
