package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type CameraShakeEffect struct{}

func (*CameraShakeEffect) ID() uint32 {
	return IDCameraShakeEffect
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       CameraShakeEffect obj = new CameraShakeEffect();
//       obj.cameraShakeId = buf.getIntLE(offset + 0);
//       obj.intensity = buf.getFloatLE(offset + 4);
//       obj.mode = AccumulationMode.fromValue(buf.getByte(offset + 8));
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.cameraShakeId);
//       buf.writeFloatLE(this.intensity);
//       buf.writeByte(this.mode.getValue());
//    }

func (pk *CameraShakeEffect) Marshal(io protocol.IO) {
	// TODO: Implement
}
