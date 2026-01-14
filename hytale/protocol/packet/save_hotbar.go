package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SaveHotbar struct{}

func (*SaveHotbar) ID() uint32 {
	return IDSaveHotbar
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SaveHotbar obj = new SaveHotbar();
//       obj.inventoryRow = buf.getByte(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.inventoryRow);
//    }

func (pk *SaveHotbar) Marshal(io protocol.IO) {
	// TODO: Implement
}
