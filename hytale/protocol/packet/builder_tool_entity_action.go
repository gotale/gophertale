package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolEntityAction struct{}

func (*BuilderToolEntityAction) ID() uint32 {
	return IDBuilderToolEntityAction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolEntityAction obj = new BuilderToolEntityAction();
//       obj.entityId = buf.getIntLE(offset + 0);
//       obj.action = EntityToolAction.fromValue(buf.getByte(offset + 4));
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.entityId);
//       buf.writeByte(this.action.getValue());
//    }

func (pk *BuilderToolEntityAction) Marshal(io protocol.IO) {
	// TODO: Implement
}
