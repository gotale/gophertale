package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateViewBobbing struct {
	UpdateType uint8
	Profiles   map[uint8]ViewBobbing
}

func (*UpdateViewBobbing) ID() uint32 {
	return IDUpdateViewBobbing
}

func (pk *UpdateViewBobbing) Marshal(io protocol.IO) {
	// TODO: Implement
}

type ViewBobbing struct {
	FirstPerson CameraShakeConfig
}

func (x *ViewBobbing) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasFirstPerson := opts.Has(x.FirstPerson != CameraShakeConfig{})
	opts.Write()

	if hasFirstPerson {
		protocol.Single(io, &x.FirstPerson)
	}
}

type CameraShakeConfig struct{}

func (x *CameraShakeConfig) Marshal(io protocol.IO) {
	// TODO: Implementation
}
