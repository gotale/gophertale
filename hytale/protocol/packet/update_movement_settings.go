package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateMovementSettings struct{}

func (*UpdateMovementSettings) ID() uint32 {
	return IDUpdateMovementSettings
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateMovementSettings obj = new UpdateMovementSettings();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.movementSettings = MovementSettings.deserialize(buf, offset + 1);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.movementSettings != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.movementSettings != null) {
//          this.movementSettings.serialize(buf);
//       } else {
//          buf.writeZero(251);
//       }
//    }

func (pk *UpdateMovementSettings) Marshal(io protocol.IO) {
	// TODO: Implement
}
