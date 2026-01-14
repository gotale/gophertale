package protocol

import (
	"fmt"
	"image/color"

	"github.com/google/uuid"
)

type IO interface {
	Uint16(x *uint16)
	Int16(x *int16)
	Uint32(x *uint32)
	Int32(x *int32)
	Uint64(x *uint64)
	Int64(x *int64)
	Float32(x *float32)
	Uint8(x *uint8)
	Int8(x *int8)
	Bool(x *bool)
	Varint64(x *int64)
	Varuint64(x *uint64)
	Varint32(x *int32)
	Varuint32(x *uint32)

	RGB(x *color.RGBA)
	RGBA(x *color.RGBA)

	FixedString(x *string, length int)
	FixedAsciiString(x *string, length int)
	VarString(x *string, maxLength int)
	VarStringAscii(x *string, maxLength int)
	ByteSlice(x *[]byte, maxLength int)

	UUID(x *uuid.UUID)

	Reading() bool
}

// Marshaler is a type that can be written to or read from an IO.
type Marshaler interface {
	Marshal(io IO)
}

// Slice reads/writes a slice of T.
func Slice[T any, S ~*[]T, A PtrMarshaler[T]](io IO, x S, maxLen int32) {
	count := int32(len(*x))
	io.Varint32(&count) // TODO: Check if signed
	SliceOfLen[T, S, A](io, count, x, maxLen)
}

// FuncSlice reads/writes a slice of T using function f.
func FuncSlice[T any, S ~*[]T](io IO, x S, maxLen int32, f func(*T)) {
	count := int32(len(*x))
	io.Varint32(&count) // TODO: Check if signed
	FuncSliceOfLen(io, count, x, maxLen, f)
}

// FuncIOSlice reads/writes a slice of T using a function f.
func FuncIOSlice[T any, S ~*[]T](io IO, x S, maxLen int32, f func(IO, *T)) {
	FuncSlice(io, x, maxLen, func(v *T) {
		f(io, v)
	})
}

// SliceOfLen reads/writes the elements of a slice of type T with length l.
func SliceOfLen[T any, S ~*[]T, A PtrMarshaler[T]](io IO, l int32, x S, maxLen int32) {
	if l > maxLen {
		panic(fmt.Errorf("slice max length exceeded: %d > %d", l, maxLen))
	} else if io.Reading() {
		*x = make([]T, l)
	}

	for i := int32(0); i < l; i++ {
		A(&(*x)[i]).Marshal(io)
	}
}

// FuncSliceOfLen reads/writes the elements of a slice of type T with length l using func f.
func FuncSliceOfLen[T any, S ~*[]T](io IO, l int32, x S, maxLen int32, f func(*T)) {
	if l > maxLen {
		panic(fmt.Errorf("slice max length exceeded: %d > %d", l, maxLen))
	} else if io.Reading() {
		*x = make([]T, l)
	}

	for i := int32(0); i < l; i++ {
		f(&(*x)[i])
	}
}

// Map reads/writes a map[K]V where K and V are PtrMarshaler types.
func Map[K comparable, V any, M ~*map[K]V, A PtrMarshaler[V]](io IO, x M, maxLen int32, fk func(*K)) {
	count := int32(len(*x))
	io.Varint32(&count) // TODO: Check if signed
	MapOfLen[K, V, M, A](io, count, x, maxLen, fk)
}

// FuncMap reads/writes a map[K]V using functions fk/fv for key/value.
func FuncMap[K comparable, V any, M ~*map[K]V](io IO, x M, maxLen int32, fk func(*K), fv func(*V)) {
	count := int32(len(*x))
	io.Varint32(&count) // TODO: Check if signed
	FuncMapOfLen(io, count, x, maxLen, fk, fv)
}

// FuncIOMap reads/writes a map[K]V using functions that receive IO.
func FuncIOMap[K comparable, V any, M ~*map[K]V](io IO, x M, maxLen int32, fk func(*K), fv func(IO, *V)) {
	FuncMap(io, x, maxLen, fk, func(v *V) {
		fv(io, v)
	})
}

// MapOfLen reads/writes exactly l entries of a map[K]V with marshalers.
func MapOfLen[K comparable, V any, M ~*map[K]V, A PtrMarshaler[V]](io IO, l int32, x M, maxLen int32, fk func(*K)) {
	if l > maxLen {
		panic(fmt.Errorf("map max length exceeded: %d > %d", l, maxLen))
	} else if io.Reading() {
		*x = make(map[K]V, l)
		for i := int32(0); i < l; i++ {
			var k K
			var v V
			fk(&k)
			A(&v).Marshal(io)
			(*x)[k] = v
		}
		return
	}

	for k, v := range *x {
		kCopy := k
		vCopy := v
		fk(&kCopy)
		A(&vCopy).Marshal(io)
	}
}

// FuncMapOfLen reads/writes exactly l entries of a map[K]V using fk/fv.
func FuncMapOfLen[K comparable, V any, M ~*map[K]V](io IO, l int32, x M, maxLen int32, fk func(*K), fv func(*V)) {
	if l > maxLen {
		panic(fmt.Errorf("map max length exceeded: %d > %d", l, maxLen))
	} else if io.Reading() {
		*x = make(map[K]V, l)
		for i := int32(0); i < l; i++ {
			var k K
			var v V
			fk(&k)
			fv(&v)
			(*x)[k] = v
		}
		return
	}

	for k, v := range *x {
		kCopy := k
		vCopy := v
		fk(&kCopy)
		fv(&vCopy)
	}
}

// PtrMarshaler represents a type that implements Marshaler for its pointer.
type PtrMarshaler[T any] interface {
	Marshaler
	*T
}

// Single reads/writes a single Marshaler x.
func Single[T any, S PtrMarshaler[T]](io IO, x S) {
	x.Marshal(io)
}
