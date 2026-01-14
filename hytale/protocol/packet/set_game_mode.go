package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

const (
	GameModeAdventure = iota
	GameModeCreative
)

type SetGameMode struct {
	GameMode uint8
}

func (*SetGameMode) ID() uint32 {
	return IDSetGameMode
}

func (pk *SetGameMode) Marshal(io protocol.IO) {
	io.Uint8(&pk.GameMode)
}
