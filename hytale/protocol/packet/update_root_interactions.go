package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateRootInteractions struct{}

func (*UpdateRootInteractions) ID() uint32 {
	return IDUpdateRootInteractions
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateRootInteractions obj = new UpdateRootInteractions();
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
//             RootInteraction val = RootInteraction.deserialize(buf, pos);
//             pos += RootInteraction.computeBytesConsumed(buf, pos);
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
//          for (Entry<Integer, RootInteraction> e : this.interactions.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             e.getValue().serialize(buf);
//          }
//       }
//    }

func (pk *UpdateRootInteractions) Marshal(io protocol.IO) {
	// TODO: Implement
}
