package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateEntityUIComponents struct{}

func (*UpdateEntityUIComponents) ID() uint32 {
	return IDUpdateEntityUIComponents
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateEntityUIComponents obj = new UpdateEntityUIComponents();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int componentsCount = VarInt.peek(buf, pos);
//          if (componentsCount < 0) {
//             throw ProtocolException.negativeLength("Components", componentsCount);
//          }
//
//          if (componentsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Components", componentsCount, 4096000);
//          }
//
//          pos += VarInt.size(componentsCount);
//          obj.components = new HashMap<>(componentsCount);
//
//          for (int i = 0; i < componentsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             EntityUIComponent val = EntityUIComponent.deserialize(buf, pos);
//             pos += EntityUIComponent.computeBytesConsumed(buf, pos);
//             if (obj.components.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("components", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.components != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.components != null) {
//          if (this.components.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Components", this.components.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.components.size());
//
//          for (Entry<Integer, EntityUIComponent> e : this.components.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateEntityUIComponents) Marshal(io protocol.IO) {
	// TODO: Implement
}
