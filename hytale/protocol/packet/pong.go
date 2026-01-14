package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

const (
	PongTypeRaw = iota
	PongTypeDirect
	PongTypeTick
)

type Pong struct {
	PingID          int32
	Time            protocol.InstantData
	Type            uint8
	PacketQueueSize int16
}

func (*Pong) ID() uint32 {
	return IDPong
}

func (pk *Pong) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	opts.Has(pk.Time != (protocol.InstantData{}))
	opts.Write()

	io.Int32(&pk.PingID)
	protocol.Single(io, &pk.Time)
	io.Uint8(&pk.Type)
	io.Int16(&pk.PacketQueueSize)
}
