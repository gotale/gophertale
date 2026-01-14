package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SetChunkHeightmap struct{}

func (*SetChunkHeightmap) ID() uint32 {
	return IDSetChunkHeightmap
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SetChunkHeightmap obj = new SetChunkHeightmap();
//       byte nullBits = buf.getByte(offset);
//       obj.x = buf.getIntLE(offset + 1);
//       obj.z = buf.getIntLE(offset + 5);
//       int pos = offset + 9;
//       if ((nullBits & 1) != 0) {
//          int heightmapCount = VarInt.peek(buf, pos);
//          if (heightmapCount < 0) {
//             throw ProtocolException.negativeLength("Heightmap", heightmapCount);
//          }
//
//          if (heightmapCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Heightmap", heightmapCount, 4096000);
//          }
//
//          int heightmapVarLen = VarInt.size(heightmapCount);
//          if (pos + heightmapVarLen + heightmapCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Heightmap", pos + heightmapVarLen + heightmapCount * 1, buf.readableBytes());
//          }
//
//          pos += heightmapVarLen;
//          obj.heightmap = new byte[heightmapCount];
//
//          for (int i = 0; i < heightmapCount; i++) {
//             obj.heightmap[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += heightmapCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.heightmap != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.z);
//       if (this.heightmap != null) {
//          if (this.heightmap.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Heightmap", this.heightmap.length, 4096000);
//          }
//
//          VarInt.write(buf, this.heightmap.length);
//
//          for (byte item : this.heightmap) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *SetChunkHeightmap) Marshal(io protocol.IO) {
	// TODO: Implement
}
