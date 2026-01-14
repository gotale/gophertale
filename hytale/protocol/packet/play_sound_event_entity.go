package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PlaySoundEventEntity struct{}

func (*PlaySoundEventEntity) ID() uint32 {
	return IDPlaySoundEventEntity
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PlaySoundEventEntity obj = new PlaySoundEventEntity();
//       obj.soundEventIndex = buf.getIntLE(offset + 0);
//       obj.networkId = buf.getIntLE(offset + 4);
//       obj.volumeModifier = buf.getFloatLE(offset + 8);
//       obj.pitchModifier = buf.getFloatLE(offset + 12);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.soundEventIndex);
//       buf.writeIntLE(this.networkId);
//       buf.writeFloatLE(this.volumeModifier);
//       buf.writeFloatLE(this.pitchModifier);
//    }

func (pk *PlaySoundEventEntity) Marshal(io protocol.IO) {
	// TODO: Implement
}
