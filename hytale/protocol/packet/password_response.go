package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PasswordResponse struct{}

func (*PasswordResponse) ID() uint32 {
	return IDPasswordResponse
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PasswordResponse obj = new PasswordResponse();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int hashCount = VarInt.peek(buf, pos);
//          if (hashCount < 0) {
//             throw ProtocolException.negativeLength("Hash", hashCount);
//          }
//
//          if (hashCount > 64) {
//             throw ProtocolException.arrayTooLong("Hash", hashCount, 64);
//          }
//
//          int hashVarLen = VarInt.size(hashCount);
//          if (pos + hashVarLen + hashCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Hash", pos + hashVarLen + hashCount * 1, buf.readableBytes());
//          }
//
//          pos += hashVarLen;
//          obj.hash = new byte[hashCount];
//
//          for (int i = 0; i < hashCount; i++) {
//             obj.hash[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += hashCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.hash != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.hash != null) {
//          if (this.hash.length > 64) {
//             throw ProtocolException.arrayTooLong("Hash", this.hash.length, 64);
//          }
//
//          VarInt.write(buf, this.hash.length);
//
//          for (byte item : this.hash) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *PasswordResponse) Marshal(io protocol.IO) {
	// TODO: Implement
}
