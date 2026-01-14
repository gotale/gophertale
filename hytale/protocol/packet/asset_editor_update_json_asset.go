package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUpdateJsonAsset struct{}

func (*AssetEditorUpdateJsonAsset) ID() uint32 {
	return IDAssetEditorUpdateJsonAsset
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUpdateJsonAsset obj = new AssetEditorUpdateJsonAsset();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       obj.assetIndex = buf.getIntLE(offset + 5);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 21 + buf.getIntLE(offset + 9);
//          int assetTypeLen = VarInt.peek(buf, varPos0);
//          if (assetTypeLen < 0) {
//             throw ProtocolException.negativeLength("AssetType", assetTypeLen);
//          }
//
//          if (assetTypeLen > 4096000) {
//             throw ProtocolException.stringTooLong("AssetType", assetTypeLen, 4096000);
//          }
//
//          obj.assetType = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 21 + buf.getIntLE(offset + 13);
//          obj.path = AssetPath.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 21 + buf.getIntLE(offset + 17);
//          int commandsCount = VarInt.peek(buf, varPos2);
//          if (commandsCount < 0) {
//             throw ProtocolException.negativeLength("Commands", commandsCount);
//          }
//
//          if (commandsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Commands", commandsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos2);
//          if (varPos2 + varIntLen + commandsCount * 7L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Commands", varPos2 + varIntLen + commandsCount * 7, buf.readableBytes());
//          }
//
//          obj.commands = new JsonUpdateCommand[commandsCount];
//          int elemPos = varPos2 + varIntLen;
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
//       if (this.assetType != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.path != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.commands != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       buf.writeIntLE(this.assetIndex);
//       int assetTypeOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int pathOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int commandsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.assetType != null) {
//          buf.setIntLE(assetTypeOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.assetType, 4096000);
//       } else {
//          buf.setIntLE(assetTypeOffsetSlot, -1);
//       }
//
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

func (pk *AssetEditorUpdateJsonAsset) Marshal(io protocol.IO) {
	// TODO: Implement
}
