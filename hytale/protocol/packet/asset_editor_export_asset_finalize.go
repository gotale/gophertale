package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorExportAssetFinalize struct{}

func (*AssetEditorExportAssetFinalize) ID() uint32 {
	return IDAssetEditorExportAssetFinalize
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new AssetEditorExportAssetFinalize();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *AssetEditorExportAssetFinalize) Marshal(io protocol.IO) {
	// TODO: Implement
}
