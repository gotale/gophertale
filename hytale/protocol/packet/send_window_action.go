package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SendWindowAction struct{}

func (*SendWindowAction) ID() uint32 {
	return IDSendWindowAction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SendWindowAction obj = new SendWindowAction();
//       obj.id = buf.getIntLE(offset + 0);
//       int pos = offset + 4;
//       obj.action = WindowAction.deserialize(buf, pos);
//       pos += WindowAction.computeBytesConsumed(buf, pos);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.id);
//       this.action.serializeWithTypeId(buf);
//    }

func (pk *SendWindowAction) Marshal(io protocol.IO) {
	// TODO: Implement
}
