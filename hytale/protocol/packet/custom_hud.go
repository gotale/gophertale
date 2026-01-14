package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type CustomHud struct{}

func (*CustomHud) ID() uint32 {
	return IDCustomHud
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       CustomHud obj = new CustomHud();
//       byte nullBits = buf.getByte(offset);
//       obj.clear = buf.getByte(offset + 1) != 0;
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int commandsCount = VarInt.peek(buf, pos);
//          if (commandsCount < 0) {
//             throw ProtocolException.negativeLength("Commands", commandsCount);
//          }
//
//          if (commandsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Commands", commandsCount, 4096000);
//          }
//
//          int commandsVarLen = VarInt.size(commandsCount);
//          if (pos + commandsVarLen + commandsCount * 2L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Commands", pos + commandsVarLen + commandsCount * 2, buf.readableBytes());
//          }
//
//          pos += commandsVarLen;
//          obj.commands = new CustomUICommand[commandsCount];
//
//          for (int i = 0; i < commandsCount; i++) {
//             obj.commands[i] = CustomUICommand.deserialize(buf, pos);
//             pos += CustomUICommand.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.commands != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.clear ? 1 : 0);
//       if (this.commands != null) {
//          if (this.commands.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Commands", this.commands.length, 4096000);
//          }
//
//          VarInt.write(buf, this.commands.length);
//
//          for (CustomUICommand item : this.commands) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *CustomHud) Marshal(io protocol.IO) {
	// TODO: Implement
}
