package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetCreativeItem struct{}

func (*SetCreativeItem) ID() uint32 {
	return IDSetCreativeItem
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetCreativeItem obj = new SetCreativeItem();
//       obj.inventorySectionId = buf.getIntLE(offset + 0);
//       obj.slotId = buf.getIntLE(offset + 4);
//       obj.override = buf.getByte(offset + 8) != 0;
//       int pos = offset + 9;
//       obj.item = ItemQuantity.deserialize(buf, pos);
//       pos += ItemQuantity.computeBytesConsumed(buf, pos);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.inventorySectionId);
//       buf.writeIntLE(this.slotId);
//       buf.writeByte(this.override ? 1 : 0);
//       this.item.serialize(buf);
//    }

func (pk *SetCreativeItem) Marshal(io protocol.IO) {
	// TODO: Implement
}
