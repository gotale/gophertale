package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClientMovement struct{}

func (*ClientMovement) ID() uint32 {
	return IDClientMovement
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ClientMovement obj = new ClientMovement();
//       byte[] nullBits = PacketIO.readBytes(buf, offset, 2);
//       if ((nullBits[0] & 1) != 0) {
//          obj.movementStates = MovementStates.deserialize(buf, offset + 2);
//       }
//
//       if ((nullBits[0] & 2) != 0) {
//          obj.relativePosition = HalfFloatPosition.deserialize(buf, offset + 24);
//       }
//
//       if ((nullBits[0] & 4) != 0) {
//          obj.absolutePosition = Position.deserialize(buf, offset + 30);
//       }
//
//       if ((nullBits[0] & 8) != 0) {
//          obj.bodyOrientation = Direction.deserialize(buf, offset + 54);
//       }
//
//       if ((nullBits[0] & 16) != 0) {
//          obj.lookOrientation = Direction.deserialize(buf, offset + 66);
//       }
//
//       if ((nullBits[0] & 32) != 0) {
//          obj.teleportAck = TeleportAck.deserialize(buf, offset + 78);
//       }
//
//       if ((nullBits[0] & 64) != 0) {
//          obj.wishMovement = Position.deserialize(buf, offset + 79);
//       }
//
//       if ((nullBits[0] & 128) != 0) {
//          obj.velocity = Vector3d.deserialize(buf, offset + 103);
//       }
//
//       obj.mountedTo = buf.getIntLE(offset + 127);
//       if ((nullBits[1] & 1) != 0) {
//          obj.riderMovementStates = MovementStates.deserialize(buf, offset + 131);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte[] nullBits = new byte[2];
//       if (this.movementStates != null) {
//          nullBits[0] = (byte)(nullBits[0] | 1);
//       }
//
//       if (this.relativePosition != null) {
//          nullBits[0] = (byte)(nullBits[0] | 2);
//       }
//
//       if (this.absolutePosition != null) {
//          nullBits[0] = (byte)(nullBits[0] | 4);
//       }
//
//       if (this.bodyOrientation != null) {
//          nullBits[0] = (byte)(nullBits[0] | 8);
//       }
//
//       if (this.lookOrientation != null) {
//          nullBits[0] = (byte)(nullBits[0] | 16);
//       }
//
//       if (this.teleportAck != null) {
//          nullBits[0] = (byte)(nullBits[0] | 32);
//       }
//
//       if (this.wishMovement != null) {
//          nullBits[0] = (byte)(nullBits[0] | 64);
//       }
//
//       if (this.velocity != null) {
//          nullBits[0] = (byte)(nullBits[0] | 128);
//       }
//
//       if (this.riderMovementStates != null) {
//          nullBits[1] = (byte)(nullBits[1] | 1);
//       }
//
//       buf.writeBytes(nullBits);
//       if (this.movementStates != null) {
//          this.movementStates.serialize(buf);
//       } else {
//          buf.writeZero(22);
//       }
//
//       if (this.relativePosition != null) {
//          this.relativePosition.serialize(buf);
//       } else {
//          buf.writeZero(6);
//       }
//
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
//       if (this.lookOrientation != null) {
//          this.lookOrientation.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       if (this.teleportAck != null) {
//          this.teleportAck.serialize(buf);
//       } else {
//          buf.writeZero(1);
//       }
//
//       if (this.wishMovement != null) {
//          this.wishMovement.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//
//       if (this.velocity != null) {
//          this.velocity.serialize(buf);
//       } else {
//          buf.writeZero(24);
//       }
//
//       buf.writeIntLE(this.mountedTo);
//       if (this.riderMovementStates != null) {
//          this.riderMovementStates.serialize(buf);
//       } else {
//          buf.writeZero(22);
//       }
//    }

func (pk *ClientMovement) Marshal(io protocol.IO) {
	// TODO: Implement
}
