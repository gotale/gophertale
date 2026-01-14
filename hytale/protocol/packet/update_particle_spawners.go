package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateParticleSpawners struct{}

func (*UpdateParticleSpawners) ID() uint32 {
	return IDUpdateParticleSpawners
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateParticleSpawners obj = new UpdateParticleSpawners();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 10 + buf.getIntLE(offset + 2);
//          int particleSpawnersCount = VarInt.peek(buf, varPos0);
//          if (particleSpawnersCount < 0) {
//             throw ProtocolException.negativeLength("ParticleSpawners", particleSpawnersCount);
//          }
//
//          if (particleSpawnersCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ParticleSpawners", particleSpawnersCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          obj.particleSpawners = new HashMap<>(particleSpawnersCount);
//          int dictPos = varPos0 + varIntLen;
//
//          for (int i = 0; i < particleSpawnersCount; i++) {
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
//             ParticleSpawner val = ParticleSpawner.deserialize(buf, dictPos);
//             dictPos += ParticleSpawner.computeBytesConsumed(buf, dictPos);
//             if (obj.particleSpawners.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("particleSpawners", key);
//             }
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 10 + buf.getIntLE(offset + 6);
//          int removedParticleSpawnersCount = VarInt.peek(buf, varPos1);
//          if (removedParticleSpawnersCount < 0) {
//             throw ProtocolException.negativeLength("RemovedParticleSpawners", removedParticleSpawnersCount);
//          }
//
//          if (removedParticleSpawnersCount > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedParticleSpawners", removedParticleSpawnersCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + removedParticleSpawnersCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("RemovedParticleSpawners", varPos1 + varIntLen + removedParticleSpawnersCount * 1, buf.readableBytes());
//          }
//
//          obj.removedParticleSpawners = new String[removedParticleSpawnersCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < removedParticleSpawnersCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("removedParticleSpawners[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("removedParticleSpawners[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.removedParticleSpawners[i] = PacketIO.readVarString(buf, elemPos);
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
//       if (this.particleSpawners != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.removedParticleSpawners != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       int particleSpawnersOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int removedParticleSpawnersOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.particleSpawners != null) {
//          buf.setIntLE(particleSpawnersOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.particleSpawners.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ParticleSpawners", this.particleSpawners.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.particleSpawners.size());
//
//          for (Entry<String, ParticleSpawner> e : this.particleSpawners.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       } else {
//          buf.setIntLE(particleSpawnersOffsetSlot, -1);
//       }
//
//       if (this.removedParticleSpawners != null) {
//          buf.setIntLE(removedParticleSpawnersOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.removedParticleSpawners.length > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedParticleSpawners", this.removedParticleSpawners.length, 4096000);
//          }
//
//          VarInt.write(buf, this.removedParticleSpawners.length);
//
//          for (String item : this.removedParticleSpawners) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(removedParticleSpawnersOffsetSlot, -1);
//       }
//    }

func (pk *UpdateParticleSpawners) Marshal(io protocol.IO) {
	// TODO: Implement
}
