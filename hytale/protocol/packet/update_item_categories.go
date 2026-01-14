package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateItemCategories struct{}

func (*UpdateItemCategories) ID() uint32 {
	return IDUpdateItemCategories
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateItemCategories obj = new UpdateItemCategories();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int itemCategoriesCount = VarInt.peek(buf, pos);
//          if (itemCategoriesCount < 0) {
//             throw ProtocolException.negativeLength("ItemCategories", itemCategoriesCount);
//          }
//
//          if (itemCategoriesCount > 4096000) {
//             throw ProtocolException.arrayTooLong("ItemCategories", itemCategoriesCount, 4096000);
//          }
//
//          int itemCategoriesVarLen = VarInt.size(itemCategoriesCount);
//          if (pos + itemCategoriesVarLen + itemCategoriesCount * 6L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("ItemCategories", pos + itemCategoriesVarLen + itemCategoriesCount * 6, buf.readableBytes());
//          }
//
//          pos += itemCategoriesVarLen;
//          obj.itemCategories = new ItemCategory[itemCategoriesCount];
//
//          for (int i = 0; i < itemCategoriesCount; i++) {
//             obj.itemCategories[i] = ItemCategory.deserialize(buf, pos);
//             pos += ItemCategory.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.itemCategories != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.itemCategories != null) {
//          if (this.itemCategories.length > 4096000) {
//             throw ProtocolException.arrayTooLong("ItemCategories", this.itemCategories.length, 4096000);
//          }
//
//          VarInt.write(buf, this.itemCategories.length);
//
//          for (ItemCategory item : this.itemCategories) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *UpdateItemCategories) Marshal(io protocol.IO) {
	// TODO: Implement
}
