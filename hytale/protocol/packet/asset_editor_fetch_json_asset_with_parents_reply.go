package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorFetchJsonAssetWithParentsReply struct{}

func (*AssetEditorFetchJsonAssetWithParentsReply) ID() uint32 {
	return IDAssetEditorFetchJsonAssetWithParentsReply
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorFetchJsonAssetWithParentsReply obj = new AssetEditorFetchJsonAssetWithParentsReply();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          int assetsCount = VarInt.peek(buf, pos);
//          if (assetsCount < 0) {
//             throw ProtocolException.negativeLength("Assets", assetsCount);
//          }
//
//          if (assetsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Assets", assetsCount, 4096000);
//          }
//
//          pos += VarInt.size(assetsCount);
//          obj.assets = new HashMap<>(assetsCount);
//
//          for (int i = 0; i < assetsCount; i++) {
//             AssetPath key = AssetPath.deserialize(buf, pos);
//             pos += AssetPath.computeBytesConsumed(buf, pos);
//             int valLen = VarInt.peek(buf, pos);
//             if (valLen < 0) {
//                throw ProtocolException.negativeLength("val", valLen);
//             }
//
//             if (valLen > 4096000) {
//                throw ProtocolException.stringTooLong("val", valLen, 4096000);
//             }
//
//             int valVarLen = VarInt.length(buf, pos);
//             String val = PacketIO.readVarString(buf, pos);
//             pos += valVarLen + valLen;
//             if (obj.assets.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("assets", key);
//             }
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
//       buf.writeIntLE(this.token);
//       if (this.assets != null) {
//          if (this.assets.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Assets", this.assets.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.assets.size());
//
//          for (Entry<AssetPath, String> e : this.assets.entrySet()) {
//             e.getKey().serialize(buf);
//             PacketIO.writeVarString(buf, e.getValue(), 4096000);
//          }
//       }
//    }

func (pk *AssetEditorFetchJsonAssetWithParentsReply) Marshal(io protocol.IO) {
	// TODO: Implement
}
