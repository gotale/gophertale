package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorFetchAssetReply struct{}

func (*AssetEditorFetchAssetReply) ID() uint32 {
	return IDAssetEditorFetchAssetReply
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorFetchAssetReply obj = new AssetEditorFetchAssetReply();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          int contentsCount = VarInt.peek(buf, pos);
//          if (contentsCount < 0) {
//             throw ProtocolException.negativeLength("Contents", contentsCount);
//          }
//
//          if (contentsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Contents", contentsCount, 4096000);
//          }
//
//          int contentsVarLen = VarInt.size(contentsCount);
//          if (pos + contentsVarLen + contentsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Contents", pos + contentsVarLen + contentsCount * 1, buf.readableBytes());
//          }
//
//          pos += contentsVarLen;
//          obj.contents = new byte[contentsCount];
//
//          for (int i = 0; i < contentsCount; i++) {
//             obj.contents[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += contentsCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.contents != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       if (this.contents != null) {
//          if (this.contents.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Contents", this.contents.length, 4096000);
//          }
//
//          VarInt.write(buf, this.contents.length);
//
//          for (byte item : this.contents) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *AssetEditorFetchAssetReply) Marshal(io protocol.IO) {
	// TODO: Implement
}
