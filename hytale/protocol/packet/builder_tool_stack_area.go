package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolStackArea struct{}

func (*BuilderToolStackArea) ID() uint32 {
	return IDBuilderToolStackArea
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolStackArea obj = new BuilderToolStackArea();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          obj.selectionMin = BlockPosition.deserialize(buf, offset + 1);
//       }
//
//       if ((nullBits & 2) != 0) {
//          obj.selectionMax = BlockPosition.deserialize(buf, offset + 13);
//       }
//
//       obj.xNormal = buf.getIntLE(offset + 25);
//       obj.yNormal = buf.getIntLE(offset + 29);
//       obj.zNormal = buf.getIntLE(offset + 33);
//       obj.numStacks = buf.getIntLE(offset + 37);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.selectionMin != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.selectionMax != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.selectionMin != null) {
//          this.selectionMin.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       if (this.selectionMax != null) {
//          this.selectionMax.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       buf.writeIntLE(this.xNormal);
//       buf.writeIntLE(this.yNormal);
//       buf.writeIntLE(this.zNormal);
//       buf.writeIntLE(this.numStacks);
//    }

func (pk *BuilderToolStackArea) Marshal(io protocol.IO) {
	// TODO: Implement
}
