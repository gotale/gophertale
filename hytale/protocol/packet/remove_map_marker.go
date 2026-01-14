package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type RemoveMapMarker struct{}

func (*RemoveMapMarker) ID() uint32 {
	return IDRemoveMapMarker
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       RemoveMapMarker obj = new RemoveMapMarker();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int markerIdLen = VarInt.peek(buf, pos);
//          if (markerIdLen < 0) {
//             throw ProtocolException.negativeLength("MarkerId", markerIdLen);
//          }
//
//          if (markerIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("MarkerId", markerIdLen, 4096000);
//          }
//
//          int markerIdVarLen = VarInt.length(buf, pos);
//          obj.markerId = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += markerIdVarLen + markerIdLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.markerId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.markerId != null) {
//          PacketIO.writeVarString(buf, this.markerId, 4096000);
//       }
//    }

func (pk *RemoveMapMarker) Marshal(io protocol.IO) {
	// TODO: Implement
}
