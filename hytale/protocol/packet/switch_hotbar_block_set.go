package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SwitchHotbarBlockSet struct{}

func (*SwitchHotbarBlockSet) ID() uint32 {
	return IDSwitchHotbarBlockSet
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SwitchHotbarBlockSet obj = new SwitchHotbarBlockSet();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int itemIdLen = VarInt.peek(buf, pos);
//          if (itemIdLen < 0) {
//             throw ProtocolException.negativeLength("ItemId", itemIdLen);
//          }
//
//          if (itemIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("ItemId", itemIdLen, 4096000);
//          }
//
//          int itemIdVarLen = VarInt.length(buf, pos);
//          obj.itemId = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += itemIdVarLen + itemIdLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.itemId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.itemId != null) {
//          PacketIO.writeVarString(buf, this.itemId, 4096000);
//       }
//    }

func (pk *SwitchHotbarBlockSet) Marshal(io protocol.IO) {
	// TODO: Implement
}
