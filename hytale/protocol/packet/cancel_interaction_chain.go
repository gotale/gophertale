package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type CancelInteractionChain struct{}

func (*CancelInteractionChain) ID() uint32 {
	return IDCancelInteractionChain
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       CancelInteractionChain obj = new CancelInteractionChain();
//       byte nullBits = buf.getByte(offset);
//       obj.chainId = buf.getIntLE(offset + 1);
//       int pos = offset + 5;
//       if ((nullBits & 1) != 0) {
//          obj.forkedId = ForkedChainId.deserialize(buf, pos);
//          pos += ForkedChainId.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.forkedId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.chainId);
//       if (this.forkedId != null) {
//          this.forkedId.serialize(buf);
//       }
//    }

func (pk *CancelInteractionChain) Marshal(io protocol.IO) {
	// TODO: Implement
}
