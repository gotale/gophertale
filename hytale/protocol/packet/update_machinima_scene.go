package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateMachinimaScene struct{}

func (*UpdateMachinimaScene) ID() uint32 {
	return IDUpdateMachinimaScene
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateMachinimaScene obj = new UpdateMachinimaScene();
//       byte nullBits = buf.getByte(offset);
//       obj.frame = buf.getFloatLE(offset + 1);
//       obj.updateType = SceneUpdateType.fromValue(buf.getByte(offset + 5));
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 18 + buf.getIntLE(offset + 6);
//          int playerLen = VarInt.peek(buf, varPos0);
//          if (playerLen < 0) {
//             throw ProtocolException.negativeLength("Player", playerLen);
//          }
//
//          if (playerLen > 4096000) {
//             throw ProtocolException.stringTooLong("Player", playerLen, 4096000);
//          }
//
//          obj.player = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 18 + buf.getIntLE(offset + 10);
//          int sceneNameLen = VarInt.peek(buf, varPos1);
//          if (sceneNameLen < 0) {
//             throw ProtocolException.negativeLength("SceneName", sceneNameLen);
//          }
//
//          if (sceneNameLen > 4096000) {
//             throw ProtocolException.stringTooLong("SceneName", sceneNameLen, 4096000);
//          }
//
//          obj.sceneName = PacketIO.readVarString(buf, varPos1, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 18 + buf.getIntLE(offset + 14);
//          int sceneCount = VarInt.peek(buf, varPos2);
//          if (sceneCount < 0) {
//             throw ProtocolException.negativeLength("Scene", sceneCount);
//          }
//
//          if (sceneCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Scene", sceneCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos2);
//          if (varPos2 + varIntLen + sceneCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Scene", varPos2 + varIntLen + sceneCount * 1, buf.readableBytes());
//          }
//
//          obj.scene = new byte[sceneCount];
//
//          for (int i = 0; i < sceneCount; i++) {
//             obj.scene[i] = buf.getByte(varPos2 + varIntLen + i * 1);
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.player != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.sceneName != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.scene != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       buf.writeFloatLE(this.frame);
//       buf.writeByte(this.updateType.getValue());
//       int playerOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int sceneNameOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int sceneOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.player != null) {
//          buf.setIntLE(playerOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.player, 4096000);
//       } else {
//          buf.setIntLE(playerOffsetSlot, -1);
//       }
//
//       if (this.sceneName != null) {
//          buf.setIntLE(sceneNameOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.sceneName, 4096000);
//       } else {
//          buf.setIntLE(sceneNameOffsetSlot, -1);
//       }
//
//       if (this.scene != null) {
//          buf.setIntLE(sceneOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.scene.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Scene", this.scene.length, 4096000);
//          }
//
//          VarInt.write(buf, this.scene.length);
//
//          for (byte item : this.scene) {
//             buf.writeByte(item);
//          }
//       } else {
//          buf.setIntLE(sceneOffsetSlot, -1);
//       }
//    }

func (pk *UpdateMachinimaScene) Marshal(io protocol.IO) {
	// TODO: Implement
}
