package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ClientPlaceBlock struct{}

func (*ClientPlaceBlock) ID() uint32 {
	return IDClientPlaceBlock
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ClientPlaceBlock obj = new ClientPlaceBlock();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.position = BlockPosition.deserialize(buf, offset + 1);
//       }
//
//       if ((nullBits & 2) != 0) {
//          obj.rotation = BlockRotation.deserialize(buf, offset + 13);
//       }
//
//       obj.placedBlockId = buf.getIntLE(offset + 16);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.position != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.rotation != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.position != null) {
//          this.position.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       if (this.rotation != null) {
//          this.rotation.serialize(buf);
//       } else {
//          buf.writeZero(3);
//       }
//
//       buf.writeIntLE(this.placedBlockId);
//    }

func (pk *ClientPlaceBlock) Marshal(io protocol.IO) {
	// TODO: Implement
}
