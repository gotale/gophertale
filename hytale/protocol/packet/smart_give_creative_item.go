package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SmartGiveCreativeItem struct{}

func (*SmartGiveCreativeItem) ID() uint32 {
	return IDSmartGiveCreativeItem
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SmartGiveCreativeItem obj = new SmartGiveCreativeItem();
//       obj.moveType = SmartMoveType.fromValue(buf.getByte(offset + 0));
//       int pos = offset + 1;
//       obj.item = ItemQuantity.deserialize(buf, pos);
//       pos += ItemQuantity.computeBytesConsumed(buf, pos);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.moveType.getValue());
//       this.item.serialize(buf);
//    }

func (pk *SmartGiveCreativeItem) Marshal(io protocol.IO) {
	// TODO: Implement
}
