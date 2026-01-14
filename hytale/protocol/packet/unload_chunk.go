package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UnloadChunk struct{}

func (*UnloadChunk) ID() uint32 {
	return IDUnloadChunk
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UnloadChunk obj = new UnloadChunk();
//       obj.chunkX = buf.getIntLE(offset + 0);
//       obj.chunkZ = buf.getIntLE(offset + 4);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.chunkX);
//       buf.writeIntLE(this.chunkZ);
//    }

func (pk *UnloadChunk) Marshal(io protocol.IO) {
	// TODO: Implement
}
