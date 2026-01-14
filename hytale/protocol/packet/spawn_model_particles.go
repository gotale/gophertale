package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SpawnModelParticles struct{}

func (*SpawnModelParticles) ID() uint32 {
	return IDSpawnModelParticles
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SpawnModelParticles obj = new SpawnModelParticles();
//       byte nullBits = buf.getByte(offset);
//       obj.entityId = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          int modelParticlesCount = VarInt.peek(buf, pos);
//          if (modelParticlesCount < 0) {
//             throw ProtocolException.negativeLength("ModelParticles", modelParticlesCount);
//          }
//
//          if (modelParticlesCount > 4096000) {
//             throw ProtocolException.arrayTooLong("ModelParticles", modelParticlesCount, 4096000);
//          }
//
//          int modelParticlesVarLen = VarInt.size(modelParticlesCount);
//          if (pos + modelParticlesVarLen + modelParticlesCount * 34L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("ModelParticles", pos + modelParticlesVarLen + modelParticlesCount * 34, buf.readableBytes());
//          }
//
//          pos += modelParticlesVarLen;
//          obj.modelParticles = new ModelParticle[modelParticlesCount];
//
//          for (int i = 0; i < modelParticlesCount; i++) {
//             obj.modelParticles[i] = ModelParticle.deserialize(buf, pos);
//             pos += ModelParticle.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.modelParticles != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.entityId);
//       if (this.modelParticles != null) {
//          if (this.modelParticles.length > 4096000) {
//             throw ProtocolException.arrayTooLong("ModelParticles", this.modelParticles.length, 4096000);
//          }
//
//          VarInt.write(buf, this.modelParticles.length);
//
//          for (ModelParticle item : this.modelParticles) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *SpawnModelParticles) Marshal(io protocol.IO) {
	// TODO: Implement
}
