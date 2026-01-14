package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

type BuilderToolsSetSoundSet struct{}

func (*BuilderToolsSetSoundSet) ID() uint32 {
	return IDBuilderToolsSetSoundSet
}

// Hytale deserialize logic:
// deserialize(@Nonnull ByteBuf buf, int offset) {
//       BuilderToolsSetSoundSet obj = new BuilderToolsSetSoundSet();
//       obj.soundSetIndex = buf.getIntLE(offset + 0);
//       return obj;
//    }

// Hytale serialize logic:
// serialize(@Nonnull ByteBuf buf) {
//       buf.writeIntLE(this.soundSetIndex);
//    }

func (pk *BuilderToolsSetSoundSet) Marshal(io protocol.IO) {
	// TODO: Implement
}
