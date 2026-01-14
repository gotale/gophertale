package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorRequestDatasetReply struct{}

func (*AssetEditorRequestDatasetReply) ID() uint32 {
	return IDAssetEditorRequestDatasetReply
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorRequestDatasetReply obj = new AssetEditorRequestDatasetReply();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          int nameLen = VarInt.peek(buf, varPos0);
//          if (nameLen < 0) {
//             throw ProtocolException.negativeLength("Name", nameLen);
//          }
//
//          if (nameLen > 4096000) {
//             throw ProtocolException.stringTooLong("Name", nameLen, 4096000);
//          }
//
//          obj.name = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int idsCount = VarInt.peek(buf, varPos1);
//          if (idsCount < 0) {
//             throw ProtocolException.negativeLength("Ids", idsCount);
//          }
//
//          if (idsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Ids", idsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + idsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Ids", varPos1 + varIntLen + idsCount * 1, buf.readableBytes());
//          }
//
//          obj.ids = new String[idsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < idsCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("ids[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("ids[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.ids[i] = PacketIO.readVarString(buf, elemPos);
//             elemPos += strVarLen + strLen;
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.name != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.ids != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int nameOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int idsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.name != null) {
//          buf.setIntLE(nameOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.name, 4096000);
//       } else {
//          buf.setIntLE(nameOffsetSlot, -1);
//       }
//
//       if (this.ids != null) {
//          buf.setIntLE(idsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.ids.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Ids", this.ids.length, 4096000);
//          }
//
//          VarInt.write(buf, this.ids.length);
//
//          for (String item : this.ids) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(idsOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorRequestDatasetReply) Marshal(io protocol.IO) {
	// TODO: Implement
}
