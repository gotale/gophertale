package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type DropCreativeItem struct{}

func (*DropCreativeItem) ID() uint32 {
	return IDDropCreativeItem
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       DropCreativeItem obj = new DropCreativeItem();
//       int pos = offset + 0;
//       obj.item = ItemQuantity.deserialize(buf, pos);
//       pos += ItemQuantity.computeBytesConsumed(buf, pos);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       this.item.serialize(buf);
//    }

func (pk *DropCreativeItem) Marshal(io protocol.IO) {
	// TODO: Implement
}
