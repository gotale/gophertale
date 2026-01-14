package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClearEditorTimeOverride struct{}

func (*ClearEditorTimeOverride) ID() uint32 {
	return IDClearEditorTimeOverride
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new ClearEditorTimeOverride();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *ClearEditorTimeOverride) Marshal(io protocol.IO) {
	// TODO: Implement
}
