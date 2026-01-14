package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetEntitySeed struct{}

func (*SetEntitySeed) ID() uint32 {
	return IDSetEntitySeed
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetEntitySeed obj = new SetEntitySeed();
//       obj.entitySeed = buf.getIntLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.entitySeed);
//    }

func (pk *SetEntitySeed) Marshal(io protocol.IO) {
	// TODO: Implement
}
