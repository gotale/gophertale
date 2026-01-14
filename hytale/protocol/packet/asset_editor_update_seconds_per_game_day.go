package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorUpdateSecondsPerGameDay struct{}

func (*AssetEditorUpdateSecondsPerGameDay) ID() uint32 {
	return IDAssetEditorUpdateSecondsPerGameDay
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorUpdateSecondsPerGameDay obj = new AssetEditorUpdateSecondsPerGameDay();
//       obj.daytimeDurationSeconds = buf.getIntLE(offset + 0);
//       obj.nighttimeDurationSeconds = buf.getIntLE(offset + 4);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.daytimeDurationSeconds);
//       buf.writeIntLE(this.nighttimeDurationSeconds);
//    }

func (pk *AssetEditorUpdateSecondsPerGameDay) Marshal(io protocol.IO) {
	// TODO: Implement
}
