package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateWorldMapVisible struct {
	Visible bool
}

func (*UpdateWorldMapVisible) ID() uint32 {
	return IDUpdateWorldMapVisible
}

func (pk *UpdateWorldMapVisible) Marshal(io protocol.IO) {
	io.Bool(&pk.Visible)
}
