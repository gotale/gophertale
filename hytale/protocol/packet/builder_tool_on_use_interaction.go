package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolOnUseInteraction struct{}

func (*BuilderToolOnUseInteraction) ID() uint32 {
	return IDBuilderToolOnUseInteraction
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolOnUseInteraction obj = new BuilderToolOnUseInteraction();
//       obj.type = InteractionType.fromValue(buf.getByte(offset + 0));
//       obj.x = buf.getIntLE(offset + 1);
//       obj.y = buf.getIntLE(offset + 5);
//       obj.z = buf.getIntLE(offset + 9);
//       obj.offsetForPaintModeX = buf.getIntLE(offset + 13);
//       obj.offsetForPaintModeY = buf.getIntLE(offset + 17);
//       obj.offsetForPaintModeZ = buf.getIntLE(offset + 21);
//       obj.isAltPlaySculptBrushModDown = buf.getByte(offset + 25) != 0;
//       obj.isHoldDownInteraction = buf.getByte(offset + 26) != 0;
//       obj.isDoServerRaytraceForPosition = buf.getByte(offset + 27) != 0;
//       obj.isShowEditNotifications = buf.getByte(offset + 28) != 0;
//       obj.maxLengthToolIgnoreHistory = buf.getIntLE(offset + 29);
//       obj.raycastOriginX = buf.getFloatLE(offset + 33);
//       obj.raycastOriginY = buf.getFloatLE(offset + 37);
//       obj.raycastOriginZ = buf.getFloatLE(offset + 41);
//       obj.raycastDirectionX = buf.getFloatLE(offset + 45);
//       obj.raycastDirectionY = buf.getFloatLE(offset + 49);
//       obj.raycastDirectionZ = buf.getFloatLE(offset + 53);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//       buf.writeIntLE(this.offsetForPaintModeX);
//       buf.writeIntLE(this.offsetForPaintModeY);
//       buf.writeIntLE(this.offsetForPaintModeZ);
//       buf.writeByte(this.isAltPlaySculptBrushModDown ? 1 : 0);
//       buf.writeByte(this.isHoldDownInteraction ? 1 : 0);
//       buf.writeByte(this.isDoServerRaytraceForPosition ? 1 : 0);
//       buf.writeByte(this.isShowEditNotifications ? 1 : 0);
//       buf.writeIntLE(this.maxLengthToolIgnoreHistory);
//       buf.writeFloatLE(this.raycastOriginX);
//       buf.writeFloatLE(this.raycastOriginY);
//       buf.writeFloatLE(this.raycastOriginZ);
//       buf.writeFloatLE(this.raycastDirectionX);
//       buf.writeFloatLE(this.raycastDirectionY);
//       buf.writeFloatLE(this.raycastDirectionZ);
//    }

func (pk *BuilderToolOnUseInteraction) Marshal(io protocol.IO) {
	// TODO: Implement
}
