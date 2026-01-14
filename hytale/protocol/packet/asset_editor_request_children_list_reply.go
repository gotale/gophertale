package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorRequestChildrenListReply struct{}

func (*AssetEditorRequestChildrenListReply) ID() uint32 {
	return IDAssetEditorRequestChildrenListReply
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorRequestChildrenListReply obj = new AssetEditorRequestChildrenListReply();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          obj.path = AssetPath.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int childrenIdsCount = VarInt.peek(buf, varPos1);
//          if (childrenIdsCount < 0) {
//             throw ProtocolException.negativeLength("ChildrenIds", childrenIdsCount);
//          }
//
//          if (childrenIdsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("ChildrenIds", childrenIdsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + childrenIdsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("ChildrenIds", varPos1 + varIntLen + childrenIdsCount * 1, buf.readableBytes());
//          }
//
//          obj.childrenIds = new String[childrenIdsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < childrenIdsCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("childrenIds[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("childrenIds[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.childrenIds[i] = PacketIO.readVarString(buf, elemPos);
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
//       if (this.path != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.childrenIds != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int pathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int childrenIdsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.path != null) {
//          buf.setIntLE(pathOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.path.serialize(buf);
//       } else {
//          buf.setIntLE(pathOffsetSlot, -1);
//       }
//
//       if (this.childrenIds != null) {
//          buf.setIntLE(childrenIdsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.childrenIds.length > 4096000) {
//             throw ProtocolException.arrayTooLong("ChildrenIds", this.childrenIds.length, 4096000);
//          }
//
//          VarInt.write(buf, this.childrenIds.length);
//
//          for (String item : this.childrenIds) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(childrenIdsOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorRequestChildrenListReply) Marshal(io protocol.IO) {
	// TODO: Implement
}
