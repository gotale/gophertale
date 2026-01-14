package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type WorldLoadFinished struct{}

func (*WorldLoadFinished) ID() uint32 {
	return IDWorldLoadFinished
}

func (pk *WorldLoadFinished) Marshal(io protocol.IO) {}
