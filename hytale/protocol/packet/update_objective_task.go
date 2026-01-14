package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateObjectiveTask struct{}

func (*UpdateObjectiveTask) ID() uint32 {
	return IDUpdateObjectiveTask
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateObjectiveTask obj = new UpdateObjectiveTask();
//       byte nullBits = buf.getByte(offset);
//       obj.objectiveUuid = PacketIO.readUUID(buf, offset + 1);
//       obj.taskIndex = buf.getIntLE(offset + 17);
//       int pos = offset + 21;
//       if ((nullBits & 1) != 0) {
//          obj.task = ObjectiveTask.deserialize(buf, pos);
//          pos += ObjectiveTask.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.task != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       PacketIO.writeUUID(buf, this.objectiveUuid);
//       buf.writeIntLE(this.taskIndex);
//       if (this.task != null) {
//          this.task.serialize(buf);
//       }
//    }

func (pk *UpdateObjectiveTask) Marshal(io protocol.IO) {
	// TODO: Implement
}
