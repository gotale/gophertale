package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSetEntityTransform struct{}

func (*BuilderToolSetEntityTransform) ID() uint32 {
	return IDBuilderToolSetEntityTransform
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSetEntityTransform obj = new BuilderToolSetEntityTransform();
//       byte nullBits = buf.getByte(offset);
//       obj.entityId = buf.getIntLE(offset + 1);
//       if ((nullBits & 1) != 0) {
//          obj.modelTransform = ModelTransform.deserialize(buf, offset + 5);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.modelTransform != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.entityId);
//       if (this.modelTransform != null) {
//          this.modelTransform.serialize(buf);
//       } else {
//          buf.writeZero(49);
//       }
//    }

func (pk *BuilderToolSetEntityTransform) Marshal(io protocol.IO) {
	// TODO: Implement
}
