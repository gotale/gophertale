package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AuthToken struct{}

func (*AuthToken) ID() uint32 {
	return IDAuthToken
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AuthToken obj = new AuthToken();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          int accessTokenLen = VarInt.peek(buf, varPos0);
//          if (accessTokenLen < 0) {
//             throw ProtocolException.negativeLength("AccessToken", accessTokenLen);
//          }
//
//          if (accessTokenLen > 8192) {
//             throw ProtocolException.stringTooLong("AccessToken", accessTokenLen, 8192);
//          }
//
//          obj.accessToken = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int serverAuthorizationGrantLen = VarInt.peek(buf, varPos1);
//          if (serverAuthorizationGrantLen < 0) {
//             throw ProtocolException.negativeLength("ServerAuthorizationGrant", serverAuthorizationGrantLen);
//          }
//
//          if (serverAuthorizationGrantLen > 4096) {
//             throw ProtocolException.stringTooLong("ServerAuthorizationGrant", serverAuthorizationGrantLen, 4096);
//          }
//
//          obj.serverAuthorizationGrant = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.accessToken != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.serverAuthorizationGrant != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int accessTokenOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int serverAuthorizationGrantOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.accessToken != null) {
//          buf.setIntLE(accessTokenOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.accessToken, 8192);
//       } else {
//          buf.setIntLE(accessTokenOffsetSlot, -1);
//       }
//
//       if (this.serverAuthorizationGrant != null) {
//          buf.setIntLE(serverAuthorizationGrantOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.serverAuthorizationGrant, 4096);
//       } else {
//          buf.setIntLE(serverAuthorizationGrantOffsetSlot, -1);
//       }
//    }

func (pk *AuthToken) Marshal(io protocol.IO) {
	// TODO: Implement
}
