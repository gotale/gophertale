package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type LoadHotbar struct{}

func (*LoadHotbar) ID() uint32 {
	return IDLoadHotbar
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       LoadHotbar obj = new LoadHotbar();
//       obj.inventoryRow = buf.getByte(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.inventoryRow);
//    }

func (pk *LoadHotbar) Marshal(io protocol.IO) {
	// TODO: Implement
}
