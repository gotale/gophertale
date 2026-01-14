package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdatePlayerInventory struct{}

func (*UpdatePlayerInventory) ID() uint32 {
	return IDUpdatePlayerInventory
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdatePlayerInventory obj = new UpdatePlayerInventory();
//       byte nullBits = buf.getByte(offset);
//       obj.sortType = SortType.fromValue(buf.getByte(offset + 1));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 30 + buf.getIntLE(offset + 2);
//          obj.storage = InventorySection.deserialize(buf, varPos0);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 30 + buf.getIntLE(offset + 6);
//          obj.armor = InventorySection.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 30 + buf.getIntLE(offset + 10);
//          obj.hotbar = InventorySection.deserialize(buf, varPos2);
//       }
//
//       if ((nullBits & 8) != 0) {
//          int varPos3 = offset + 30 + buf.getIntLE(offset + 14);
//          obj.utility = InventorySection.deserialize(buf, varPos3);
//       }
//
//       if ((nullBits & 16) != 0) {
//          int varPos4 = offset + 30 + buf.getIntLE(offset + 18);
//          obj.builderMaterial = InventorySection.deserialize(buf, varPos4);
//       }
//
//       if ((nullBits & 32) != 0) {
//          int varPos5 = offset + 30 + buf.getIntLE(offset + 22);
//          obj.tools = InventorySection.deserialize(buf, varPos5);
//       }
//
//       if ((nullBits & 64) != 0) {
//          int varPos6 = offset + 30 + buf.getIntLE(offset + 26);
//          obj.backpack = InventorySection.deserialize(buf, varPos6);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.storage != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.armor != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.hotbar != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       if (this.utility != null) {
//          nullBits = (byte)(nullBits | 8);
//       }
//
//       if (this.builderMaterial != null) {
//          nullBits = (byte)(nullBits | 16);
//       }
//
//       if (this.tools != null) {
//          nullBits = (byte)(nullBits | 32);
//       }
//
//       if (this.backpack != null) {
//          nullBits = (byte)(nullBits | 64);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.sortType.getValue());
//       int storageOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int armorOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int hotbarOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int utilityOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int builderMaterialOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int toolsOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int backpackOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.storage != null) {
//          buf.setIntLE(storageOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.storage.serialize(buf);
//       } else {
//          buf.setIntLE(storageOffsetSlot, -1);
//       }
//
//       if (this.armor != null) {
//          buf.setIntLE(armorOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.armor.serialize(buf);
//       } else {
//          buf.setIntLE(armorOffsetSlot, -1);
//       }
//
//       if (this.hotbar != null) {
//          buf.setIntLE(hotbarOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.hotbar.serialize(buf);
//       } else {
//          buf.setIntLE(hotbarOffsetSlot, -1);
//       }
//
//       if (this.utility != null) {
//          buf.setIntLE(utilityOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.utility.serialize(buf);
//       } else {
//          buf.setIntLE(utilityOffsetSlot, -1);
//       }
//
//       if (this.builderMaterial != null) {
//          buf.setIntLE(builderMaterialOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.builderMaterial.serialize(buf);
//       } else {
//          buf.setIntLE(builderMaterialOffsetSlot, -1);
//       }
//
//       if (this.tools != null) {
//          buf.setIntLE(toolsOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.tools.serialize(buf);
//       } else {
//          buf.setIntLE(toolsOffsetSlot, -1);
//       }
//
//       if (this.backpack != null) {
//          buf.setIntLE(backpackOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.backpack.serialize(buf);
//       } else {
//          buf.setIntLE(backpackOffsetSlot, -1);
//       }
//    }

func (pk *UpdatePlayerInventory) Marshal(io protocol.IO) {
	// TODO: Implement
}
