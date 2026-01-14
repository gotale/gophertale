package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type ServerSetFluids struct{}

func (*ServerSetFluids) ID() uint32 {
	return IDServerSetFluids
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       ServerSetFluids obj = new ServerSetFluids();
//       obj.x = buf.getIntLE(offset + 0);
//       obj.y = buf.getIntLE(offset + 4);
//       obj.z = buf.getIntLE(offset + 8);
//       int pos = offset + 12;
//       int cmdsCount = VarInt.peek(buf, pos);
//       if (cmdsCount < 0) {
//          throw ProtocolException.negativeLength("Cmds", cmdsCount);
//       } else if (cmdsCount > 4096000) {
//          throw ProtocolException.arrayTooLong("Cmds", cmdsCount, 4096000);
//       } else {
//          int cmdsVarLen = VarInt.size(cmdsCount);
//          if (pos + cmdsVarLen + cmdsCount * 7L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Cmds", pos + cmdsVarLen + cmdsCount * 7, buf.readableBytes());
//          } else {
//             pos += cmdsVarLen;
//             obj.cmds = new SetFluidCmd[cmdsCount];
//
//             for (int i = 0; i < cmdsCount; i++) {
//                obj.cmds[i] = SetFluidCmd.deserialize(buf, pos);
//                pos += SetFluidCmd.computeBytesConsumed(buf, pos);
//             }
//
//             return obj;
//          }
//       }
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.x);
//       buf.writeIntLE(this.y);
//       buf.writeIntLE(this.z);
//       if (this.cmds.length > 4096000) {
//          throw ProtocolException.arrayTooLong("Cmds", this.cmds.length, 4096000);
//       } else {
//          VarInt.write(buf, this.cmds.length);
//
//          for (SetFluidCmd item : this.cmds) {
//             item.serialize(buf);
//          }
//       }
//    }

func (pk *ServerSetFluids) Marshal(io protocol.IO) {
	// TODO: Implement
}
