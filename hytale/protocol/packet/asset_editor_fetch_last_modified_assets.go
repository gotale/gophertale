package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorFetchLastModifiedAssets struct{}

func (*AssetEditorFetchLastModifiedAssets) ID() uint32 {
	return IDAssetEditorFetchLastModifiedAssets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new AssetEditorFetchLastModifiedAssets();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *AssetEditorFetchLastModifiedAssets) Marshal(io protocol.IO) {
	// TODO: Implement
}
