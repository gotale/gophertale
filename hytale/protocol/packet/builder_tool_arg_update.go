package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolArgUpdate struct{}

func (*BuilderToolArgUpdate) ID() uint32 {
	return IDBuilderToolArgUpdate
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolArgUpdate obj = new BuilderToolArgUpdate();
//       byte nullBits = buf.getByte(offset);
//       obj.token = buf.getIntLE(offset + 1);
//       obj.section = buf.getIntLE(offset + 5);
//       obj.slot = buf.getIntLE(offset + 9);
//       obj.group = BuilderToolArgGroup.fromValue(buf.getByte(offset + 13));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 22 + buf.getIntLE(offset + 14);
//          int idLen = VarInt.peek(buf, varPos0);
//          if (idLen < 0) {
//             throw ProtocolException.negativeLength("Id", idLen);
//          }
//
//          if (idLen > 4096000) {
//             throw ProtocolException.stringTooLong("Id", idLen, 4096000);
//          }
//
//          obj.id = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 22 + buf.getIntLE(offset + 18);
//          int valueLen = VarInt.peek(buf, varPos1);
//          if (valueLen < 0) {
//             throw ProtocolException.negativeLength("Value", valueLen);
//          }
//
//          if (valueLen > 4096000) {
//             throw ProtocolException.stringTooLong("Value", valueLen, 4096000);
//          }
//
//          obj.value = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.id != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.value != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.token);
//       buf.writeIntLE(this.section);
//       buf.writeIntLE(this.slot);
//       buf.writeByte(this.group.getValue());
//       int idOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int valueOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.id != null) {
//          buf.setIntLE(idOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.id, 4096000);
//       } else {
//          buf.setIntLE(idOffsetSlot, -1);
//       }
//
//       if (this.value != null) {
//          buf.setIntLE(valueOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.value, 4096000);
//       } else {
//          buf.setIntLE(valueOffsetSlot, -1);
//       }
//    }

func (pk *BuilderToolArgUpdate) Marshal(io protocol.IO) {
	// TODO: Implement
}
