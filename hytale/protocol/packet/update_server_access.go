package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateServerAccess struct{}

func (*UpdateServerAccess) ID() uint32 {
	return IDUpdateServerAccess
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateServerAccess obj = new UpdateServerAccess();
//       byte nullBits = buf.getByte(offset);
//       obj.access = Access.fromValue(buf.getByte(offset + 1));
//       int pos = offset + 2;
//       if ((nullBits & 1) != 0) {
//          int hostsCount = VarInt.peek(buf, pos);
//          if (hostsCount < 0) {
//             throw ProtocolException.negativeLength("Hosts", hostsCount);
//          }
//
//          if (hostsCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Hosts", hostsCount, 4096000);
//          }
//
//          int hostsVarLen = VarInt.size(hostsCount);
//          if (pos + hostsVarLen + hostsCount * 2L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Hosts", pos + hostsVarLen + hostsCount * 2, buf.readableBytes());
//          }
//
//          pos += hostsVarLen;
//          obj.hosts = new HostAddress[hostsCount];
//
//          for (int i = 0; i < hostsCount; i++) {
//             obj.hosts[i] = HostAddress.deserialize(buf, pos);
//             pos += HostAddress.computeBytesConsumed(buf, pos);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       byte nullBits = 0;
//       if (this.hosts != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeByte(this.access.getValue());
//       if (this.hosts != null) {
//          if (this.hosts.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Hosts", this.hosts.length, 4096000);
//          }
//
//          VarInt.write(buf, this.hosts.length);
//
//          for (HostAddress item : this.hosts) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *UpdateServerAccess) Marshal(io protocol.IO) {
	// TODO: Implement
}
