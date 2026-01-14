package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEditorWeatherOverride struct{}

func (*UpdateEditorWeatherOverride) ID() uint32 {
	return IDUpdateEditorWeatherOverride
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEditorWeatherOverride obj = new UpdateEditorWeatherOverride();
//       obj.weatherIndex = buf.getIntLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.weatherIndex);
//    }

func (pk *UpdateEditorWeatherOverride) Marshal(io protocol.IO) {
	// TODO: Implement
}
