package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorRebuildCaches struct{}

func (*AssetEditorRebuildCaches) ID() uint32 {
	return IDAssetEditorRebuildCaches
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorRebuildCaches obj = new AssetEditorRebuildCaches();
//       obj.blockTextures = buf.getByte(offset + 0) != 0;
//       obj.models = buf.getByte(offset + 1) != 0;
//       obj.modelTextures = buf.getByte(offset + 2) != 0;
//       obj.mapGeometry = buf.getByte(offset + 3) != 0;
//       obj.itemIcons = buf.getByte(offset + 4) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.blockTextures ? 1 : 0);
//       buf.writeByte(this.models ? 1 : 0);
//       buf.writeByte(this.modelTextures ? 1 : 0);
//       buf.writeByte(this.mapGeometry ? 1 : 0);
//       buf.writeByte(this.itemIcons ? 1 : 0);
//    }

func (pk *AssetEditorRebuildCaches) Marshal(io protocol.IO) {
	// TODO: Implement
}
