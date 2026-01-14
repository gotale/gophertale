package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AuthGrant struct{}

func (*AuthGrant) ID() uint32 {
	return IDAuthGrant
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AuthGrant obj = new AuthGrant();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          int authorizationGrantLen = VarInt.peek(buf, varPos0);
//          if (authorizationGrantLen < 0) {
//             throw ProtocolException.negativeLength("AuthorizationGrant", authorizationGrantLen);
//          }
//
//          if (authorizationGrantLen > 4096) {
//             throw ProtocolException.stringTooLong("AuthorizationGrant", authorizationGrantLen, 4096);
//          }
//
//          obj.authorizationGrant = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int serverIdentityTokenLen = VarInt.peek(buf, varPos1);
//          if (serverIdentityTokenLen < 0) {
//             throw ProtocolException.negativeLength("ServerIdentityToken", serverIdentityTokenLen);
//          }
//
//          if (serverIdentityTokenLen > 8192) {
//             throw ProtocolException.stringTooLong("ServerIdentityToken", serverIdentityTokenLen, 8192);
//          }
//
//          obj.serverIdentityToken = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.authorizationGrant != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.serverIdentityToken != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int authorizationGrantOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int serverIdentityTokenOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.authorizationGrant != null) {
//          buf.setIntLE(authorizationGrantOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.authorizationGrant, 4096);
//       } else {
//          buf.setIntLE(authorizationGrantOffsetSlot, -1);
//       }
//
//       if (this.serverIdentityToken != null) {
//          buf.setIntLE(serverIdentityTokenOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.serverIdentityToken, 8192);
//       } else {
//          buf.setIntLE(serverIdentityTokenOffsetSlot, -1);
//       }
//    }

func (pk *AuthGrant) Marshal(io protocol.IO) {
	// TODO: Implement
}
