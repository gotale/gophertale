package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

const (
	DisconnectTypeDisconnect = iota
	DisconnectTypeCrash
)

type Disconnect struct {
	Reason string
	Type   uint8
}

func (*Disconnect) ID() uint32 {
	return IDDisconnect
}

func (pk *Disconnect) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasReason := opts.Has(pk.Reason != "")
	opts.Write()

	io.Uint8(&pk.Type)
	if hasReason {
		io.VarString(&pk.Reason, 4096000)
	}
}
