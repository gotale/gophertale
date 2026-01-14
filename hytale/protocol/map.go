package protocol

type MapChunk struct{}

func (x *MapChunk) Marshal(io IO) {}

type MapMarker struct{}

func (x *MapMarker) Marshal(io IO) {}
