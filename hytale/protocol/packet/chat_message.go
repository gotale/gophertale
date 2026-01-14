package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ChatMessage struct{}

func (*ChatMessage) ID() uint32 {
	return IDChatMessage
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ChatMessage obj = new ChatMessage();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int messageLen = VarInt.peek(buf, pos);
//          if (messageLen < 0) {
//             throw ProtocolException.negativeLength("Message", messageLen);
//          }
//
//          if (messageLen > 4096000) {
//             throw ProtocolException.stringTooLong("Message", messageLen, 4096000);
//          }
//
//          int messageVarLen = VarInt.length(buf, pos);
//          obj.message = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += messageVarLen + messageLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.message != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.message != null) {
//          PacketIO.writeVarString(buf, this.message, 4096000);
//       }
//    }

func (pk *ChatMessage) Marshal(io protocol.IO) {
	// TODO: Implement
}
