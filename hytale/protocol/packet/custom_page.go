package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type CustomPage struct{}

func (*CustomPage) ID() uint32 {
	return IDCustomPage
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       CustomPage obj = new CustomPage();
//       byte nullBits = buf.getByte(offset);
//       obj.isInitial = buf.getByte(offset + 1) != 0;
//       obj.clear = buf.getByte(offset + 2) != 0;
//       obj.lifetime = CustomPageLifetime.fromValue(buf.getByte(offset + 3));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 16 + buf.getIntLE(offset + 4);
//          int keyLen = VarInt.peek(buf, varPos0);
//          if (keyLen < 0) {
//             throw ProtocolException.negativeLength("Key", keyLen);
//          }
//
//          if (keyLen > 4096000) {
//             throw ProtocolException.stringTooLong("Key", keyLen, 4096000);
//          }
//
//          obj.key = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 16 + buf.getIntLE(offset + 8);
//          int commandsCount = VarInt.peek(buf, varPos1);
//          if (commandsCount < 0) {
//             throw ProtocolException.negativeLength("Commands", commandsCount);
//          }
//
//          if (commandsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Commands", commandsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + commandsCount * 2L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Commands", varPos1 + varIntLen + commandsCount * 2, buf.readableBytes());
//          }
//
//          obj.commands = new CustomUICommand[commandsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < commandsCount; i++) {
//             obj.commands[i] = CustomUICommand.deserialize(buf, elemPos);
//             elemPos += CustomUICommand.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 16 + buf.getIntLE(offset + 12);
//          int eventBindingsCount = VarInt.peek(buf, varPos2);
//          if (eventBindingsCount < 0) {
//             throw ProtocolException.negativeLength("EventBindings", eventBindingsCount);
//          }
//
//          if (eventBindingsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("EventBindings", eventBindingsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos2);
//          if (varPos2 + varIntLen + eventBindingsCount * 3L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("EventBindings", varPos2 + varIntLen + eventBindingsCount * 3, buf.readableBytes());
//          }
//
//          obj.eventBindings = new CustomUIEventBinding[eventBindingsCount];
//          int elemPos = varPos2 + varIntLen;
//
//          for (int i = 0; i < eventBindingsCount; i++) {
//             obj.eventBindings[i] = CustomUIEventBinding.deserialize(buf, elemPos);
//             elemPos += CustomUIEventBinding.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.key != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.commands != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.eventBindings != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.isInitial ? 1 : 0);
//       buf.writeByte(this.clear ? 1 : 0);
//       buf.writeByte(this.lifetime.getValue());
//       int keyOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int commandsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int eventBindingsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.key != null) {
//          buf.setIntLE(keyOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.key, 4096000);
//       } else {
//          buf.setIntLE(keyOffsetSlot, -1);
//       }
//
//       if (this.commands != null) {
//          buf.setIntLE(commandsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.commands.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Commands", this.commands.length, 4096000);
//          }
//
//          VarInt.write(buf, this.commands.length);
//
//          for (CustomUICommand item : this.commands) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(commandsOffsetSlot, -1);
//       }
//
//       if (this.eventBindings != null) {
//          buf.setIntLE(eventBindingsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.eventBindings.length > 4096000) {
//             throw ProtocolException.arrayTooLong("EventBindings", this.eventBindings.length, 4096000);
//          }
//
//          VarInt.write(buf, this.eventBindings.length);
//
//          for (CustomUIEventBinding item : this.eventBindings) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(eventBindingsOffsetSlot, -1);
//       }
//    }

func (pk *CustomPage) Marshal(io protocol.IO) {
	// TODO: Implement
}
