package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorCreateAssetPack struct{}

func (*AssetEditorCreateAssetPack) ID() uint32 {
	return IDAssetEditorCreateAssetPack
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorCreateAssetPack obj = new AssetEditorCreateAssetPack();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          obj.manifest = AssetPackManifest.deserialize(buf, pos);
//          pos += AssetPackManifest.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.manifest != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       if (this.manifest != null) {
//          this.manifest.serialize(buf);
//       }
//    }

func (pk *AssetEditorCreateAssetPack) Marshal(io protocol.IO) {
	// TODO: Implement
}
