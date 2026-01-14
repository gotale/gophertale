package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetChunkEnvironments struct{}

func (*SetChunkEnvironments) ID() uint32 {
	return IDSetChunkEnvironments
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetChunkEnvironments obj = new SetChunkEnvironments();
//       byte nullBits = buf.getByte(offset);
//       obj.x = buf.getIntLE(offset + 1);
//       obj.z = buf.getIntLE(offset + 5);
//       int pos = offset + 9;
//       if ((nullBits & 1) != 0) {
//          int environmentsCount = VarInt.peek(buf, pos);
//          if (environmentsCount < 0) {
//             throw ProtocolException.negativeLength("Environments", environmentsCount);
//          }
//
//          if (environmentsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Environments", environmentsCount, 4096000);
//          }
//
//          int environmentsVarLen = VarInt.size(environmentsCount);
//          if (pos + environmentsVarLen + environmentsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Environments", pos + environmentsVarLen + environmentsCount * 1, buf.readableBytes());
//          }
//
//          pos += environmentsVarLen;
//          obj.environments = new byte[environmentsCount];
//
//          for (int i = 0; i < environmentsCount; i++) {
//             obj.environments[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += environmentsCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.environments != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.z);
//       if (this.environments != null) {
//          if (this.environments.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Environments", this.environments.length, 4096000);
//          }
//
//          VarInt.write(buf, this.environments.length);
//
//          for (byte item : this.environments) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *SetChunkEnvironments) Marshal(io protocol.IO) {
	// TODO: Implement
}
