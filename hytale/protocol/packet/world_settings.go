package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type WorldSettings struct {
	WorldHeight int32
	Assets      []protocol.Asset
}

func (*WorldSettings) ID() uint32 {
	return IDWorldSettings
}

func (pk *WorldSettings) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasAssets := opts.Has(len(pk.Assets) > 0)
	opts.Write()

	io.Int32(&pk.WorldHeight)
	if hasAssets {
		protocol.Slice(io, &pk.Assets, 4096000)
	}
}
