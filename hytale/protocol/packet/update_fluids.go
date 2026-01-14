package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateFluids struct{}

func (*UpdateFluids) ID() uint32 {
	return IDUpdateFluids
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateFluids obj = new UpdateFluids();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int fluidsCount = VarInt.peek(buf, pos);
//          if (fluidsCount < 0) {
//             throw ProtocolException.negativeLength("Fluids", fluidsCount);
//          }
//
//          if (fluidsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Fluids", fluidsCount, 4096000);
//          }
//
//          pos += VarInt.size(fluidsCount);
//          obj.fluids = new HashMap<>(fluidsCount);
//
//          for (int i = 0; i < fluidsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             Fluid val = Fluid.deserialize(buf, pos);
//             pos += Fluid.computeBytesConsumed(buf, pos);
//             if (obj.fluids.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("fluids", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.fluids != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.fluids != null) {
//          if (this.fluids.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Fluids", this.fluids.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.fluids.size());
//
//          for (Entry<Integer, Fluid> e : this.fluids.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateFluids) Marshal(io protocol.IO) {
	// TODO: Implement
}
