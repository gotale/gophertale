package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type MountMovement struct{}

func (*MountMovement) ID() uint32 {
	return IDMountMovement
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       MountMovement obj = new MountMovement();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.absolutePosition = Position.deserialize(buf, offset + 1);
//       }
//
//       if ((nullBits & 2) != 0) {
//          obj.bodyOrientation = Direction.deserialize(buf, offset + 25);
//       }
//
//       if ((nullBits & 4) != 0) {
//          obj.movementStates = MovementStates.deserialize(buf, offset + 37);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.absolutePosition != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.bodyOrientation != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.movementStates != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.absolutePosition != null) {
//          this.absolutePosition.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//
//       if (this.bodyOrientation != null) {
//          this.bodyOrientation.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       if (this.movementStates != null) {
//          this.movementStates.serialize(buf);
//       } else {
//          buf.writeZero(22);
//       }
//    }

func (pk *MountMovement) Marshal(io protocol.IO) {
	// TODO: Implement
}
