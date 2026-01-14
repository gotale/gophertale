package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetServerCamera struct{}

func (*SetServerCamera) ID() uint32 {
	return IDSetServerCamera
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetServerCamera obj = new SetServerCamera();
//       byte nullBits = buf.getByte(offset);
//       obj.clientCameraView = ClientCameraView.fromValue(buf.getByte(offset + 1));
//       obj.isLocked = buf.getByte(offset + 2) != 0;
//       if ((nullBits & 1) != 0) {
//          obj.cameraSettings = ServerCameraSettings.deserialize(buf, offset + 3);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.cameraSettings != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.clientCameraView.getValue());
//       buf.writeByte(this.isLocked ? 1 : 0);
//       if (this.cameraSettings != null) {
//          this.cameraSettings.serialize(buf);
//       } else {
//          buf.writeZero(154);
//       }
//    }

func (pk *SetServerCamera) Marshal(io protocol.IO) {
	// TODO: Implement
}
