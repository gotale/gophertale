package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorCapabilities struct{}

func (*AssetEditorCapabilities) ID() uint32 {
	return IDAssetEditorCapabilities
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorCapabilities obj = new AssetEditorCapabilities();
//       obj.canDiscardAssets = buf.getByte(offset + 0) != 0;
//       obj.canEditAssets = buf.getByte(offset + 1) != 0;
//       obj.canCreateAssetPacks = buf.getByte(offset + 2) != 0;
//       obj.canEditAssetPacks = buf.getByte(offset + 3) != 0;
//       obj.canDeleteAssetPacks = buf.getByte(offset + 4) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.canDiscardAssets ? 1 : 0);
//       buf.writeByte(this.canEditAssets ? 1 : 0);
//       buf.writeByte(this.canCreateAssetPacks ? 1 : 0);
//       buf.writeByte(this.canEditAssetPacks ? 1 : 0);
//       buf.writeByte(this.canDeleteAssetPacks ? 1 : 0);
//    }

func (pk *AssetEditorCapabilities) Marshal(io protocol.IO) {
	// TODO: Implement
}
