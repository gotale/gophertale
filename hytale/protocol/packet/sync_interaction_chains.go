package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SyncInteractionChains struct{}

func (*SyncInteractionChains) ID() uint32 {
	return IDSyncInteractionChains
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SyncInteractionChains obj = new SyncInteractionChains();
//       int pos = offset + 0;
//       int updatesCount = VarInt.peek(buf, pos);
//       if (updatesCount < 0) {
//          throw ProtocolException.negativeLength("Updates", updatesCount);
//       } else if (updatesCount > 4096000) {
//          throw ProtocolException.arrayTooLong("Updates", updatesCount, 4096000);
//       } else {
//          int updatesVarLen = VarInt.size(updatesCount);
//          if (pos + updatesVarLen + updatesCount * 33L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Updates", pos + updatesVarLen + updatesCount * 33, buf.readableBytes());
//          } else {
//             pos += updatesVarLen;
//             obj.updates = new SyncInteractionChain[updatesCount];
//
//             for (int i = 0; i < updatesCount; i++) {
//                obj.updates[i] = SyncInteractionChain.deserialize(buf, pos);
//                pos += SyncInteractionChain.computeBytesConsumed(buf, pos);
//             }
//
//             return obj;
//          }
//       }
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       if (this.updates.length > 4096000) {
//          throw ProtocolException.arrayTooLong("Updates", this.updates.length, 4096000);
//       } else {
//          VarInt.write(buf, this.updates.length);
//
//          for (SyncInteractionChain item : this.updates) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *SyncInteractionChains) Marshal(io protocol.IO) {
	// TODO: Implement
}
