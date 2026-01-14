package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateServerPlayerListPing struct{}

func (*UpdateServerPlayerListPing) ID() uint32 {
	return IDUpdateServerPlayerListPing
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateServerPlayerListPing obj = new UpdateServerPlayerListPing();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int playersCount = VarInt.peek(buf, pos);
//          if (playersCount < 0) {
//             throw ProtocolException.negativeLength("Players", playersCount);
//          }
//
//          if (playersCount > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Players", playersCount, 4096000);
//          }
//
//          pos += VarInt.size(playersCount);
//          obj.players = new HashMap<>(playersCount);
//
//          for (int i = 0; i < playersCount; i++) {
//             UUID key = PacketIO.readUUID(buf, pos);
//             pos += 16;
//             int val = buf.getIntLE(pos);
//             pos += 4;
//             if (obj.players.put(key, val) != null) {
//                throw ProtocolException.duplicateKey("players", key);
//             }
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.players != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       if (this.players != null) {
//          if (this.players.size() > 4096000) {
//             throw ProtocolException.dictionaryTooLarge("Players", this.players.size(), 4096000);
//          }
//
//          VarInt.write(buf, this.players.size());
//
//          for (Entry<UUID, Integer> e : this.players.entrySet()) {
//             PacketIO.writeUUID(buf, e.getKey());
//             buf.writeIntLE(e.getValue());
//          }
//       }
//    }

func (pk *UpdateServerPlayerListPing) Marshal(io protocol.IO) {
	// TODO: Implement
}
