package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerAuthToken struct{}

func (*ServerAuthToken) ID() uint32 {
	return IDServerAuthToken
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerAuthToken obj = new ServerAuthToken();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          int serverAccessTokenLen = VarInt.peek(buf, varPos0);
//          if (serverAccessTokenLen < 0) {
//             throw ProtocolException.negativeLength("ServerAccessToken", serverAccessTokenLen);
//          }
//
//          if (serverAccessTokenLen > 8192) {
//             throw ProtocolException.stringTooLong("ServerAccessToken", serverAccessTokenLen, 8192);
//          }
//
//          obj.serverAccessToken = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int passwordChallengeCount = VarInt.peek(buf, varPos1);
//          if (passwordChallengeCount < 0) {
//             throw ProtocolException.negativeLength("PasswordChallenge", passwordChallengeCount);
//          }
//
//          if (passwordChallengeCount > 64) {
//             throw ProtocolException.arrayTooLong("PasswordChallenge", passwordChallengeCount, 64);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + passwordChallengeCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("PasswordChallenge", varPos1 + varIntLen + passwordChallengeCount * 1, buf.readableBytes());
//          }
//
//          obj.passwordChallenge = new byte[passwordChallengeCount];
//
//          for (int i = 0; i < passwordChallengeCount; i++) {
//             obj.passwordChallenge[i] = buf.getByte(varPos1 + varIntLen + i * 1);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.serverAccessToken != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.passwordChallenge != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int serverAccessTokenOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int passwordChallengeOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.serverAccessToken != null) {
//          buf.setIntLE(serverAccessTokenOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.serverAccessToken, 8192);
//       } else {
//          buf.setIntLE(serverAccessTokenOffsetSlot, -1);
//       }
//
//       if (this.passwordChallenge != null) {
//          buf.setIntLE(passwordChallengeOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.passwordChallenge.length > 64) {
//             throw ProtocolException.arrayTooLong("PasswordChallenge", this.passwordChallenge.length, 64);
//          }
//
//          VarInt.write(buf, this.passwordChallenge.length);
//
//          for (byte item : this.passwordChallenge) {
//             buf.writeByte(item);
//          }
//       } else {
//          buf.setIntLE(passwordChallengeOffsetSlot, -1);
//       }
//    }

func (pk *ServerAuthToken) Marshal(io protocol.IO) {
	// TODO: Implement
}
