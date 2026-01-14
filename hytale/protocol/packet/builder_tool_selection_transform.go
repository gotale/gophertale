package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolSelectionTransform struct{}

func (*BuilderToolSelectionTransform) ID() uint32 {
	return IDBuilderToolSelectionTransform
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolSelectionTransform obj = new BuilderToolSelectionTransform();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 2) != 0) {
//          obj.initialSelectionMin = BlockPosition.deserialize(buf, offset + 1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          obj.initialSelectionMax = BlockPosition.deserialize(buf, offset + 13);
//       }
//
//       if ((nullBits & 8) != 0) {
//          obj.initialRotationOrigin = Vector3f.deserialize(buf, offset + 25);
//       }
//
//       obj.cutOriginal = buf.getByte(offset + 37) != 0;
//       obj.applyTransformationToSelectionMinMax = buf.getByte(offset + 38) != 0;
//       obj.isExitingTransformMode = buf.getByte(offset + 39) != 0;
//       if ((nullBits & 16) != 0) {
//          obj.initialPastePointForClipboardPaste = BlockPosition.deserialize(buf, offset + 40);
//       }
//
//       int pos = offset + 52;
//       if ((nullBits & 1) != 0) {
//          int transformationMatrixCount = VarInt.peek(buf, pos);
//          if (transformationMatrixCount < 0) {
//             throw ProtocolException.negativeLength("TransformationMatrix", transformationMatrixCount);
//          }
//
//          if (transformationMatrixCount > 4096000) {
//             throw ProtocolException.arrayTooLong("TransformationMatrix", transformationMatrixCount, 4096000);
//          }
//
//          int transformationMatrixVarLen = VarInt.size(transformationMatrixCount);
//          if (pos + transformationMatrixVarLen + transformationMatrixCount * 4L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall(
//                "TransformationMatrix", pos + transformationMatrixVarLen + transformationMatrixCount * 4, buf.readableBytes()
//             );
//          }
//
//          pos += transformationMatrixVarLen;
//          obj.transformationMatrix = new float[transformationMatrixCount];
//
//          for (int i = 0; i < transformationMatrixCount; i++) {
//             obj.transformationMatrix[i] = buf.getFloatLE(pos + i * 4);
//          }
//
//          pos += transformationMatrixCount * 4;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.transformationMatrix != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.initialSelectionMin != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.initialSelectionMax != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       if (this.initialRotationOrigin != null) {
//          nullBits = (byte)(nullBits | 8);
//       }
//
//       if (this.initialPastePointForClipboardPaste != null) {
//          nullBits = (byte)(nullBits | 16);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.initialSelectionMin != null) {
//          this.initialSelectionMin.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       if (this.initialSelectionMax != null) {
//          this.initialSelectionMax.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       if (this.initialRotationOrigin != null) {
//          this.initialRotationOrigin.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       buf.writeByte(this.cutOriginal ? 1 : 0);
//       buf.writeByte(this.applyTransformationToSelectionMinMax ? 1 : 0);
//       buf.writeByte(this.isExitingTransformMode ? 1 : 0);
//       if (this.initialPastePointForClipboardPaste != null) {
//          this.initialPastePointForClipboardPaste.serialize(buf);
//       } else {
//          buf.writeZero(12);
//       }
//
//       if (this.transformationMatrix != null) {
//          if (this.transformationMatrix.length > 4096000) {
//             throw ProtocolException.arrayTooLong("TransformationMatrix", this.transformationMatrix.length, 4096000);
//          }
//
//          VarInt.write(buf, this.transformationMatrix.length);
//
//          for (float item : this.transformationMatrix) {
//             buf.writeFloatLE(item);
//          }
//       }
//    }

func (pk *BuilderToolSelectionTransform) Marshal(io protocol.IO) {
	// TODO: Implement
}
