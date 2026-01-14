package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type TeleportToWorldMapMarker struct{}

func (*TeleportToWorldMapMarker) ID() uint32 {
	return IDTeleportToWorldMapMarker
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       TeleportToWorldMapMarker obj = new TeleportToWorldMapMarker();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int idLen = VarInt.peek(buf, pos);
//          if (idLen < 0) {
//             throw ProtocolException.negativeLength("Id", idLen);
//          }
//
//          if (idLen > 4096000) {
//             throw ProtocolException.stringTooLong("Id", idLen, 4096000);
//          }
//
//          int idVarLen = VarInt.length(buf, pos);
//          obj.id = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += idVarLen + idLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.id != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.id != null) {
//          PacketIO.writeVarString(buf, this.id, 4096000);
//       }
//    }

func (pk *TeleportToWorldMapMarker) Marshal(io protocol.IO) {
	// TODO: Implement
}
