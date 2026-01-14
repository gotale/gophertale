package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockParticleSets struct{}

func (*UpdateBlockParticleSets) ID() uint32 {
	return IDUpdateBlockParticleSets
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockParticleSets obj = new UpdateBlockParticleSets();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int blockParticleSetsCount = VarInt.peek(buf, pos);
//          if (blockParticleSetsCount < 0) {
//             throw ProtocolException.negativeLength("BlockParticleSets", blockParticleSetsCount);
//          }
//
//          if (blockParticleSetsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockParticleSets", blockParticleSetsCount, 4096000);
//          }
//
//          pos += VarInt.size(blockParticleSetsCount);
//          obj.blockParticleSets = new HashMap<>(blockParticleSetsCount);
//
//          for (int i = 0; i < blockParticleSetsCount; i++) {
//             int keyLen = VarInt.peek(buf, pos);
//             if (keyLen < 0) {
//                throw ProtocolException.negativeLength("key", keyLen);
//             }
//
//             if (keyLen > 4096000) {
//                throw ProtocolException.stringTooLong("key", keyLen, 4096000);
//             }
//
//             int keyVarLen = VarInt.length(buf, pos);
//             String key = PacketIO.readVarString(buf, pos);
//             pos += keyVarLen + keyLen;
//             BlockParticleSet val = BlockParticleSet.deserialize(buf, pos);
//             pos += BlockParticleSet.computeBytesConsumed(buf, pos);
//             if (obj.blockParticleSets.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("blockParticleSets", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.blockParticleSets != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.blockParticleSets != null) {
//          if (this.blockParticleSets.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockParticleSets", this.blockParticleSets.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.blockParticleSets.size());
//
//          for (Entry<String, BlockParticleSet> e : this.blockParticleSets.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateBlockParticleSets) Marshal(io protocol.IO) {
	// TODO: Implement
}
