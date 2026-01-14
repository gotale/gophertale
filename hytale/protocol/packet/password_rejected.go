package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PasswordRejected struct{}

func (*PasswordRejected) ID() uint32 {
	return IDPasswordRejected
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PasswordRejected obj = new PasswordRejected();
//       byte nullBits = buf.getByte(offset);
//       obj.attemptsRemaining = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          int newChallengeCount = VarInt.peek(buf, pos);
//          if (newChallengeCount < 0) {
//             throw ProtocolException.negativeLength("NewChallenge", newChallengeCount);
//          }
//
//          if (newChallengeCount > 64) {
//             throw ProtocolException.arrayTooLong("NewChallenge", newChallengeCount, 64);
//          }
//
//          int newChallengeVarLen = VarInt.size(newChallengeCount);
//          if (pos + newChallengeVarLen + newChallengeCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("NewChallenge", pos + newChallengeVarLen + newChallengeCount * 1, buf.readableBytes());
//          }
//
//          pos += newChallengeVarLen;
//          obj.newChallenge = new byte[newChallengeCount];
//
//          for (int i = 0; i < newChallengeCount; i++) {
//             obj.newChallenge[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += newChallengeCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.newChallenge != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.attemptsRemaining);
//       if (this.newChallenge != null) {
//          if (this.newChallenge.length > 64) {
//             throw ProtocolException.arrayTooLong("NewChallenge", this.newChallenge.length, 64);
//          }
//
//          VarInt.write(buf, this.newChallenge.length);
//
//          for (byte item : this.newChallenge) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *PasswordRejected) Marshal(io protocol.IO) {
	// TODO: Implement
}
