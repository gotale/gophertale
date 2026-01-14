package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorRequestChildrenList struct{}

func (*AssetEditorRequestChildrenList) ID() uint32 {
	return IDAssetEditorRequestChildrenList
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorRequestChildrenList obj = new AssetEditorRequestChildrenList();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
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
//       if (this.path != null) {
//          this.path.serialize(buf);
//       }
//    }

func (pk *AssetEditorRequestChildrenList) Marshal(io protocol.IO) {
	// TODO: Implement
}
