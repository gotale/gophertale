package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateSleepState struct{}

func (*UpdateSleepState) ID() uint32 {
	return IDUpdateSleepState
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateSleepState obj = new UpdateSleepState();
//       byte nullBits = buf.getByte(offset);
//       obj.grayFade = buf.getByte(offset + 1) != 0;
//       obj.sleepUi = buf.getByte(offset + 2) != 0;
//       if ((nullBits & 1) != 0) {
//          obj.clock = SleepClock.deserialize(buf, offset + 3);
//       }
//
//       int pos = offset + 36;
//       if ((nullBits & 2) != 0) {
//          obj.multiplayer = SleepMultiplayer.deserialize(buf, pos);
//          pos += SleepMultiplayer.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.clock != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.multiplayer != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.grayFade ? 1 : 0);
//       buf.writeByte(this.sleepUi ? 1 : 0);
//       if (this.clock != null) {
//          this.clock.serialize(buf);
//       } else {
//          buf.writeZero(33);
//       }
//
//       if (this.multiplayer != null) {
//          this.multiplayer.serialize(buf);
//       }
//    }

func (pk *UpdateSleepState) Marshal(io protocol.IO) {
	// TODO: Implement
}
