package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateSunSettings struct{}

func (*UpdateSunSettings) ID() uint32 {
	return IDUpdateSunSettings
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateSunSettings obj = new UpdateSunSettings();
//       obj.heightPercentage = buf.getFloatLE(offset + 0);
//       obj.angleRadians = buf.getFloatLE(offset + 4);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeFloatLE(this.heightPercentage);
//       buf.writeFloatLE(this.angleRadians);
//    }

func (pk *UpdateSunSettings) Marshal(io protocol.IO) {
	// TODO: Implement
}
