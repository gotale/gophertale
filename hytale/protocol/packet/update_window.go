package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateWindow struct{}

func (*UpdateWindow) ID() uint32 {
	return IDUpdateWindow
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateWindow obj = new UpdateWindow();
//       byte nullBits = buf.getByte(offset);
//       obj.id = buf.getIntLE(offset + 1);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 17 + buf.getIntLE(offset + 5);
//          int windowDataLen = VarInt.peek(buf, varPos0);
//          if (windowDataLen < 0) {
//             throw ProtocolException.negativeLength("WindowData", windowDataLen);
//          }
//
//          if (windowDataLen > 4096000) {
//             throw ProtocolException.stringTooLong("WindowData", windowDataLen, 4096000);
//          }
//
//          obj.windowData = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 17 + buf.getIntLE(offset + 9);
//          obj.inventory = InventorySection.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 17 + buf.getIntLE(offset + 13);
//          obj.extraResources = ExtraResources.deserialize(buf, varPos2);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.windowData != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.inventory != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.extraResources != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeIntLE(this.id);
//       int windowDataOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int inventoryOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int extraResourcesOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.windowData != null) {
//          buf.setIntLE(windowDataOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.windowData, 4096000);
//       } else {
//          buf.setIntLE(windowDataOffsetSlot, -1);
//       }
//
//       if (this.inventory != null) {
//          buf.setIntLE(inventoryOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.inventory.serialize(buf);
//       } else {
//          buf.setIntLE(inventoryOffsetSlot, -1);
//       }
//
//       if (this.extraResources != null) {
//          buf.setIntLE(extraResourcesOffsetSlot, buf.writerIndex() - varBlockStart);
//          this.extraResources.serialize(buf);
//       } else {
//          buf.setIntLE(extraResourcesOffsetSlot, -1);
//       }
//    }

func (pk *UpdateWindow) Marshal(io protocol.IO) {
	// TODO: Implement
}
