package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UntrackObjective struct{}

func (*UntrackObjective) ID() uint32 {
	return IDUntrackObjective
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UntrackObjective obj = new UntrackObjective();
//       obj.objectiveUuid = PacketIO.readUUID(buf, offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       PacketIO.writeUUID(buf, this.objectiveUuid);
//    }

func (pk *UntrackObjective) Marshal(io protocol.IO) {
	// TODO: Implement
}
