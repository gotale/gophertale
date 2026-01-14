package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateBlockHitboxes struct{}

func (*UpdateBlockHitboxes) ID() uint32 {
	return IDUpdateBlockHitboxes
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateBlockHitboxes obj = new UpdateBlockHitboxes();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.maxId = buf.getIntLE(offset + 2);
//       int pos = offset + 6;
//       if ((nullBits & 1) != 0) {
//          int blockBaseHitboxesCount = VarInt.peek(buf, pos);
//          if (blockBaseHitboxesCount < 0) {
//             throw ProtocolException.negativeLength("BlockBaseHitboxes", blockBaseHitboxesCount);
//          }
//
//          if (blockBaseHitboxesCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockBaseHitboxes", blockBaseHitboxesCount, 4096000);
//          }
//
//          pos += VarInt.size(blockBaseHitboxesCount);
//          obj.blockBaseHitboxes = new HashMap<>(blockBaseHitboxesCount);
//
//          for (int i = 0; i < blockBaseHitboxesCount; i++) {
//             int key = buf.getIntLE(pos);
//             pos += 4;
//             int valLen = VarInt.peek(buf, pos);
//             if (valLen < 0) {
//                throw ProtocolException.negativeLength("val", valLen);
//             }
//
//             if (valLen > 64) {
//                throw ProtocolException.arrayTooLong("val", valLen, 64);
//             }
//
//             int valVarLen = VarInt.length(buf, pos);
//             if (pos + valVarLen + valLen * 24L > buf.readableBytes()) {
//                throw ProtocolException.bufferTooSmall("val", pos + valVarLen + valLen * 24, buf.readableBytes());
//             }
//
//             pos += valVarLen;
//             Hitbox[] val = new Hitbox[valLen];
//
//             for (int valIdx = 0; valIdx < valLen; valIdx++) {
//                val[valIdx] = Hitbox.deserialize(buf, pos);
//                pos += Hitbox.computeBytesConsumed(buf, pos);
//             }
//
//             if (obj.blockBaseHitboxes.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("blockBaseHitboxes", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.blockBaseHitboxes != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeIntLE(this.maxId);
//       if (this.blockBaseHitboxes != null) {
//          if (this.blockBaseHitboxes.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("BlockBaseHitboxes", this.blockBaseHitboxes.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.blockBaseHitboxes.size());
//
//          for (Entry<Integer, Hitbox[]> e : this.blockBaseHitboxes.entrySet()) {
//             buf.writeIntLE(e.getKey());
//             VarInt.write(buf, e.getValue().length);
//
//             for (Hitbox arrItem : e.getValue()) {
//                arrItem.serialize(buf);
//             }
//          }
//       }
//    }

func (pk *UpdateBlockHitboxes) Marshal(io protocol.IO) {
	// TODO: Implement
}
