package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetFluids struct{}

func (*SetFluids) ID() uint32 {
	return IDSetFluids
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetFluids obj = new SetFluids();
//       byte nullBits = buf.getByte(offset);
//       obj.x = buf.getIntLE(offset + 1);
//       obj.y = buf.getIntLE(offset + 5);
//       obj.z = buf.getIntLE(offset + 9);
//       int pos = offset + 13;
//       if ((nullBits & 1) != 0) {
//          int dataCount = VarInt.peek(buf, pos);
//          if (dataCount < 0) {
//             throw ProtocolException.negativeLength("Data", dataCount);
//          }
//
//          if (dataCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Data", dataCount, 4096000);
//          }
//
//          int dataVarLen = VarInt.size(dataCount);
//          if (pos + dataVarLen + dataCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Data", pos + dataVarLen + dataCount * 1, buf.readableBytes());
//          }
//
//          pos += dataVarLen;
//          obj.data = new byte[dataCount];
//
//          for (int i = 0; i < dataCount; i++) {
//             obj.data[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += dataCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.data != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//       if (this.data != null) {
//          if (this.data.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Data", this.data.length, 4096000);
//          }
//
//          VarInt.write(buf, this.data.length);
//
//          for (byte item : this.data) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *SetFluids) Marshal(io protocol.IO) {
	// TODO: Implement
}
