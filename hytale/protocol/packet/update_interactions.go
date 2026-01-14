package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateInteractions struct{}

func (*UpdateInteractions) ID() uint32 {
	return IDUpdateInteractions
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateInteractions obj = new UpdateInteractions();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int interactionsCount = VarInt.peek(buf, pos);
//          if (interactionsCount < 0) {
//             throw ProtocolException.negativeLength("Interactions", interactionsCount);
//          }
//
//          if (interactionsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Interactions", interactionsCount, 4096000);
//          }
//
//          pos += VarInt.size(interactionsCount);
//          obj.interactions = new HashMap<>(interactionsCount);
//
//          for (int i = 0; i < interactionsCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             Interaction val = Interaction.deserialize(buf, pos);
//             pos += Interaction.computeBytesConsumed(buf, pos);
//             if (obj.interactions.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("interactions", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.interactions != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.interactions != null) {
//          if (this.interactions.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Interactions", this.interactions.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.interactions.size());
//
//          for (Entry<Integer, Interaction> e : this.interactions.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serializeWithTypeId(buf);
//          }
//       }
//    }

func (pk *UpdateInteractions) Marshal(io protocol.IO) {
	// TODO: Implement
}
