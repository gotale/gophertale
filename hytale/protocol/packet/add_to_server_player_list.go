package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type AddToServerPlayerList struct{}

func (*AddToServerPlayerList) ID() uint32 {
	return IDAddToServerPlayerList
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       AddToServerPlayerList obj = new AddToServerPlayerList();
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
//          if (pos + playersVarLen + playersCount * 37L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Players", pos + playersVarLen + playersCount * 37, buf.readableBytes());
//          }
//
//          pos += playersVarLen;
//          obj.players = new ServerPlayerListPlayer[playersCount];
//
//          for (int i = 0; i < playersCount; i++) {
//             obj.players[i] = ServerPlayerListPlayer.deserialize(buf, pos);
//             pos += ServerPlayerListPlayer.computeBytesConsumed(buf, pos);
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
//          for (ServerPlayerListPlayer item : this.players) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *AddToServerPlayerList) Marshal(io protocol.IO) {
	// TODO: Implement
}
