package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type EntityUpdates struct{}

func (*EntityUpdates) ID() uint32 {
	return IDEntityUpdates
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       EntityUpdates obj = new EntityUpdates();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          int removedCount = VarInt.peek(buf, varPos0);
//          if (removedCount < 0) {
//             throw ProtocolException.negativeLength("Removed", removedCount);
//          }
//
//          if (removedCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Removed", removedCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          if (varPos0 + varIntLen + removedCount * 4L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Removed", varPos0 + varIntLen + removedCount * 4, buf.readableBytes());
//          }
//
//          obj.removed = new int[removedCount];
//
//          for (int i = 0; i < removedCount; i++) {
//             obj.removed[i] = buf.getIntLE(varPos0 + varIntLen + i * 4);
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int updatesCount = VarInt.peek(buf, varPos1);
//          if (updatesCount < 0) {
//             throw ProtocolException.negativeLength("Updates", updatesCount);
//          }
//
//          if (updatesCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Updates", updatesCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + updatesCount * 5L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Updates", varPos1 + varIntLen + updatesCount * 5, buf.readableBytes());
//          }
//
//          obj.updates = new EntityUpdate[updatesCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < updatesCount; i++) {
//             obj.updates[i] = EntityUpdate.deserialize(buf, elemPos);
//             elemPos += EntityUpdate.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.removed != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.updates != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int removedOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int updatesOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.removed != null) {
//          buf.setIntLE(removedOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.removed.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Removed", this.removed.length, 4096000);
//          }
//
//          VarInt.write(buf, this.removed.length);
//
//          for (int item : this.removed) {
//             buf.writeIntLE(item);
//          }
//       } else {
//          buf.setIntLE(removedOffsetSlot, -1);
//       }
//
//       if (this.updates != null) {
//          buf.setIntLE(updatesOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.updates.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Updates", this.updates.length, 4096000);
//          }
//
//          VarInt.write(buf, this.updates.length);
//
//          for (EntityUpdate item : this.updates) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(updatesOffsetSlot, -1);
//       }
//    }

func (pk *EntityUpdates) Marshal(io protocol.IO) {
	// TODO: Implement
}
