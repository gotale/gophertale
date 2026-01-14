package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type Ping struct {
	PingID              int32
	Time                protocol.InstantData
	LastPingValueRaw    int32
	LastPingValueDirect int32
	LastPingValueTick   int32
}

func (*Ping) ID() uint32 {
	return IDPing
}

func (pk *Ping) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	opts.Has(pk.Time != (protocol.InstantData{}))
	opts.Write()

	io.Int32(&pk.PingID)
	protocol.Single(io, &pk.Time)
	io.Int32(&pk.LastPingValueRaw)
	io.Int32(&pk.LastPingValueDirect)
	io.Int32(&pk.LastPingValueTick)
}
