package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetClientId struct {
	ClientID int32
}

func (*SetClientId) ID() uint32 {
	return IDSetClientId
}

func (pk *SetClientId) Marshal(io protocol.IO) {
	io.Int32(&pk.ClientID)
}
