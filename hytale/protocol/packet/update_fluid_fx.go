package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateFluidFX struct{}

func (*UpdateFluidFX) ID() uint32 {
	return IDUpdateFluidFX
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateFluidFX obj = new UpdateFluidFX();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int fluidFXCount = VarInt.peek(buf, pos);
//          if (fluidFXCount < 0) {
//             throw ProtocolException.negativeLength("FluidFX", fluidFXCount);
//          }
//
//          if (fluidFXCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("FluidFX", fluidFXCount, 4096000);
//          }
//
//          pos += VarInt.size(fluidFXCount);
//          obj.fluidFX = new HashMap<>(fluidFXCount);
//
//          for (int i = 0; i < fluidFXCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             FluidFX val = FluidFX.deserialize(buf, pos);
//             pos += FluidFX.computeBytesConsumed(buf, pos);
//             if (obj.fluidFX.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("fluidFX", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.fluidFX != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.fluidFX != null) {
//          if (this.fluidFX.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("FluidFX", this.fluidFX.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.fluidFX.size());
//
//          for (Entry<Integer, FluidFX> e : this.fluidFX.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateFluidFX) Marshal(io protocol.IO) {
	// TODO: Implement
}
