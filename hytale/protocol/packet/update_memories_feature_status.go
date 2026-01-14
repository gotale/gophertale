package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateMemoriesFeatureStatus struct{}

func (*UpdateMemoriesFeatureStatus) ID() uint32 {
	return IDUpdateMemoriesFeatureStatus
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateMemoriesFeatureStatus obj = new UpdateMemoriesFeatureStatus();
//       obj.isFeatureUnlocked = buf.getByte(offset + 0) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.isFeatureUnlocked ? 1 : 0);
//    }

func (pk *UpdateMemoriesFeatureStatus) Marshal(io protocol.IO) {
	// TODO: Implement
}
