package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type InventoryAction struct{}

func (*InventoryAction) ID() uint32 {
	return IDInventoryAction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       InventoryAction obj = new InventoryAction();
//       obj.inventorySectionId = buf.getIntLE(offset + 0);
//       obj.inventoryActionType = InventoryActionType.fromValue(buf.getByte(offset + 4));
//       obj.actionData = buf.getByte(offset + 5);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.inventorySectionId);
//       buf.writeByte(this.inventoryActionType.getValue());
//       buf.writeByte(this.actionData);
//    }

func (pk *InventoryAction) Marshal(io protocol.IO) {
	// TODO: Implement
}
