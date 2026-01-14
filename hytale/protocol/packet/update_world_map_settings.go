package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateWorldMapSettings struct {
	Enabled                    bool
	AllowTeleportToCoordinates bool
	AllowTeleportToMarkers     bool
	DefaultScale               float32
	MinScale                   float32
	MaxScale                   float32
	BiomeData                  map[int16]protocol.BiomeData
}

func (*UpdateWorldMapSettings) ID() uint32 {
	return IDUpdateWorldMapSettings
}

func (pk *UpdateWorldMapSettings) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasBiomeData := opts.Has(len(pk.BiomeData) > 0)
	opts.Write()

	io.Bool(&pk.Enabled)
	io.Bool(&pk.AllowTeleportToCoordinates)
	io.Bool(&pk.AllowTeleportToMarkers)
	io.Float32(&pk.DefaultScale)
	io.Float32(&pk.MinScale)
	io.Float32(&pk.MaxScale)
	if hasBiomeData {
		protocol.Map(io, &pk.BiomeData, 4096000, io.Int16)
	}
}
