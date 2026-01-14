package packet

import (
	"github.com/google/uuid"
	"github.com/gotale/gophertale/hytale/protocol"
)

const (
	ClientTypeGame = iota
	ClientTypeEditor
)

type Connect struct {
	ProtocolHash   string
	ClientType     uint8
	Language       string
	IdentityToken  string
	UUID           uuid.UUID
	Username       string
	ReferralData   []byte
	ReferralSource protocol.HostAddress
}

func (*Connect) ID() uint32 {
	return IDConnect
}

func (pk *Connect) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasLanguage := opts.Has(pk.Language != "")
	hasIdentityToken := opts.Has(pk.IdentityToken != "")
	hasReferralData := opts.Has(pk.ReferralData != nil)
	hasReferralSource := opts.Has(pk.ReferralSource != (protocol.HostAddress{}))
	opts.Write()

	io.FixedAsciiString(&pk.ProtocolHash, 64)
	io.Uint8(&pk.ClientType)
	io.UUID(&pk.UUID)
	if hasLanguage {
		io.VarStringAscii(&pk.Language, 128)
	}
	if hasIdentityToken {
		io.VarString(&pk.IdentityToken, 8192)
	}
	io.VarStringAscii(&pk.Username, 16)
	if hasReferralData {
		io.ByteSlice(&pk.ReferralData, 4096)
	}
	if hasReferralSource {
		protocol.Single(io, &pk.ReferralSource)
	}
}
