package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type DisplayDebug struct{}

func (*DisplayDebug) ID() uint32 {
	return IDDisplayDebug
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       DisplayDebug obj = new DisplayDebug();
//       byte nullBits = buf.getByte(offset);
//       obj.shape = DebugShape.fromValue(buf.getByte(offset + 1));
//       if ((nullBits & 2) != 0) {
//          obj.color = Vector3f.deserialize(buf, offset + 2);
//       }
//
//       obj.time = buf.getFloatLE(offset + 14);
//       obj.fade = buf.getByte(offset + 18) != 0;
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 27 + buf.getIntLE(offset + 19);
//          int matrixCount = VarInt.peek(buf, varPos0);
//          if (matrixCount < 0) {
//             throw ProtocolException.negativeLength("Matrix", matrixCount);
//          }
//
//          if (matrixCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Matrix", matrixCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          if (varPos0 + varIntLen + matrixCount * 4L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Matrix", varPos0 + varIntLen + matrixCount * 4, buf.readableBytes());
//          }
//
//          obj.matrix = new float[matrixCount];
//
//          for (int i = 0; i < matrixCount; i++) {
//             obj.matrix[i] = buf.getFloatLE(varPos0 + varIntLen + i * 4);
//          }
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos1 = offset + 27 + buf.getIntLE(offset + 23);
//          int frustumProjectionCount = VarInt.peek(buf, varPos1);
//          if (frustumProjectionCount < 0) {
//             throw ProtocolException.negativeLength("FrustumProjection", frustumProjectionCount);
//          }
//
//          if (frustumProjectionCount > 4096000) {
//             throw ProtocolException.arrayTooLong("FrustumProjection", frustumProjectionCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + frustumProjectionCount * 4L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("FrustumProjection", varPos1 + varIntLen + frustumProjectionCount * 4, buf.readableBytes());
//          }
//
//          obj.frustumProjection = new float[frustumProjectionCount];
//
//          for (int i = 0; i < frustumProjectionCount; i++) {
//             obj.frustumProjection[i] = buf.getFloatLE(varPos1 + varIntLen + i * 4);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.matrix != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.color != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.frustumProjection != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.shape.getValue());
//       if (this.color != null) {
//          this.color.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       buf.writeFloatLE(this.time);
//       buf.writeByte(this.fade ? 1 : 0);
//       int matrixOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int frustumProjectionOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.matrix != null) {
//          buf.setIntLE(matrixOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.matrix.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Matrix", this.matrix.length, 4096000);
//          }
//
//          VarInt.write(buf, this.matrix.length);
//
//          for (float item : this.matrix) {
//             buf.writeFloatLE(item);
//          }
//       } else {
//          buf.setIntLE(matrixOffsetSlot, -1);
//       }
//
//       if (this.frustumProjection != null) {
//          buf.setIntLE(frustumProjectionOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.frustumProjection.length > 4096000) {
//             throw ProtocolException.arrayTooLong("FrustumProjection", this.frustumProjection.length, 4096000);
//          }
//
//          VarInt.write(buf, this.frustumProjection.length);
//
//          for (float item : this.frustumProjection) {
//             buf.writeFloatLE(item);
//          }
//       } else {
//          buf.setIntLE(frustumProjectionOffsetSlot, -1);
//       }
//    }

func (pk *DisplayDebug) Marshal(io protocol.IO) {
	// TODO: Implement
}
