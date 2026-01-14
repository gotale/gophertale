package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateFeatures struct{}

func (*UpdateFeatures) ID() uint32 {
	return IDUpdateFeatures
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateFeatures obj = new UpdateFeatures();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int featuresCount = VarInt.peek(buf, pos);
//          if (featuresCount < 0) {
//             throw ProtocolException.negativeLength("Features", featuresCount);
//          }
//
//          if (featuresCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Features", featuresCount, 4096000);
//          }
//
//          pos += VarInt.size(featuresCount);
//          obj.features = new HashMap<>(featuresCount);
//
//          for (int i = 0; i < featuresCount; i++) {
//             ClientFeature key = ClientFeature.fromValue(buf.getByte(pos));
//             pos++;
//             boolean val = buf.getByte(pos) != 0;
//             pos++;
//             if (obj.features.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("features", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.features != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.features != null) {
//          if (this.features.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Features", this.features.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.features.size());
//
//          for (Entry<ClientFeature, Boolean> e : this.features.entrySet()) {
//             buf.writeByte(e.getKey().getValue());
//             buf.writeByte(e.getValue() ? 1 : 0);
//          }
//       }
//    }

func (pk *UpdateFeatures) Marshal(io protocol.IO) {
	// TODO: Implement
}
