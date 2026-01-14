package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEnvironmentMusic struct{}

func (*UpdateEnvironmentMusic) ID() uint32 {
	return IDUpdateEnvironmentMusic
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEnvironmentMusic obj = new UpdateEnvironmentMusic();
//       obj.environmentIndex = buf.getIntLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.environmentIndex);
//    }

func (pk *UpdateEnvironmentMusic) Marshal(io protocol.IO) {
	// TODO: Implement
}
