package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSetTransformationModeState struct{}

func (*BuilderToolSetTransformationModeState) ID() uint32 {
	return IDBuilderToolSetTransformationModeState
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSetTransformationModeState obj = new BuilderToolSetTransformationModeState();
//       obj.enabled = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.enabled ? 1 : 0);
//    }

func (pk *BuilderToolSetTransformationModeState) Marshal(io protocol.IO) {
	// TODO: Implement
}
