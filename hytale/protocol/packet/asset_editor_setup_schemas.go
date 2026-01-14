package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorSetupSchemas struct{}

func (*AssetEditorSetupSchemas) ID() uint32 {
	return IDAssetEditorSetupSchemas
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorSetupSchemas obj = new AssetEditorSetupSchemas();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int schemasCount = VarInt.peek(buf, pos);
//          if (schemasCount < 0) {
//             throw ProtocolException.negativeLength("Schemas", schemasCount);
//          }
//
//          if (schemasCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Schemas", schemasCount, 4096000);
//          }
//
//          int schemasVarLen = VarInt.size(schemasCount);
//          if (pos + schemasVarLen + schemasCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Schemas", pos + schemasVarLen + schemasCount * 1, buf.readableBytes());
//          }
//
//          pos += schemasVarLen;
//          obj.schemas = new SchemaFile[schemasCount];
//
//          for (int i = 0; i < schemasCount; i++) {
//             obj.schemas[i] = SchemaFile.deserialize(buf, pos);
//             pos += SchemaFile.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.schemas != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.schemas != null) {
//          if (this.schemas.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Schemas", this.schemas.length, 4096000);
//          }
//
//          VarInt.write(buf, this.schemas.length);
//
//          for (SchemaFile item : this.schemas) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *AssetEditorSetupSchemas) Marshal(io protocol.IO) {
	// TODO: Implement
}
