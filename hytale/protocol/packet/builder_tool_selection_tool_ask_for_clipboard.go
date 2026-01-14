package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSelectionToolAskForClipboard struct{}

func (*BuilderToolSelectionToolAskForClipboard) ID() uint32 {
	return IDBuilderToolSelectionToolAskForClipboard
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new BuilderToolSelectionToolAskForClipboard();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *BuilderToolSelectionToolAskForClipboard) Marshal(io protocol.IO) {
	// TODO: Implement
}
