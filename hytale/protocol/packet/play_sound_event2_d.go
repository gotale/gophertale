package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PlaySoundEvent2D struct{}

func (*PlaySoundEvent2D) ID() uint32 {
	return IDPlaySoundEvent2D
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PlaySoundEvent2D obj = new PlaySoundEvent2D();
//       obj.soundEventIndex = buf.getIntLE(offset + 0);
//       obj.category = SoundCategory.fromValue(buf.getByte(offset + 4));
//       obj.volumeModifier = buf.getFloatLE(offset + 5);
//       obj.pitchModifier = buf.getFloatLE(offset + 9);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.soundEventIndex);
//       buf.writeByte(this.category.getValue());
//       buf.writeFloatLE(this.volumeModifier);
//       buf.writeFloatLE(this.pitchModifier);
//    }

func (pk *PlaySoundEvent2D) Marshal(io protocol.IO) {
	// TODO: Implement
}
