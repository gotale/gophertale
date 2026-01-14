package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUndoRedoReply struct{}

func (*AssetEditorUndoRedoReply) ID() uint32 {
	return IDAssetEditorUndoRedoReply
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUndoRedoReply obj = new AssetEditorUndoRedoReply();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          obj.command = JsonUpdateCommand.deserialize(buf, pos);
//          pos += JsonUpdateCommand.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.command != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       if (this.command != null) {
//          this.command.serialize(buf);
//       }
//    }

func (pk *AssetEditorUndoRedoReply) Marshal(io protocol.IO) {
	// TODO: Implement
}
