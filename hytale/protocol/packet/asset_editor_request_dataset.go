package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorRequestDataset struct{}

func (*AssetEditorRequestDataset) ID() uint32 {
	return IDAssetEditorRequestDataset
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorRequestDataset obj = new AssetEditorRequestDataset();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int nameLen = VarInt.peek(buf, pos);
//          if (nameLen < 0) {
//             throw ProtocolException.negativeLength("Name", nameLen);
//          }
//
//          if (nameLen > 4096000) {
//             throw ProtocolException.stringTooLong("Name", nameLen, 4096000);
//          }
//
//          int nameVarLen = VarInt.length(buf, pos);
//          obj.name = PacketIO.readVarString(buf, pos, PacketIO.UTF8);
//          pos += nameVarLen + nameLen;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.name != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.name != null) {
//          PacketIO.writeVarString(buf, this.name, 4096000);
//       }
//    }

func (pk *AssetEditorRequestDataset) Marshal(io protocol.IO) {
	// TODO: Implement
}
