package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type TrackOrUpdateObjective struct{}

func (*TrackOrUpdateObjective) ID() uint32 {
	return IDTrackOrUpdateObjective
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       TrackOrUpdateObjective obj = new TrackOrUpdateObjective();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          obj.objective = Objective.deserialize(buf, pos);
//          pos += Objective.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.objective != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.objective != null) {
//          this.objective.serialize(buf);
//       }
//    }

func (pk *TrackOrUpdateObjective) Marshal(io protocol.IO) {
	// TODO: Implement
}
