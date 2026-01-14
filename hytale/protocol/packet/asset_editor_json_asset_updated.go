package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorJsonAssetUpdated struct{}

func (*AssetEditorJsonAssetUpdated) ID() uint32 {
	return IDAssetEditorJsonAssetUpdated
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorJsonAssetUpdated obj = new AssetEditorJsonAssetUpdated();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          obj.path = AssetPath.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int commandsCount = VarInt.peek(buf, varPos1);
//          if (commandsCount < 0) {
//             throw ProtocolException.negativeLength("Commands", commandsCount);
//          }
//
//          if (commandsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Commands", commandsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + commandsCount * 7L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Commands", varPos1 + varIntLen + commandsCount * 7, buf.readableBytes());
//          }
//
//          obj.commands = new JsonUpdateCommand[commandsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < commandsCount; i++) {
//             obj.commands[i] = JsonUpdateCommand.deserialize(buf, elemPos);
//             elemPos += JsonUpdateCommand.computeBytesConsumed(buf, elemPos);
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
//       if (this.commands != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int pathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int commandsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.path != null) {
//          buf.setIntLE(pathOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.path.serialize(buf);
//       } else {
//          buf.setIntLE(pathOffsetSlot, -1);
//       }
//
//       if (this.commands != null) {
//          buf.setIntLE(commandsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.commands.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Commands", this.commands.length, 4096000);
//          }
//
//          VarInt.write(buf, this.commands.length);
//
//          for (JsonUpdateCommand item : this.commands) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(commandsOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorJsonAssetUpdated) Marshal(io protocol.IO) {
	// TODO: Implement
}
