package protocol

type HostAddress struct {
	Host string
	Port int16
}

func (x *HostAddress) Marshal(io IO) {
	io.Int16(&x.Port)
	io.VarString(&x.Host, 256)
}
