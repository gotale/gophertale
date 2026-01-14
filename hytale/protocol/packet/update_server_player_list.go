package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateServerPlayerList struct{}

func (*UpdateServerPlayerList) ID() uint32 {
	return IDUpdateServerPlayerList
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateServerPlayerList obj = new UpdateServerPlayerList();
//       byte nullBits = buf.getByte(offset);
//       int pos = offset + 1;
//       if ((nullBits & 1) != 0) {
//          int playersCount = VarInt.peek(buf, pos);
//          if (playersCount < 0) {
//             throw ProtocolException.negativeLength("Players", playersCount);
//          }
//
//          if (playersCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Players", playersCount, 4096000);
//          }
//
//          int playersVarLen = VarInt.size(playersCount);
//          if (pos + playersVarLen + playersCount * 32L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Players", pos + playersVarLen + playersCount * 32, buf.readableBytes());
//          }
//
//          pos += playersVarLen;
//          obj.players = new ServerPlayerListUpdate[playersCount];
//
//          for (int i = 0; i < playersCount; i++) {
//             obj.players[i] = ServerPlayerListUpdate.deserialize(buf, pos);
//             pos += ServerPlayerListUpdate.computeBytesConsumed(buf, pos);
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
//          if (this.players.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Players", this.players.length, 4096000);
//          }
//
//          VarInt.write(buf, this.players.length);
//
//          for (ServerPlayerListUpdate item : this.players) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *UpdateServerPlayerList) Marshal(io protocol.IO) {
	// TODO: Implement
}
