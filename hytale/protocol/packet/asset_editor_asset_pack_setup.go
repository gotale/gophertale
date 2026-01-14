package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorAssetPackSetup struct{}

func (*AssetEditorAssetPackSetup) ID() uint32 {
	return IDAssetEditorAssetPackSetup
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorAssetPackSetup obj = new AssetEditorAssetPackSetup();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int packsCount = VarInt.peek(buf, pos);
//          if (packsCount < 0) {
//             throw ProtocolException.negativeLength("Packs", packsCount);
//          }
//
//          if (packsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Packs", packsCount, 4096000);
//          }
//
//          pos += VarInt.size(packsCount);
//          obj.packs = new HashMap<>(packsCount);
//
//          for (int i = 0; i < packsCount; i++) {
//             int keyLen = VarInt.peek(buf, pos);
//             if (keyLen < 0) {
//                throw ProtocolException.negativeLength("key", keyLen);
//             }
//
//             if (keyLen > 4096000) {
//                throw ProtocolException.stringTooLong("key", keyLen, 4096000);
//             }
//
//             int keyVarLen = VarInt.length(buf, pos);
//             String key = PacketIO.readVarString(buf, pos);
//             pos += keyVarLen + keyLen;
//             AssetPackManifest val = AssetPackManifest.deserialize(buf, pos);
//             pos += AssetPackManifest.computeBytesConsumed(buf, pos);
//             if (obj.packs.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("packs", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.packs != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.packs != null) {
//          if (this.packs.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Packs", this.packs.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.packs.size());
//
//          for (Entry<String, AssetPackManifest> e : this.packs.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *AssetEditorAssetPackSetup) Marshal(io protocol.IO) {
	// TODO: Implement
}
