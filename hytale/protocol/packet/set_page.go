package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetPage struct{}

func (*SetPage) ID() uint32 {
	return IDSetPage
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetPage obj = new SetPage();
//       obj.page = Page.fromValue(buf.getByte(offset + 0));
//       obj.canCloseThroughInteraction = buf.getByte(offset + 1) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.page.getValue());
//       buf.writeByte(this.canCloseThroughInteraction ? 1 : 0);
//    }

func (pk *SetPage) Marshal(io protocol.IO) {
	// TODO: Implement
}
