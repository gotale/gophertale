package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateTranslations struct {
	UpdateType   uint8
	Translations map[string]string
}

func (*UpdateTranslations) ID() uint32 {
	return IDUpdateTranslations
}

func (pk *UpdateTranslations) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasTranslations := opts.Has(len(pk.Translations) > 0)
	opts.Write()

	io.Uint8(&pk.UpdateType)
	if hasTranslations {
		protocol.FuncMap(io, &pk.Translations, 4096000, func(s *string) {
			io.VarString(s, 4096000)
		}, func(s *string) {
			io.VarString(s, 4096000)
		})
	}
}
