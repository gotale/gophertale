package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SpawnParticleSystem struct{}

func (*SpawnParticleSystem) ID() uint32 {
	return IDSpawnParticleSystem
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SpawnParticleSystem obj = new SpawnParticleSystem();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 2) != 0) {
//          obj.position = Position.deserialize(buf, offset + 1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          obj.rotation = Direction.deserialize(buf, offset + 25);
//       }
//
//       obj.scale = buf.getFloatLE(offset + 37);
//       if ((nullBits & 8) != 0) {
//          obj.color = Color.deserialize(buf, offset + 41);
//       }
//
//       int pos = offset + 44;
//       if ((nullBits & 1) != 0) {
//          int particleSystemIdLen = VarInt.peek(buf, pos);
//          if (particleSystemIdLen < 0) {
//             throw ProtocolException.negativeLength("ParticleSystemId", particleSystemIdLen);
//          }
//
//          if (particleSystemIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("ParticleSystemId", particleSystemIdLen, 4096000);
//          }
//
//          int particleSystemIdVarLen = VarInt.length(buf, pos);
//          obj.particleSystemId = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += particleSystemIdVarLen + particleSystemIdLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.particleSystemId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.position != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.rotation != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       if (this.color != null) {
//          nullBits = (byte)(nullBits | 8);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.position != null) {
//          this.position.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//
//       if (this.rotation != null) {
//          this.rotation.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       buf.writeFloatLE(this.scale);
//       if (this.color != null) {
//          this.color.serialize(buf);
//       } else {
//          buf.writeZero(3);
//       }
//
//       if (this.particleSystemId != null) {
//          PacketIO.writeVarString(buf, this.particleSystemId, 4096000);
//       }
//    }

func (pk *SpawnParticleSystem) Marshal(io protocol.IO) {
	// TODO: Implement
}
