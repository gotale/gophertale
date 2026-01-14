package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorFetchAutoCompleteDataReply struct{}

func (*AssetEditorFetchAutoCompleteDataReply) ID() uint32 {
	return IDAssetEditorFetchAutoCompleteDataReply
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorFetchAutoCompleteDataReply obj = new AssetEditorFetchAutoCompleteDataReply();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          int resultsCount = VarInt.peek(buf, pos);
//          if (resultsCount < 0) {
//             throw ProtocolException.negativeLength("Results", resultsCount);
//          }
//
//          if (resultsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Results", resultsCount, 4096000);
//          }
//
//          int resultsVarLen = VarInt.size(resultsCount);
//          if (pos + resultsVarLen + resultsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Results", pos + resultsVarLen + resultsCount * 1, buf.readableBytes());
//          }
//
//          pos += resultsVarLen;
//          obj.results = new String[resultsCount];
//
//          for (int i = 0; i < resultsCount; i++) {
//             int strLen = VarInt.peek(buf, pos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("results[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("results[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, pos);
//             obj.results[i] = PacketIO.readVarString(buf, pos);
//             pos += strVarLen + strLen;
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.results != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       if (this.results != null) {
//          if (this.results.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Results", this.results.length, 4096000);
//          }
//
//          VarInt.write(buf, this.results.length);
//
//          for (String item : this.results) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       }
//    }

func (pk *AssetEditorFetchAutoCompleteDataReply) Marshal(io protocol.IO) {
	// TODO: Implement
}
