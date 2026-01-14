package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PrefabUnselectPrefab struct{}

func (*PrefabUnselectPrefab) ID() uint32 {
	return IDPrefabUnselectPrefab
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       return new PrefabUnselectPrefab();
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//    }

func (pk *PrefabUnselectPrefab) Marshal(io protocol.IO) {
	// TODO: Implement
}
