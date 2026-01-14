package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorActivateButton struct{}

func (*AssetEditorActivateButton) ID() uint32 {
	return IDAssetEditorActivateButton
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorActivateButton obj = new AssetEditorActivateButton();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int buttonIdLen = VarInt.peek(buf, pos);
//          if (buttonIdLen < 0) {
//             throw ProtocolException.negativeLength("ButtonId", buttonIdLen);
//          }
//
//          if (buttonIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("ButtonId", buttonIdLen, 4096000);
//          }
//
//          int buttonIdVarLen = VarInt.length(buf, pos);
//          obj.buttonId = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += buttonIdVarLen + buttonIdLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.buttonId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.buttonId != null) {
//          PacketIO.writeVarString(buf, this.buttonId, 4096000);
//       }
//    }

func (pk *AssetEditorActivateButton) Marshal(io protocol.IO) {
	// TODO: Implement
}
