package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type DropItemStack struct{}

func (*DropItemStack) ID() uint32 {
	return IDDropItemStack
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       DropItemStack obj = new DropItemStack();
//       obj.inventorySectionId = buf.getIntLE(offset + 0);
//       obj.slotId = buf.getIntLE(offset + 4);
//       obj.quantity = buf.getIntLE(offset + 8);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.inventorySectionId);
//       buf.writeIntLE(this.slotId);
//       buf.writeIntLE(this.quantity);
//    }

func (pk *DropItemStack) Marshal(io protocol.IO) {
	// TODO: Implement
}
