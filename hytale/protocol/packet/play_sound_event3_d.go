package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PlaySoundEvent3D struct{}

func (*PlaySoundEvent3D) ID() uint32 {
	return IDPlaySoundEvent3D
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PlaySoundEvent3D obj = new PlaySoundEvent3D();
//       byte nullBits = buf.getByte(offset);
//       obj.soundEventIndex = buf.getIntLE(offset + 1);
//       obj.category = SoundCategory.fromValue(buf.getByte(offset + 5));
//       if ((nullBits & 1) != 0) {
//          obj.position = Position.deserialize(buf, offset + 6);
//       }
//
//       obj.volumeModifier = buf.getFloatLE(offset + 30);
//       obj.pitchModifier = buf.getFloatLE(offset + 34);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.position != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.soundEventIndex);
//       buf.writeByte(this.category.getValue());
//       if (this.position != null) {
//          this.position.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//
//       buf.writeFloatLE(this.volumeModifier);
//       buf.writeFloatLE(this.pitchModifier);
//    }

func (pk *PlaySoundEvent3D) Marshal(io protocol.IO) {
	// TODO: Implement
}
