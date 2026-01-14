package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorSetupAssetTypes struct{}

func (*AssetEditorSetupAssetTypes) ID() uint32 {
	return IDAssetEditorSetupAssetTypes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorSetupAssetTypes obj = new AssetEditorSetupAssetTypes();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int assetTypesCount = VarInt.peek(buf, pos);
//          if (assetTypesCount < 0) {
//             throw ProtocolException.negativeLength("AssetTypes", assetTypesCount);
//          }
//
//          if (assetTypesCount > 4096000) {
//             throw ProtocolException.arrayTooLong("AssetTypes", assetTypesCount, 4096000);
//          }
//
//          int assetTypesVarLen = VarInt.size(assetTypesCount);
//          if (pos + assetTypesVarLen + assetTypesCount * 3L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("AssetTypes", pos + assetTypesVarLen + assetTypesCount * 3, buf.readableBytes());
//          }
//
//          pos += assetTypesVarLen;
//          obj.assetTypes = new AssetEditorAssetType[assetTypesCount];
//
//          for (int i = 0; i < assetTypesCount; i++) {
//             obj.assetTypes[i] = AssetEditorAssetType.deserialize(buf, pos);
//             pos += AssetEditorAssetType.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.assetTypes != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.assetTypes != null) {
//          if (this.assetTypes.length > 4096000) {
//             throw ProtocolException.arrayTooLong("AssetTypes", this.assetTypes.length, 4096000);
//          }
//
//          VarInt.write(buf, this.assetTypes.length);
//
//          for (AssetEditorAssetType item : this.assetTypes) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *AssetEditorSetupAssetTypes) Marshal(io protocol.IO) {
	// TODO: Implement
}
