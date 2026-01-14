package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type SyncPlayerPreferences struct{}

func (*SyncPlayerPreferences) ID() uint32 {
	return IDSyncPlayerPreferences
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       SyncPlayerPreferences obj = new SyncPlayerPreferences();
//       obj.showEntityMarkers = buf.getByte(offset + 0) != 0;
//       obj.armorItemsPreferredPickupLocation = PickupLocation.fromValue(buf.getByte(offset + 1));
//       obj.weaponAndToolItemsPreferredPickupLocation = PickupLocation.fromValue(buf.getByte(offset + 2));
//       obj.usableItemsItemsPreferredPickupLocation = PickupLocation.fromValue(buf.getByte(offset + 3));
//       obj.solidBlockItemsPreferredPickupLocation = PickupLocation.fromValue(buf.getByte(offset + 4));
//       obj.miscItemsPreferredPickupLocation = PickupLocation.fromValue(buf.getByte(offset + 5));
//       obj.allowNPCDetection = buf.getByte(offset + 6) != 0;
//       obj.respondToHit = buf.getByte(offset + 7) != 0;
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeByte(this.showEntityMarkers ? 1 : 0);
//       buf.writeByte(this.armorItemsPreferredPickupLocation.getValue());
//       buf.writeByte(this.weaponAndToolItemsPreferredPickupLocation.getValue());
//       buf.writeByte(this.usableItemsItemsPreferredPickupLocation.getValue());
//       buf.writeByte(this.solidBlockItemsPreferredPickupLocation.getValue());
//       buf.writeByte(this.miscItemsPreferredPickupLocation.getValue());
//       buf.writeByte(this.allowNPCDetection ? 1 : 0);
//       buf.writeByte(this.respondToHit ? 1 : 0);
//    }

func (pk *SyncPlayerPreferences) Marshal(io protocol.IO) {
	// TODO: Implement
}
