package protocol

import "fmt"

type Optionals struct {
	io      IO
	val     uint8
	bit     uint8
	reading bool
	written bool
}

func NewOptionals(io IO) *Optionals {
	opts := &Optionals{
		io:      io,
		reading: io.Reading(),
	}
	if opts.reading {
		io.Uint8(&opts.val)
	}
	return opts
}

func (opts *Optionals) Has(present bool) bool {
	if opts.bit >= 8 {
		panic(fmt.Errorf("has exceeds 8 bits"))
	}
	mask := uint8(1) << opts.bit
	opts.bit++
	if opts.reading {
		return (opts.val & mask) != 0
	} else if present {
		opts.val |= mask
	}
	return present
}

func (opts *Optionals) Write() {
	if opts.reading {
		panic(fmt.Errorf("cannot write optionals on a reader"))
	} else if opts.written {
		panic(fmt.Errorf("already written optionals"))
	}
	opts.written = true
	opts.io.Uint8(&opts.val)
}
