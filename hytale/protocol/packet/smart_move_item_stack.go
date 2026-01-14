package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SmartMoveItemStack struct{}

func (*SmartMoveItemStack) ID() uint32 {
	return IDSmartMoveItemStack
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SmartMoveItemStack obj = new SmartMoveItemStack();
//       obj.fromSectionId = buf.getIntLE(offset + 0);
//       obj.fromSlotId = buf.getIntLE(offset + 4);
//       obj.quantity = buf.getIntLE(offset + 8);
//       obj.moveType = SmartMoveType.fromValue(buf.getByte(offset + 12));
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.fromSectionId);
//       buf.writeIntLE(this.fromSlotId);
//       buf.writeIntLE(this.quantity);
//       buf.writeByte(this.moveType.getValue());
//    }

func (pk *SmartMoveItemStack) Marshal(io protocol.IO) {
	// TODO: Implement
}
