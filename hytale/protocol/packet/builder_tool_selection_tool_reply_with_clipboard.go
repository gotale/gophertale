package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSelectionToolReplyWithClipboard struct{}

func (*BuilderToolSelectionToolReplyWithClipboard) ID() uint32 {
	return IDBuilderToolSelectionToolReplyWithClipboard
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSelectionToolReplyWithClipboard obj = new BuilderToolSelectionToolReplyWithClipboard();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 9 + buf.getIntLE(offset + 1);
//          int blocksChangeCount = VarInt.peek(buf, varPos0);
//          if (blocksChangeCount < 0) {
//             throw ProtocolException.negativeLength("BlocksChange", blocksChangeCount);
//          }
//
//          if (blocksChangeCount > 4096000) {
//             throw ProtocolException.arrayTooLong("BlocksChange", blocksChangeCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          if (varPos0 + varIntLen + blocksChangeCount * 17L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("BlocksChange", varPos0 + varIntLen + blocksChangeCount * 17, buf.readableBytes());
//          }
//
//          obj.blocksChange = new BlockChange[blocksChangeCount];
//          int elemPos = varPos0 + varIntLen;
//
//          for (int i = 0; i < blocksChangeCount; i++) {
//             obj.blocksChange[i] = BlockChange.deserialize(buf, elemPos);
//             elemPos += BlockChange.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 9 + buf.getIntLE(offset + 5);
//          int fluidsChangeCount = VarInt.peek(buf, varPos1);
//          if (fluidsChangeCount < 0) {
//             throw ProtocolException.negativeLength("FluidsChange", fluidsChangeCount);
//          }
//
//          if (fluidsChangeCount > 4096000) {
//             throw ProtocolException.arrayTooLong("FluidsChange", fluidsChangeCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + fluidsChangeCount * 17L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("FluidsChange", varPos1 + varIntLen + fluidsChangeCount * 17, buf.readableBytes());
//          }
//
//          obj.fluidsChange = new FluidChange[fluidsChangeCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < fluidsChangeCount; i++) {
//             obj.fluidsChange[i] = FluidChange.deserialize(buf, elemPos);
//             elemPos += FluidChange.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.blocksChange != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.fluidsChange != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       int blocksChangeOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int fluidsChangeOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.blocksChange != null) {
//          buf.setIntLE(blocksChangeOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.blocksChange.length > 4096000) {
//             throw ProtocolException.arrayTooLong("BlocksChange", this.blocksChange.length, 4096000);
//          }
//
//          VarInt.write(buf, this.blocksChange.length);
//
//          for (BlockChange item : this.blocksChange) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(blocksChangeOffsetSlot, -1);
//       }
//
//       if (this.fluidsChange != null) {
//          buf.setIntLE(fluidsChangeOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.fluidsChange.length > 4096000) {
//             throw ProtocolException.arrayTooLong("FluidsChange", this.fluidsChange.length, 4096000);
//          }
//
//          VarInt.write(buf, this.fluidsChange.length);
//
//          for (FluidChange item : this.fluidsChange) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(fluidsChangeOffsetSlot, -1);
//       }
//    }

func (pk *BuilderToolSelectionToolReplyWithClipboard) Marshal(io protocol.IO) {
	// TODO: Implement
}
