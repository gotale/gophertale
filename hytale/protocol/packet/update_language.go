package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateLanguage struct{}

func (*UpdateLanguage) ID() uint32 {
	return IDUpdateLanguage
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateLanguage obj = new UpdateLanguage();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int languageLen = VarInt.peek(buf, pos);
//          if (languageLen < 0) {
//             throw ProtocolException.negativeLength("Language", languageLen);
//          }
//
//          if (languageLen > 4096000) {
//             throw ProtocolException.stringTooLong("Language", languageLen, 4096000);
//          }
//
//          int languageVarLen = VarInt.length(buf, pos);
//          obj.language = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += languageVarLen + languageLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.language != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.language != null) {
//          PacketIO.writeVarString(buf, this.language, 4096000);
//       }
//    }

func (pk *UpdateLanguage) Marshal(io protocol.IO) {
	// TODO: Implement
}
