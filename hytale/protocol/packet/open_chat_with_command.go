package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type OpenChatWithCommand struct{}

func (*OpenChatWithCommand) ID() uint32 {
	return IDOpenChatWithCommand
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       OpenChatWithCommand obj = new OpenChatWithCommand();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int commandLen = VarInt.peek(buf, pos);
//          if (commandLen < 0) {
//             throw ProtocolException.negativeLength("Command", commandLen);
//          }
//
//          if (commandLen > 4096000) {
//             throw ProtocolException.stringTooLong("Command", commandLen, 4096000);
//          }
//
//          int commandVarLen = VarInt.length(buf, pos);
//          obj.command = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += commandVarLen + commandLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.command != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.command != null) {
//          PacketIO.writeVarString(buf, this.command, 4096000);
//       }
//    }

func (pk *OpenChatWithCommand) Marshal(io protocol.IO) {
	// TODO: Implement
}
