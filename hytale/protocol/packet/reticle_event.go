package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ReticleEvent struct{}

func (*ReticleEvent) ID() uint32 {
	return IDReticleEvent
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ReticleEvent obj = new ReticleEvent();
//       obj.eventIndex = buf.getIntLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.eventIndex);
//    }

func (pk *ReticleEvent) Marshal(io protocol.IO) {
	// TODO: Implement
}
