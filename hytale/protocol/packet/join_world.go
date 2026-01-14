package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type JoinWorld struct{}

func (*JoinWorld) ID() uint32 {
	return IDJoinWorld
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       JoinWorld obj = new JoinWorld();
//       obj.clearWorld = buf.getByte(offset + 0) != 0;
//       obj.fadeInOut = buf.getByte(offset + 1) != 0;
//       obj.worldUuid = PacketIO.readUUID(buf, offset + 2);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.clearWorld ? 1 : 0);
//       buf.writeByte(this.fadeInOut ? 1 : 0);
//       PacketIO.writeUUID(buf, this.worldUuid);
//    }

func (pk *JoinWorld) Marshal(io protocol.IO) {
	// TODO: Implement
}
