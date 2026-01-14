package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorFetchAutoCompleteData struct{}

func (*AssetEditorFetchAutoCompleteData) ID() uint32 {
	return IDAssetEditorFetchAutoCompleteData
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorFetchAutoCompleteData obj = new AssetEditorFetchAutoCompleteData();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 13 + buf.getIntLE(offset + 5);
//          int datasetLen = VarInt.peek(buf, varPos0);
//          if (datasetLen < 0) {
//             throw ProtocolException.negativeLength("Dataset", datasetLen);
//          }
//
//          if (datasetLen > 4096000) {
//             throw ProtocolException.stringTooLong("Dataset", datasetLen, 4096000);
//          }
//
//          obj.dataset = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 13 + buf.getIntLE(offset + 9);
//          int queryLen = VarInt.peek(buf, varPos1);
//          if (queryLen < 0) {
//             throw ProtocolException.negativeLength("Query", queryLen);
//          }
//
//          if (queryLen > 4096000) {
//             throw ProtocolException.stringTooLong("Query", queryLen, 4096000);
//          }
//
//          obj.query = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.dataset != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.query != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       int datasetOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int queryOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.dataset != null) {
//          buf.setIntLE(datasetOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.dataset, 4096000);
//       } else {
//          buf.setIntLE(datasetOffsetSlot, -1);
//       }
//
//       if (this.query != null) {
//          buf.setIntLE(queryOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.query, 4096000);
//       } else {
//          buf.setIntLE(queryOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorFetchAutoCompleteData) Marshal(io protocol.IO) {
	// TODO: Implement
}
