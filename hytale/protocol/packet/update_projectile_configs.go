package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateProjectileConfigs struct{}

func (*UpdateProjectileConfigs) ID() uint32 {
	return IDUpdateProjectileConfigs
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateProjectileConfigs obj = new UpdateProjectileConfigs();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 10 + buf.getIntLE(offset + 2);
//          int configsCount = VarInt.peek(buf, varPos0);
//          if (configsCount < 0) {
//             throw ProtocolException.negativeLength("Configs", configsCount);
//          }
//
//          if (configsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Configs", configsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          obj.configs = new HashMap<>(configsCount);
//          int dictPos = varPos0 + varIntLen;
//
//          for (int i = 0; i < configsCount; i++) {
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
//             ProjectileConfig val = ProjectileConfig.deserialize(buf, dictPos);
//             dictPos += ProjectileConfig.computeBytesConsumed(buf, dictPos);
//             if (obj.configs.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("configs", key);
//             }
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 10 + buf.getIntLE(offset + 6);
//          int removedConfigsCount = VarInt.peek(buf, varPos1);
//          if (removedConfigsCount < 0) {
//             throw ProtocolException.negativeLength("RemovedConfigs", removedConfigsCount);
//          }
//
//          if (removedConfigsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedConfigs", removedConfigsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + removedConfigsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("RemovedConfigs", varPos1 + varIntLen + removedConfigsCount * 1, buf.readableBytes());
//          }
//
//          obj.removedConfigs = new String[removedConfigsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < removedConfigsCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("removedConfigs[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("removedConfigs[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.removedConfigs[i] = PacketIO.readVarString(buf, elemPos);
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
//       if (this.configs != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.removedConfigs != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       int configsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int removedConfigsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.configs != null) {
//          buf.setIntLE(configsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.configs.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Configs", this.configs.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.configs.size());
//
//          for (Entry<String, ProjectileConfig> e : this.configs.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       } else {
//          buf.setIntLE(configsOffsetSlot, -1);
//       }
//
//       if (this.removedConfigs != null) {
//          buf.setIntLE(removedConfigsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.removedConfigs.length > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedConfigs", this.removedConfigs.length, 4096000);
//          }
//
//          VarInt.write(buf, this.removedConfigs.length);
//
//          for (String item : this.removedConfigs) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(removedConfigsOffsetSlot, -1);
//       }
//    }

func (pk *UpdateProjectileConfigs) Marshal(io protocol.IO) {
	// TODO: Implement
}
