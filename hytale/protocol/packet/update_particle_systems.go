package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateParticleSystems struct{}

func (*UpdateParticleSystems) ID() uint32 {
	return IDUpdateParticleSystems
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateParticleSystems obj = new UpdateParticleSystems();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 10 + buf.getIntLE(offset + 2);
//          int particleSystemsCount = VarInt.peek(buf, varPos0);
//          if (particleSystemsCount < 0) {
//             throw ProtocolException.negativeLength("ParticleSystems", particleSystemsCount);
//          }
//
//          if (particleSystemsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ParticleSystems", particleSystemsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          obj.particleSystems = new HashMap<>(particleSystemsCount);
//          int dictPos = varPos0 + varIntLen;
//
//          for (int i = 0; i < particleSystemsCount; i++) {
//             int keyLen = VarInt.peek(buf, dictPos);
//             if (keyLen < 0) {
//                throw ProtocolException.negativeLength("key", keyLen);
//             }
//
//             if (keyLen > 4096000) {
//                throw ProtocolException.stringTooLong("key", keyLen, 4096000);
//             }
//
//             int keyVarLen = VarInt.length(buf, dictPos);
//             String key = PacketIO.readVarString(buf, dictPos);
//             dictPos += keyVarLen + keyLen;
//             ParticleSystem val = ParticleSystem.deserialize(buf, dictPos);
//             dictPos += ParticleSystem.computeBytesConsumed(buf, dictPos);
//             if (obj.particleSystems.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("particleSystems", key);
//             }
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 10 + buf.getIntLE(offset + 6);
//          int removedParticleSystemsCount = VarInt.peek(buf, varPos1);
//          if (removedParticleSystemsCount < 0) {
//             throw ProtocolException.negativeLength("RemovedParticleSystems", removedParticleSystemsCount);
//          }
//
//          if (removedParticleSystemsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedParticleSystems", removedParticleSystemsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + removedParticleSystemsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("RemovedParticleSystems", varPos1 + varIntLen + removedParticleSystemsCount * 1, buf.readableBytes());
//          }
//
//          obj.removedParticleSystems = new String[removedParticleSystemsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < removedParticleSystemsCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("removedParticleSystems[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("removedParticleSystems[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.removedParticleSystems[i] = PacketIO.readVarString(buf, elemPos);
//             elemPos += strVarLen + strLen;
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.particleSystems != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.removedParticleSystems != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       int particleSystemsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int removedParticleSystemsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.particleSystems != null) {
//          buf.setIntLE(particleSystemsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.particleSystems.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ParticleSystems", this.particleSystems.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.particleSystems.size());
//
//          for (Entry<String, ParticleSystem> e : this.particleSystems.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       } else {
//          buf.setIntLE(particleSystemsOffsetSlot, -1);
//       }
//
//       if (this.removedParticleSystems != null) {
//          buf.setIntLE(removedParticleSystemsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.removedParticleSystems.length > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedParticleSystems", this.removedParticleSystems.length, 4096000);
//          }
//
//          VarInt.write(buf, this.removedParticleSystems.length);
//
//          for (String item : this.removedParticleSystems) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(removedParticleSystemsOffsetSlot, -1);
//       }
//    }

func (pk *UpdateParticleSystems) Marshal(io protocol.IO) {
	// TODO: Implement
}
