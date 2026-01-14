package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorAuthorization struct{}

func (*AssetEditorAuthorization) ID() uint32 {
	return IDAssetEditorAuthorization
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorAuthorization obj = new AssetEditorAuthorization();
//       obj.canUse = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.canUse ? 1 : 0);
//    }

func (pk *AssetEditorAuthorization) Marshal(io protocol.IO) {
	// TODO: Implement
}
