package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateItems struct{}

func (*UpdateItems) ID() uint32 {
	return IDUpdateItems
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateItems obj = new UpdateItems();
//       byte nullBits = buf.getByte(offset);
//       obj.type = UpdateType.fromValue(buf.getByte(offset + 1));
//       obj.updateModels = buf.getByte(offset + 2) != 0;
//       obj.updateIcons = buf.getByte(offset + 3) != 0;
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 12 + buf.getIntLE(offset + 4);
//          int itemsCount = VarInt.peek(buf, varPos0);
//          if (itemsCount < 0) {
//             throw ProtocolException.negativeLength("Items", itemsCount);
//          }
//
//          if (itemsCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Items", itemsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          obj.items = new HashMap<>(itemsCount);
//          int dictPos = varPos0 + varIntLen;
//
//          for (int i = 0; i < itemsCount; i++) {
//             int keyLen = VarInt.peek(buf, dictPos);
//             if (keyLen < 0) {
//                throw ProtocolException.negativeLength("key", keyLen);
//             }
//
//             if (keyLen > 4096000) {
//                throw ProtocolException.stringTooLong("key", keyLen, 4096000);
//             }
//
//             int keyVarLen = VarInt.length(buf, dictPos);
//             String key = PacketIO.readVarString(buf, dictPos);
//             dictPos += keyVarLen + keyLen;
//             ItemBase val = ItemBase.deserialize(buf, dictPos);
//             dictPos += ItemBase.computeBytesConsumed(buf, dictPos);
//             if (obj.items.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("items", key);
//             }
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 12 + buf.getIntLE(offset + 8);
//          int removedItemsCount = VarInt.peek(buf, varPos1);
//          if (removedItemsCount < 0) {
//             throw ProtocolException.negativeLength("RemovedItems", removedItemsCount);
//          }
//
//          if (removedItemsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedItems", removedItemsCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + removedItemsCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("RemovedItems", varPos1 + varIntLen + removedItemsCount * 1, buf.readableBytes());
//          }
//
//          obj.removedItems = new String[removedItemsCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < removedItemsCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("removedItems[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("removedItems[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.removedItems[i] = PacketIO.readVarString(buf, elemPos);
//             elemPos += strVarLen + strLen;
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.items != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.removedItems != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       buf.writeByte(this.updateModels ? 1 : 0);
//       buf.writeByte(this.updateIcons ? 1 : 0);
//       int itemsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int removedItemsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.items != null) {
//          buf.setIntLE(itemsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.items.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Items", this.items.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.items.size());
//
//          for (Entry<String, ItemBase> e : this.items.entrySet()) {
//             PacketIO.writeVarString(buf, e.getKey(), 4096000);
//             e.getValue().serialize(buf);
//          }
//       } else {
//          buf.setIntLE(itemsOffsetSlot, -1);
//       }
//
//       if (this.removedItems != null) {
//          buf.setIntLE(removedItemsOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.removedItems.length > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedItems", this.removedItems.length, 4096000);
//          }
//
//          VarInt.write(buf, this.removedItems.length);
//
//          for (String item : this.removedItems) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(removedItemsOffsetSlot, -1);
//       }
//    }

func (pk *UpdateItems) Marshal(io protocol.IO) {
	// TODO: Implement
}
