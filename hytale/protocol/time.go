package protocol

type InstantData struct {
	Seconds int64
	Nanos   int32
}

func (x *InstantData) Marshal(io IO) {
	io.Int64(&x.Seconds)
	io.Int32(&x.Nanos)
}
