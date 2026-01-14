package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUpdateWeatherPreviewLock struct{}

func (*AssetEditorUpdateWeatherPreviewLock) ID() uint32 {
	return IDAssetEditorUpdateWeatherPreviewLock
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUpdateWeatherPreviewLock obj = new AssetEditorUpdateWeatherPreviewLock();
//       obj.locked = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.locked ? 1 : 0);
//    }

func (pk *AssetEditorUpdateWeatherPreviewLock) Marshal(io protocol.IO) {
	// TODO: Implement
}
