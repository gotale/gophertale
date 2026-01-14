package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerMessage struct{}

func (*ServerMessage) ID() uint32 {
	return IDServerMessage
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerMessage obj = new ServerMessage();
//       byte nullBits = buf.getByte(offset);
//       obj.type = ChatType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          obj.message = FormattedMessage.deserialize(buf, pos);
//          pos += FormattedMessage.computeBytesConsumed(buf, pos);
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
//       buf.writeByte(this.type.getValue());
//       if (this.message != null) {
//          this.message.serialize(buf);
//       }
//    }

func (pk *ServerMessage) Marshal(io protocol.IO) {
	// TODO: Implement
}
