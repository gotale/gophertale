package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorFetchJsonAssetWithParents struct{}

func (*AssetEditorFetchJsonAssetWithParents) ID() uint32 {
	return IDAssetEditorFetchJsonAssetWithParents
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorFetchJsonAssetWithParents obj = new AssetEditorFetchJsonAssetWithParents();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       obj.isFromOpenedTab = buf.getByte(offset + 5) != 0;
//       int pos = offset + 6;
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
//       buf.writeByte(this.isFromOpenedTab ? 1 : 0);
//       if (this.path != null) {
//          this.path.serialize(buf);
//       }
//    }

func (pk *AssetEditorFetchJsonAssetWithParents) Marshal(io protocol.IO) {
	// TODO: Implement
}
