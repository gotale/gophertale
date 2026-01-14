package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AssetEditorPopupNotification struct{}

func (*AssetEditorPopupNotification) ID() uint32 {
	return IDAssetEditorPopupNotification
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AssetEditorPopupNotification obj = new AssetEditorPopupNotification();
//       byte nullBits = buf.getByte(offset);
//       obj.type = AssetEditorPopupNotificationType.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          obj.message = FormattedMessage.deserialize(buf, pos);
//          pos += FormattedMessage.computeBytesConsumed(buf, pos);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.message != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.type.getValue());
//       if (this.message != null) {
//          this.message.serialize(buf);
//       }
//    }

func (pk *AssetEditorPopupNotification) Marshal(io protocol.IO) {
	// TODO: Implement
}
