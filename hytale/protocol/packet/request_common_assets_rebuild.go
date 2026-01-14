package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type RequestCommonAssetsRebuild struct{}

func (*RequestCommonAssetsRebuild) ID() uint32 {
	return IDRequestCommonAssetsRebuild
}

func (pk *RequestCommonAssetsRebuild) Marshal(io protocol.IO) {}
