package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUndoChanges struct{}

func (*AssetEditorUndoChanges) ID() uint32 {
	return IDAssetEditorUndoChanges
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUndoChanges obj = new AssetEditorUndoChanges();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          obj.path = AssetPath.deserialize(buf, pos);
//          pos += AssetPath.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.path != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       if (this.path != null) {
//          this.path.serialize(buf);
//       }
//    }

func (pk *AssetEditorUndoChanges) Marshal(io protocol.IO) {
	// TODO: Implement
}
