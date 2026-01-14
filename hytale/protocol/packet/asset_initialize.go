package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetInitialize struct{}

func (*AssetInitialize) ID() uint32 {
	return IDAssetInitialize
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetInitialize obj = new AssetInitialize();
//       obj.size = buf.getIntLE(offset + 0);
//       int pos = offset + 4;
//       obj.asset = Asset.deserialize(buf, pos);
//       pos += Asset.computeBytesConsumed(buf, pos);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.size);
//       this.asset.serialize(buf);
//    }

func (pk *AssetInitialize) Marshal(io protocol.IO) {
	// TODO: Implement
}
