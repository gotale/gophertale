package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type WorldSavingStatus struct {
	WorldSaving bool
}

func (*WorldSavingStatus) ID() uint32 {
	return IDWorldSavingStatus
}

func (pk *WorldSavingStatus) Marshal(io protocol.IO) {
	io.Bool(&pk.WorldSaving)
}
