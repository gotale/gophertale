package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorRenameDirectory struct{}

func (*AssetEditorRenameDirectory) ID() uint32 {
	return IDAssetEditorRenameDirectory
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorRenameDirectory obj = new AssetEditorRenameDirectory();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 13 + buf.getIntLE(offset + 5);
//          obj.path = AssetPath.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 13 + buf.getIntLE(offset + 9);
//          obj.newPath = AssetPath.deserialize(buf, varPos1);
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
//       if (this.newPath != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       int pathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int newPathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.path != null) {
//          buf.setIntLE(pathOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.path.serialize(buf);
//       } else {
//          buf.setIntLE(pathOffsetSlot, -1);
//       }
//
//       if (this.newPath != null) {
//          buf.setIntLE(newPathOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.newPath.serialize(buf);
//       } else {
//          buf.setIntLE(newPathOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorRenameDirectory) Marshal(io protocol.IO) {
	// TODO: Implement
}
