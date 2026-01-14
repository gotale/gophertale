package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEditorTimeOverride struct{}

func (*UpdateEditorTimeOverride) ID() uint32 {
	return IDUpdateEditorTimeOverride
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEditorTimeOverride obj = new UpdateEditorTimeOverride();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.gameTime = InstantData.deserialize(buf, offset + 1);
//       }
//
//       obj.paused = buf.getByte(offset + 13) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.gameTime != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.gameTime != null) {
//          this.gameTime.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       buf.writeByte(this.paused ? 1 : 0);
//    }

func (pk *UpdateEditorTimeOverride) Marshal(io protocol.IO) {
	// TODO: Implement
}
