package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorInitialize struct{}

func (*AssetEditorInitialize) ID() uint32 {
	return IDAssetEditorInitialize
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new AssetEditorInitialize();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *AssetEditorInitialize) Marshal(io protocol.IO) {
	// TODO: Implement
}
