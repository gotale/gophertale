package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type TeleportToWorldMapPosition struct{}

func (*TeleportToWorldMapPosition) ID() uint32 {
	return IDTeleportToWorldMapPosition
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       TeleportToWorldMapPosition obj = new TeleportToWorldMapPosition();
//       obj.x = buf.getIntLE(offset + 0);
//       obj.y = buf.getIntLE(offset + 4);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//    }

func (pk *TeleportToWorldMapPosition) Marshal(io protocol.IO) {
	// TODO: Implement
}
