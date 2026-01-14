package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateModelvfxs struct{}

func (*UpdateModelvfxs) ID() uint32 {
	return IDUpdateModelvfxs
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateModelvfxs obj = new UpdateModelvfxs();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int modelVFXsCount = VarInt.peek(buf, pos);
//          if (modelVFXsCount < 0) {
//             throw ProtocolException.negativeLength("ModelVFXs", modelVFXsCount);
//          }
//
//          if (modelVFXsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ModelVFXs", modelVFXsCount, 4096000);
//          }
//
//          pos += VarInt.size(modelVFXsCount);
//          obj.modelVFXs = new HashMap<>(modelVFXsCount);
//
//          for (int i = 0; i < modelVFXsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             ModelVFX val = ModelVFX.deserialize(buf, pos);
//             pos += ModelVFX.computeBytesConsumed(buf, pos);
//             if (obj.modelVFXs.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("modelVFXs", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.modelVFXs != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.modelVFXs != null) {
//          if (this.modelVFXs.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("ModelVFXs", this.modelVFXs.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.modelVFXs.size());
//
//          for (Entry<Integer, ModelVFX> e : this.modelVFXs.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateModelvfxs) Marshal(io protocol.IO) {
	// TODO: Implement
}
