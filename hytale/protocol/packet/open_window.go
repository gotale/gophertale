package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type OpenWindow struct{}

func (*OpenWindow) ID() uint32 {
	return IDOpenWindow
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       OpenWindow obj = new OpenWindow();
//       byte nullBits = buf.getByte(offset);
//       obj.id = buf.getIntLE(offset + 1);
//       obj.windowType = WindowType.fromValue(buf.getByte(offset + 5));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 18 + buf.getIntLE(offset + 6);
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
//          int varPos1 = offset + 18 + buf.getIntLE(offset + 10);
//          obj.inventory = InventorySection.deserialize(buf, varPos1);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 18 + buf.getIntLE(offset + 14);
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
//       buf.writeByte(this.windowType.getValue());
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

func (pk *OpenWindow) Marshal(io protocol.IO) {
	// TODO: Implement
}
