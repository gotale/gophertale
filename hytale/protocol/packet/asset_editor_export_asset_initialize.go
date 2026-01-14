package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorExportAssetInitialize struct{}

func (*AssetEditorExportAssetInitialize) ID() uint32 {
	return IDAssetEditorExportAssetInitialize
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorExportAssetInitialize obj = new AssetEditorExportAssetInitialize();
//       byte nullBits = buf.getByte(offset);
//       obj.size = buf.getIntLE(offset + 1);
//       obj.failed = buf.getByte(offset + 5) != 0;
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 14 + buf.getIntLE(offset + 6);
//          obj.asset = AssetEditorAsset.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 14 + buf.getIntLE(offset + 10);
//          obj.oldPath = AssetPath.deserialize(buf, varPos1);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.asset != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.oldPath != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.size);
//       buf.writeByte(this.failed ? 1 : 0);
//       int assetOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int oldPathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.asset != null) {
//          buf.setIntLE(assetOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.asset.serialize(buf);
//       } else {
//          buf.setIntLE(assetOffsetSlot, -1);
//       }
//
//       if (this.oldPath != null) {
//          buf.setIntLE(oldPathOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.oldPath.serialize(buf);
//       } else {
//          buf.setIntLE(oldPathOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorExportAssetInitialize) Marshal(io protocol.IO) {
	// TODO: Implement
}
