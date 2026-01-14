package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ConnectAccept struct{}

func (*ConnectAccept) ID() uint32 {
	return IDConnectAccept
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ConnectAccept obj = new ConnectAccept();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int passwordChallengeCount = VarInt.peek(buf, pos);
//          if (passwordChallengeCount < 0) {
//             throw ProtocolException.negativeLength("PasswordChallenge", passwordChallengeCount);
//          }
//
//          if (passwordChallengeCount > 64) {
//             throw ProtocolException.arrayTooLong("PasswordChallenge", passwordChallengeCount, 64);
//          }
//
//          int passwordChallengeVarLen = VarInt.size(passwordChallengeCount);
//          if (pos + passwordChallengeVarLen + passwordChallengeCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("PasswordChallenge", pos + passwordChallengeVarLen + passwordChallengeCount * 1, buf.readableBytes());
//          }
//
//          pos += passwordChallengeVarLen;
//          obj.passwordChallenge = new byte[passwordChallengeCount];
//
//          for (int i = 0; i < passwordChallengeCount; i++) {
//             obj.passwordChallenge[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += passwordChallengeCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.passwordChallenge != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.passwordChallenge != null) {
//          if (this.passwordChallenge.length > 64) {
//             throw ProtocolException.arrayTooLong("PasswordChallenge", this.passwordChallenge.length, 64);
//          }
//
//          VarInt.write(buf, this.passwordChallenge.length);
//
//          for (byte item : this.passwordChallenge) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *ConnectAccept) Marshal(io protocol.IO) {
	// TODO: Implement
}
