package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetPart struct{}

func (*AssetPart) ID() uint32 {
	return IDAssetPart
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetPart obj = new AssetPart();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int partCount = VarInt.peek(buf, pos);
//          if (partCount < 0) {
//             throw ProtocolException.negativeLength("Part", partCount);
//          }
//
//          if (partCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Part", partCount, 4096000);
//          }
//
//          int partVarLen = VarInt.size(partCount);
//          if (pos + partVarLen + partCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Part", pos + partVarLen + partCount * 1, buf.readableBytes());
//          }
//
//          pos += partVarLen;
//          obj.part = new byte[partCount];
//
//          for (int i = 0; i < partCount; i++) {
//             obj.part[i] = buf.getByte(pos + i * 1);
//          }
//
//          pos += partCount * 1;
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.part != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.part != null) {
//          if (this.part.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Part", this.part.length, 4096000);
//          }
//
//          VarInt.write(buf, this.part.length);
//
//          for (byte item : this.part) {
//             buf.writeByte(item);
//          }
//       }
//    }

func (pk *AssetPart) Marshal(io protocol.IO) {
	// TODO: Implement
}
