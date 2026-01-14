package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type MoveItemStack struct{}

func (*MoveItemStack) ID() uint32 {
	return IDMoveItemStack
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       MoveItemStack obj = new MoveItemStack();
//       obj.fromSectionId = buf.getIntLE(offset + 0);
//       obj.fromSlotId = buf.getIntLE(offset + 4);
//       obj.quantity = buf.getIntLE(offset + 8);
//       obj.toSectionId = buf.getIntLE(offset + 12);
//       obj.toSlotId = buf.getIntLE(offset + 16);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.fromSectionId);
//       buf.writeIntLE(this.fromSlotId);
//       buf.writeIntLE(this.quantity);
//       buf.writeIntLE(this.toSectionId);
//       buf.writeIntLE(this.toSlotId);
//    }

func (pk *MoveItemStack) Marshal(io protocol.IO) {
	// TODO: Implement
}
