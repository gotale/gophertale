package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type Status struct{}

func (*Status) ID() uint32 {
	return IDStatus
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       Status obj = new Status();
//       byte nullBits = buf.getByte(offset);
//       obj.playerCount = buf.getIntLE(offset + 1);
//       obj.maxPlayers = buf.getIntLE(offset + 5);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 17 + buf.getIntLE(offset + 9);
//          int nameLen = VarInt.peek(buf, varPos0);
//          if (nameLen < 0) {
//             throw ProtocolException.negativeLength("Name", nameLen);
//          }
//
//          if (nameLen > 128) {
//             throw ProtocolException.stringTooLong("Name", nameLen, 128);
//          }
//
//          obj.name = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 17 + buf.getIntLE(offset + 13);
//          int motdLen = VarInt.peek(buf, varPos1);
//          if (motdLen < 0) {
//             throw ProtocolException.negativeLength("Motd", motdLen);
//          }
//
//          if (motdLen > 512) {
//             throw ProtocolException.stringTooLong("Motd", motdLen, 512);
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
//       if (this.name != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.motd != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.playerCount);
//       buf.writeIntLE(this.maxPlayers);
//       int nameOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int motdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.name != null) {
//          buf.setIntLE(nameOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.name, 128);
//       } else {
//          buf.setIntLE(nameOffsetSlot, -1);
//       }
//
//       if (this.motd != null) {
//          buf.setIntLE(motdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.motd, 512);
//       } else {
//          buf.setIntLE(motdOffsetSlot, -1);
//       }
//    }

func (pk *Status) Marshal(io protocol.IO) {
	// TODO: Implement
}
