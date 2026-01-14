package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorExportDeleteAssets struct{}

func (*AssetEditorExportDeleteAssets) ID() uint32 {
	return IDAssetEditorExportDeleteAssets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorExportDeleteAssets obj = new AssetEditorExportDeleteAssets();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int assetCount = VarInt.peek(buf, pos);
//          if (assetCount < 0) {
//             throw ProtocolException.negativeLength("Asset", assetCount);
//          }
//
//          if (assetCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Asset", assetCount, 4096000);
//          }
//
//          int assetVarLen = VarInt.size(assetCount);
//          if (pos + assetVarLen + assetCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Asset", pos + assetVarLen + assetCount * 1, buf.readableBytes());
//          }
//
//          pos += assetVarLen;
//          obj.asset = new AssetEditorAsset[assetCount];
//
//          for (int i = 0; i < assetCount; i++) {
//             obj.asset[i] = AssetEditorAsset.deserialize(buf, pos);
//             pos += AssetEditorAsset.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.asset != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.asset != null) {
//          if (this.asset.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Asset", this.asset.length, 4096000);
//          }
//
//          VarInt.write(buf, this.asset.length);
//
//          for (AssetEditorAsset item : this.asset) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *AssetEditorExportDeleteAssets) Marshal(io protocol.IO) {
	// TODO: Implement
}
