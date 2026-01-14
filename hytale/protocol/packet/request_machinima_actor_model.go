package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type RequestMachinimaActorModel struct{}

func (*RequestMachinimaActorModel) ID() uint32 {
	return IDRequestMachinimaActorModel
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       RequestMachinimaActorModel obj = new RequestMachinimaActorModel();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 13 + buf.getIntLE(offset + 1);
//          int modelIdLen = VarInt.peek(buf, varPos0);
//          if (modelIdLen < 0) {
//             throw ProtocolException.negativeLength("ModelId", modelIdLen);
//          }
//
//          if (modelIdLen > 4096000) {
//             throw ProtocolException.stringTooLong("ModelId", modelIdLen, 4096000);
//          }
//
//          obj.modelId = PacketIO.readVarString(buf, varPos0, PacketIO.UTF8);
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 13 + buf.getIntLE(offset + 5);
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
//          int varPos2 = offset + 13 + buf.getIntLE(offset + 9);
//          int actorNameLen = VarInt.peek(buf, varPos2);
//          if (actorNameLen < 0) {
//             throw ProtocolException.negativeLength("ActorName", actorNameLen);
//          }
//
//          if (actorNameLen > 4096000) {
//             throw ProtocolException.stringTooLong("ActorName", actorNameLen, 4096000);
//          }
//
//          obj.actorName = PacketIO.readVarString(buf, varPos2, PacketIO.UTF8);
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.modelId != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.sceneName != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.actorName != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       int modelIdOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int sceneNameOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int actorNameOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.modelId != null) {
//          buf.setIntLE(modelIdOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.modelId, 4096000);
//       } else {
//          buf.setIntLE(modelIdOffsetSlot, -1);
//       }
//
//       if (this.sceneName != null) {
//          buf.setIntLE(sceneNameOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.sceneName, 4096000);
//       } else {
//          buf.setIntLE(sceneNameOffsetSlot, -1);
//       }
//
//       if (this.actorName != null) {
//          buf.setIntLE(actorNameOffsetSlot, buf.writerIndex() - varBlockStart);
//          PacketIO.writeVarString(buf, this.actorName, 4096000);
//       } else {
//          buf.setIntLE(actorNameOffsetSlot, -1);
//       }
//    }

func (pk *RequestMachinimaActorModel) Marshal(io protocol.IO) {
	// TODO: Implement
}
