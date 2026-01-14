package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetFinalize struct{}

func (*AssetFinalize) ID() uint32 {
	return IDAssetFinalize
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new AssetFinalize();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *AssetFinalize) Marshal(io protocol.IO) {
	// TODO: Implement
}
