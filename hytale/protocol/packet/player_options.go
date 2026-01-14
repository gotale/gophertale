package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type PlayerOptions struct{}

func (*PlayerOptions) ID() uint32 {
	return IDPlayerOptions
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       PlayerOptions obj = new PlayerOptions();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          obj.skin = PlayerSkin.deserialize(buf, pos);
//          pos += PlayerSkin.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.skin != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.skin != null) {
//          this.skin.serialize(buf);
//       }
//    }

func (pk *PlayerOptions) Marshal(io protocol.IO) {
	// TODO: Implement
}
