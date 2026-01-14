package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSetEntityLight struct{}

func (*BuilderToolSetEntityLight) ID() uint32 {
	return IDBuilderToolSetEntityLight
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSetEntityLight obj = new BuilderToolSetEntityLight();
//       byte nullBits = buf.getByte(offset);
//       obj.entityId = buf.getIntLE(offset + 1);
//       if ((nullBits & 1) != 0) {
//          obj.light = ColorLight.deserialize(buf, offset + 5);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.light != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.entityId);
//       if (this.light != null) {
//          this.light.serialize(buf);
//       } else {
//          buf.writeZero(4);
//       }
//    }

func (pk *BuilderToolSetEntityLight) Marshal(io protocol.IO) {
	// TODO: Implement
}
