package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdatePostFxSettings struct{}

func (*UpdatePostFxSettings) ID() uint32 {
	return IDUpdatePostFxSettings
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdatePostFxSettings obj = new UpdatePostFxSettings();
//       obj.globalIntensity = buf.getFloatLE(offset + 0);
//       obj.power = buf.getFloatLE(offset + 4);
//       obj.sunshaftScale = buf.getFloatLE(offset + 8);
//       obj.sunIntensity = buf.getFloatLE(offset + 12);
//       obj.sunshaftIntensity = buf.getFloatLE(offset + 16);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeFloatLE(this.globalIntensity);
//       buf.writeFloatLE(this.power);
//       buf.writeFloatLE(this.sunshaftScale);
//       buf.writeFloatLE(this.sunIntensity);
//       buf.writeFloatLE(this.sunshaftIntensity);
//    }

func (pk *UpdatePostFxSettings) Marshal(io protocol.IO) {
	// TODO: Implement
}
