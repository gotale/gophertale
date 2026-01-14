package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUpdateModelPreview struct{}

func (*AssetEditorUpdateModelPreview) ID() uint32 {
	return IDAssetEditorUpdateModelPreview
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUpdateModelPreview obj = new AssetEditorUpdateModelPreview();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 8) != 0) {
//          obj.camera = AssetEditorPreviewCameraSettings.deserialize(buf, offset + 1);
//       }
//
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 42 + buf.getIntLE(offset + 30);
//          obj.assetPath = AssetPath.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 42 + buf.getIntLE(offset + 34);
//          obj.model = Model.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 42 + buf.getIntLE(offset + 38);
//          obj.block = BlockType.deserialize(buf, varPos2);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.assetPath != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.model != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.block != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       if (this.camera != null) {
//          nullBits = (byte)(nullBits | 8);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.camera != null) {
//          this.camera.serialize(buf);
//       } else {
//          buf.writeZero(29);
//       }
//
//       int assetPathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int modelOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int blockOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.assetPath != null) {
//          buf.setIntLE(assetPathOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.assetPath.serialize(buf);
//       } else {
//          buf.setIntLE(assetPathOffsetSlot, -1);
//       }
//
//       if (this.model != null) {
//          buf.setIntLE(modelOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.model.serialize(buf);
//       } else {
//          buf.setIntLE(modelOffsetSlot, -1);
//       }
//
//       if (this.block != null) {
//          buf.setIntLE(blockOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.block.serialize(buf);
//       } else {
//          buf.setIntLE(blockOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorUpdateModelPreview) Marshal(io protocol.IO) {
	// TODO: Implement
}
