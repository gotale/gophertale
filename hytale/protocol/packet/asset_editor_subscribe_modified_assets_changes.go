package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorSubscribeModifiedAssetsChanges struct{}

func (*AssetEditorSubscribeModifiedAssetsChanges) ID() uint32 {
	return IDAssetEditorSubscribeModifiedAssetsChanges
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorSubscribeModifiedAssetsChanges obj = new AssetEditorSubscribeModifiedAssetsChanges();
//       obj.subscribe = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.subscribe ? 1 : 0);
//    }

func (pk *AssetEditorSubscribeModifiedAssetsChanges) Marshal(io protocol.IO) {
	// TODO: Implement
}
