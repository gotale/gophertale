package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolHideAnchors struct{}

func (*BuilderToolHideAnchors) ID() uint32 {
	return IDBuilderToolHideAnchors
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new BuilderToolHideAnchors();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *BuilderToolHideAnchors) Marshal(io protocol.IO) {
	// TODO: Implement
}
