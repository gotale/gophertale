package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type FailureReply struct{}

func (*FailureReply) ID() uint32 {
	return IDFailureReply
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       FailureReply obj = new FailureReply();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
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
//       buf.writeIntLE(this.token);
//       if (this.message != null) {
//          this.message.serialize(buf);
//       }
//    }

func (pk *FailureReply) Marshal(io protocol.IO) {
	// TODO: Implement
}
