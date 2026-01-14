package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SpawnBlockParticleSystem struct{}

func (*SpawnBlockParticleSystem) ID() uint32 {
	return IDSpawnBlockParticleSystem
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SpawnBlockParticleSystem obj = new SpawnBlockParticleSystem();
//       byte nullBits = buf.getByte(offset);
//       obj.blockId = buf.getIntLE(offset + 1);
//       obj.particleType = BlockParticleEvent.fromValue(buf.getByte(offset + 5));
//       if ((nullBits & 1) != 0) {
//          obj.position = Position.deserialize(buf, offset + 6);
//       }
//
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
//       buf.writeIntLE(this.blockId);
//       buf.writeByte(this.particleType.getValue());
//       if (this.position != null) {
//          this.position.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//    }

func (pk *SpawnBlockParticleSystem) Marshal(io protocol.IO) {
	// TODO: Implement
}
