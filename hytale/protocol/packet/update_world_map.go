package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type UpdateWorldMap struct {
	Chunks         []protocol.MapChunk
	AddedMarkers   []protocol.MapMarker
	RemovedMarkers []string
}

func (*UpdateWorldMap) ID() uint32 {
	return IDUpdateWorldMap
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       UpdateWorldMap obj = new UpdateWorldMap();
//       byte nullBits = buf.getByte(offset);
//       if ((nullBits & 1) != 0) {
//          int varPos0 = offset + 13 + buf.getIntLE(offset + 1);
//          int chunksCount = VarInt.peek(buf, varPos0);
//          if (chunksCount < 0) {
//             throw ProtocolException.negativeLength("Chunks", chunksCount);
//          }
//
//          if (chunksCount > 4096000) {
//             throw ProtocolException.arrayTooLong("Chunks", chunksCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos0);
//          if (varPos0 + varIntLen + chunksCount * 9L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("Chunks", varPos0 + varIntLen + chunksCount * 9, buf.readableBytes());
//          }
//
//          obj.chunks = new MapChunk[chunksCount];
//          int elemPos = varPos0 + varIntLen;
//
//          for (int i = 0; i < chunksCount; i++) {
//             obj.chunks[i] = MapChunk.deserialize(buf, elemPos);
//             elemPos += MapChunk.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       if ((nullBits & 2) != 0) {
//          int varPos1 = offset + 13 + buf.getIntLE(offset + 5);
//          int addedMarkersCount = VarInt.peek(buf, varPos1);
//          if (addedMarkersCount < 0) {
//             throw ProtocolException.negativeLength("AddedMarkers", addedMarkersCount);
//          }
//
//          if (addedMarkersCount > 4096000) {
//             throw ProtocolException.arrayTooLong("AddedMarkers", addedMarkersCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos1);
//          if (varPos1 + varIntLen + addedMarkersCount * 38L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("AddedMarkers", varPos1 + varIntLen + addedMarkersCount * 38, buf.readableBytes());
//          }
//
//          obj.addedMarkers = new MapMarker[addedMarkersCount];
//          int elemPos = varPos1 + varIntLen;
//
//          for (int i = 0; i < addedMarkersCount; i++) {
//             obj.addedMarkers[i] = MapMarker.deserialize(buf, elemPos);
//             elemPos += MapMarker.computeBytesConsumed(buf, elemPos);
//          }
//       }
//
//       if ((nullBits & 4) != 0) {
//          int varPos2 = offset + 13 + buf.getIntLE(offset + 9);
//          int removedMarkersCount = VarInt.peek(buf, varPos2);
//          if (removedMarkersCount < 0) {
//             throw ProtocolException.negativeLength("RemovedMarkers", removedMarkersCount);
//          }
//
//          if (removedMarkersCount > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedMarkers", removedMarkersCount, 4096000);
//          }
//
//          int varIntLen = VarInt.length(buf, varPos2);
//          if (varPos2 + varIntLen + removedMarkersCount * 1L > buf.readableBytes()) {
//             throw ProtocolException.bufferTooSmall("RemovedMarkers", varPos2 + varIntLen + removedMarkersCount * 1, buf.readableBytes());
//          }
//
//          obj.removedMarkers = new String[removedMarkersCount];
//          int elemPos = varPos2 + varIntLen;
//
//          for (int i = 0; i < removedMarkersCount; i++) {
//             int strLen = VarInt.peek(buf, elemPos);
//             if (strLen < 0) {
//                throw ProtocolException.negativeLength("removedMarkers[" + i + "]", strLen);
//             }
//
//             if (strLen > 4096000) {
//                throw ProtocolException.stringTooLong("removedMarkers[" + i + "]", strLen, 4096000);
//             }
//
//             int strVarLen = VarInt.length(buf, elemPos);
//             obj.removedMarkers[i] = PacketIO.readVarString(buf, elemPos);
//             elemPos += strVarLen + strLen;
//          }
//       }
//
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       int startPos = buf.writerIndex();
//       byte nullBits = 0;
//       if (this.chunks != null) {
//          nullBits = (byte)(nullBits | 1);
//       }
//
//       if (this.addedMarkers != null) {
//          nullBits = (byte)(nullBits | 2);
//       }
//
//       if (this.removedMarkers != null) {
//          nullBits = (byte)(nullBits | 4);
//       }
//
//       buf.writeByte(nullBits);
//       int chunksOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int addedMarkersOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int removedMarkersOffsetSlot = buf.writerIndex();
//       buf.writeIntLE(0);
//       int varBlockStart = buf.writerIndex();
//       if (this.chunks != null) {
//          buf.setIntLE(chunksOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.chunks.length > 4096000) {
//             throw ProtocolException.arrayTooLong("Chunks", this.chunks.length, 4096000);
//          }
//
//          VarInt.write(buf, this.chunks.length);
//
//          for (MapChunk item : this.chunks) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(chunksOffsetSlot, -1);
//       }
//
//       if (this.addedMarkers != null) {
//          buf.setIntLE(addedMarkersOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.addedMarkers.length > 4096000) {
//             throw ProtocolException.arrayTooLong("AddedMarkers", this.addedMarkers.length, 4096000);
//          }
//
//          VarInt.write(buf, this.addedMarkers.length);
//
//          for (MapMarker item : this.addedMarkers) {
//             item.serialize(buf);
//          }
//       } else {
//          buf.setIntLE(addedMarkersOffsetSlot, -1);
//       }
//
//       if (this.removedMarkers != null) {
//          buf.setIntLE(removedMarkersOffsetSlot, buf.writerIndex() - varBlockStart);
//          if (this.removedMarkers.length > 4096000) {
//             throw ProtocolException.arrayTooLong("RemovedMarkers", this.removedMarkers.length, 4096000);
//          }
//
//          VarInt.write(buf, this.removedMarkers.length);
//
//          for (String item : this.removedMarkers) {
//             PacketIO.writeVarString(buf, item, 4096000);
//          }
//       } else {
//          buf.setIntLE(removedMarkersOffsetSlot, -1);
//       }
//    }

func (pk *UpdateWorldMap) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasChunks := opts.Has(len(pk.Chunks) > 0)
	hasAddedMarkers := opts.Has(len(pk.AddedMarkers) > 0)
	hasRemovedMarkers := opts.Has(len(pk.RemovedMarkers) > 0)
	opts.Write()

	// TODO: Fix this to correctly be offsets to different blocks
	var x int32
	io.Int32(&x)
	io.Int32(&x)
	io.Int32(&x)

	if hasChunks {
		protocol.Slice(io, &pk.Chunks, 4096000)
	}
	if hasAddedMarkers {
		protocol.Slice(io, &pk.AddedMarkers, 4096000)
	}
	if hasRemovedMarkers {
		protocol.FuncSlice(io, &pk.RemovedMarkers, 4096000, func(x *string) {
			io.VarString(x, 4096000)
		})
	}
}
