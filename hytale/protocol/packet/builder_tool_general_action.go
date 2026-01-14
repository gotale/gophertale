package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolGeneralAction struct{}

func (*BuilderToolGeneralAction) ID() uint32 {
	return IDBuilderToolGeneralAction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolGeneralAction obj = new BuilderToolGeneralAction();
//       obj.action = BuilderToolAction.fromValue(buf.getByte(offset + 0));
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.action.getValue());
//    }

func (pk *BuilderToolGeneralAction) Marshal(io protocol.IO) {
	// TODO: Implement
}
