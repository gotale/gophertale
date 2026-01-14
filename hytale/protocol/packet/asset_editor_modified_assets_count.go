package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorModifiedAssetsCount struct{}

func (*AssetEditorModifiedAssetsCount) ID() uint32 {
	return IDAssetEditorModifiedAssetsCount
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorModifiedAssetsCount obj = new AssetEditorModifiedAssetsCount();
//       obj.count = buf.getIntLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.count);
//    }

func (pk *AssetEditorModifiedAssetsCount) Marshal(io protocol.IO) {
	// TODO: Implement
}
