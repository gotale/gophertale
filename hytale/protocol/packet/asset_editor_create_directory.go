package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorCreateDirectory struct{}

func (*AssetEditorCreateDirectory) ID() uint32 {
	return IDAssetEditorCreateDirectory
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorCreateDirectory obj = new AssetEditorCreateDirectory();
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

func (pk *AssetEditorCreateDirectory) Marshal(io protocol.IO) {
	// TODO: Implement
}
