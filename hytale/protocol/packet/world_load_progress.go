package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type WorldLoadProgress struct {
	Status                 string
	PercentComplete        int32
	PercentCompleteSubitem int32
}

func (*WorldLoadProgress) ID() uint32 {
	return IDWorldLoadProgress
}

func (pk *WorldLoadProgress) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasStatus := opts.Has(pk.Status != "")
	opts.Write()

	io.Int32(&pk.PercentComplete)
	io.Int32(&pk.PercentCompleteSubitem)
	if hasStatus {
		io.VarString(&pk.Status, 4096000)
	}
}
