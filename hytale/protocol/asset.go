package protocol

type Asset struct {
	Hash string
	Name string
}

func (x *Asset) Marshal(io IO) {
	io.FixedAsciiString(&x.Hash, 64)
	io.VarStringUTF8(&x.Name, 512)
}
