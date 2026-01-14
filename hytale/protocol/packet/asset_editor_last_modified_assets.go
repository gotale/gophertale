package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorLastModifiedAssets struct{}

func (*AssetEditorLastModifiedAssets) ID() uint32 {
	return IDAssetEditorLastModifiedAssets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorLastModifiedAssets obj = new AssetEditorLastModifiedAssets();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int assetsCount = VarInt.peek(buf, pos);
//          if (assetsCount < 0) {
//             throw ProtocolException.negativeLength("Assets", assetsCount);
//          }
//
//          if (assetsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Assets", assetsCount, 4096000);
//          }
//
//          int assetsVarLen = VarInt.size(assetsCount);
//          if (pos + assetsVarLen + assetsCount * 11L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Assets", pos + assetsVarLen + assetsCount * 11, buf.readableBytes());
//          }
//
//          pos += assetsVarLen;
//          obj.assets = new AssetInfo[assetsCount];
//
//          for (int i = 0; i < assetsCount; i++) {
//             obj.assets[i] = AssetInfo.deserialize(buf, pos);
//             pos += AssetInfo.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.assets != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.assets != null) {
//          if (this.assets.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Assets", this.assets.length, 4096000);
//          }
//
//          VarInt.write(buf, this.assets.length);
//
//          for (AssetInfo item : this.assets) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *AssetEditorLastModifiedAssets) Marshal(io protocol.IO) {
	// TODO: Implement
}
