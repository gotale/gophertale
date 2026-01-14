package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetActiveSlot struct{}

func (*SetActiveSlot) ID() uint32 {
	return IDSetActiveSlot
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetActiveSlot obj = new SetActiveSlot();
//       obj.inventorySectionId = buf.getIntLE(offset + 0);
//       obj.activeSlot = buf.getIntLE(offset + 4);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.inventorySectionId);
//       buf.writeIntLE(this.activeSlot);
//    }

func (pk *SetActiveSlot) Marshal(io protocol.IO) {
	// TODO: Implement
}
