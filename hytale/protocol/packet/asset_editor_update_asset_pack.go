package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUpdateAssetPack struct{}

func (*AssetEditorUpdateAssetPack) ID() uint32 {
	return IDAssetEditorUpdateAssetPack
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUpdateAssetPack obj = new AssetEditorUpdateAssetPack();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          int idLen = VarInt.peek(buf, varPos0);
//          if (idLen < 0) {
//             throw ProtocolException.negativeLength("Id", idLen);
//          }
//
//          if (idLen > 4096000) {
//             throw ProtocolException.stringTooLong("Id", idLen, 4096000);
//          }
//
//          obj.id = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          obj.manifest = AssetPackManifest.deserialize(buf, varPos1);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.id != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.manifest != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int idOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int manifestOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.id != null) {
//          buf.setIntLE(idOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.id, 4096000);
//       } else {
//          buf.setIntLE(idOffsetSlot, -1);
//       }
//
//       if (this.manifest != null) {
//          buf.setIntLE(manifestOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.manifest.serialize(buf);
//       } else {
//          buf.setIntLE(manifestOffsetSlot, -1);
//       }
//    }

func (pk *AssetEditorUpdateAssetPack) Marshal(io protocol.IO) {
	// TODO: Implement
}
