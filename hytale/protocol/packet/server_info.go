package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerInfo struct{}

func (*ServerInfo) ID() uint32 {
	return IDServerInfo
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerInfo obj = new ServerInfo();
//       byte nullBits = buf.getByte(offset);
//       obj.maxPlayers = buf.getIntLE(offset + 1);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 13 + buf.getIntLE(offset + 5);
//          int serverNameLen = VarInt.peek(buf, varPos0);
//          if (serverNameLen < 0) {
//             throw ProtocolException.negativeLength("ServerName", serverNameLen);
//          }
//
//          if (serverNameLen > 4096000) {
//             throw ProtocolException.stringTooLong("ServerName", serverNameLen, 4096000);
//          }
//
//          obj.serverName = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 13 + buf.getIntLE(offset + 9);
//          int motdLen = VarInt.peek(buf, varPos1);
//          if (motdLen < 0) {
//             throw ProtocolException.negativeLength("Motd", motdLen);
//          }
//
//          if (motdLen > 4096000) {
//             throw ProtocolException.stringTooLong("Motd", motdLen, 4096000);
//          }
//
//          obj.motd = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.serverName != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.motd != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.maxPlayers);
//       int serverNameOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int motdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.serverName != null) {
//          buf.setIntLE(serverNameOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.serverName, 4096000);
//       } else {
//          buf.setIntLE(serverNameOffsetSlot, -1);
//       }
//
//       if (this.motd != null) {
//          buf.setIntLE(motdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.motd, 4096000);
//       } else {
//          buf.setIntLE(motdOffsetSlot, -1);
//       }
//    }

func (pk *ServerInfo) Marshal(io protocol.IO) {
	// TODO: Implement
}
