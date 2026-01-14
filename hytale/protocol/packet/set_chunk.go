package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetChunk struct{}

func (*SetChunk) ID() uint32 {
	return IDSetChunk
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetChunk obj = new SetChunk();
//       byte nullBits = buf.getByte(offset);
//       obj.x = buf.getIntLE(offset + 1);
//       obj.y = buf.getIntLE(offset + 5);
//       obj.z = buf.getIntLE(offset + 9);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 25 + buf.getIntLE(offset + 13);
//          int localLightCount = VarInt.peek(buf, varPos0);
//          if (localLightCount < 0) {
//             throw ProtocolException.negativeLength("LocalLight", localLightCount);
//          }
//
//          if (localLightCount > 4096000) {
//             throw ProtocolException.arrayTooLong("LocalLight", localLightCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          if (varPos0 + varIntLen + localLightCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("LocalLight", varPos0 + varIntLen + localLightCount * 1, buf.readableBytes());
//          }
//
//          obj.localLight = new byte[localLightCount];
//
//          for (int i = 0; i < localLightCount; i++) {
//             obj.localLight[i] = buf.getByte(varPos0 + varIntLen + i * 1);
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 25 + buf.getIntLE(offset + 17);
//          int globalLightCount = VarInt.peek(buf, varPos1);
//          if (globalLightCount < 0) {
//             throw ProtocolException.negativeLength("GlobalLight", globalLightCount);
//          }
//
//          if (globalLightCount > 4096000) {
//             throw ProtocolException.arrayTooLong("GlobalLight", globalLightCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + globalLightCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("GlobalLight", varPos1 + varIntLen + globalLightCount * 1, buf.readableBytes());
//          }
//
//          obj.globalLight = new byte[globalLightCount];
//
//          for (int i = 0; i < globalLightCount; i++) {
//             obj.globalLight[i] = buf.getByte(varPos1 + varIntLen + i * 1);
//          }
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 25 + buf.getIntLE(offset + 21);
//          int dataCount = VarInt.peek(buf, varPos2);
//          if (dataCount < 0) {
//             throw ProtocolException.negativeLength("Data", dataCount);
//          }
//
//          if (dataCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Data", dataCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos2);
//          if (varPos2 + varIntLen + dataCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Data", varPos2 + varIntLen + dataCount * 1, buf.readableBytes());
//          }
//
//          obj.data = new byte[dataCount];
//
//          for (int i = 0; i < dataCount; i++) {
//             obj.data[i] = buf.getByte(varPos2 + varIntLen + i * 1);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.localLight != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.globalLight != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.data != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//       int localLightOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int globalLightOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int dataOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.localLight != null) {
//          buf.setIntLE(localLightOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.localLight.length > 4096000) {
//             throw ProtocolException.arrayTooLong("LocalLight", this.localLight.length, 4096000);
//          }
//
//          VarInt.write(buf, this.localLight.length);
//
//          for (byte item : this.localLight) {
//             buf.writeByte(item);
//          }
//       } else {
//          buf.setIntLE(localLightOffsetSlot, -1);
//       }
//
//       if (this.globalLight != null) {
//          buf.setIntLE(globalLightOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.globalLight.length > 4096000) {
//             throw ProtocolException.arrayTooLong("GlobalLight", this.globalLight.length, 4096000);
//          }
//
//          VarInt.write(buf, this.globalLight.length);
//
//          for (byte item : this.globalLight) {
//             buf.writeByte(item);
//          }
//       } else {
//          buf.setIntLE(globalLightOffsetSlot, -1);
//       }
//
//       if (this.data != null) {
//          buf.setIntLE(dataOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.data.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Data", this.data.length, 4096000);
//          }
//
//          VarInt.write(buf, this.data.length);
//
//          for (byte item : this.data) {
//             buf.writeByte(item);
//          }
//       } else {
//          buf.setIntLE(dataOffsetSlot, -1);
//       }
//    }

func (pk *SetChunk) Marshal(io protocol.IO) {
	// TODO: Implement
}
