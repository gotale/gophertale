package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetChunkTintmap struct{}

func (*SetChunkTintmap) ID() uint32 {
	return IDSetChunkTintmap
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetChunkTintmap obj = new SetChunkTintmap();
//       byte nullBits = buf.getByte(offset);
//       obj.x = buf.getIntLE(offset + 1);
//       obj.z = buf.getIntLE(offset + 5);
//       int pos = offset + 9;
//       if ((nullBits & 1) != 0) {
//          int tintmapCount = VarInt.peek(buf, pos);
//          if (tintmapCount < 0) {
//             throw ProtocolException.negativeLength("Tintmap", tintmapCount);
//          }
//
//          if (tintmapCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Tintmap", tintmapCount, 4096000);
//          }
//
//          int tintmapVarLen = VarInt.size(tintmapCount);
//          if (pos + tintmapVarLen + tintmapCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Tintmap", pos + tintmapVarLen + tintmapCount * 1, buf.readableBytes());
//          }
//
//          pos += tintmapVarLen;
//          obj.tintmap = new byte[tintmapCount];
//
//          for (int i = 0; i < tintmapCount; i++) {
//             obj.tintmap[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += tintmapCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.tintmap != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.z);
//       if (this.tintmap != null) {
//          if (this.tintmap.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Tintmap", this.tintmap.length, 4096000);
//          }
//
//          VarInt.write(buf, this.tintmap.length);
//
//          for (byte item : this.tintmap) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *SetChunkTintmap) Marshal(io protocol.IO) {
	// TODO: Implement
}
