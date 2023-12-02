package bitstream

import (
	"encoding/binary"
	"io"
	"math"
)

type BOStream struct {
	endian binary.ByteOrder
	writer io.Writer
	err    error
}

func NewBOStream(endian binary.ByteOrder, writer io.Writer) *BOStream {
	return &BOStream{
		endian: endian,
		writer: writer,
	}
}

func (p *BOStream) catchError(f func()) *BOStream {
	if nil == p.err {
		f()
	}
	return p
}

func (p *BOStream) Error() error {
	return p.err
}

// WriteBool write bool in io.Writer.
func (p *BOStream) WriteBool(b bool) *BOStream {
	if b {
		return p.WriteBytes([]byte{1})
	}
	return p.WriteBytes([]byte{0})
}

// WriteByte write 1 byte in io.Writer.
func (p *BOStream) WriteByte(b byte) *BOStream {
	return p.WriteBytes([]byte{b})
}

// WriteBytes write the bytes in io.Writer.
func (p *BOStream) WriteBytes(bytes []byte) *BOStream {
	return p.catchError(func() {
		_, p.err = p.writer.Write(bytes)
	})
}

// WriteUint8 write an uint8 in io.Writer
func (p *BOStream) WriteUint8(n uint8) *BOStream {
	return p.WriteByte(n)
}

// WriteInt8 write an int8 in io.Writer
func (p *BOStream) WriteInt8(n int8) *BOStream {
	return p.WriteByte(byte(n))
}

// WriteUShort write an unsigned short in io.Writer
func (p *BOStream) WriteUShort(n uint16) *BOStream {
	return p.WriteUint16(n)
}

// WriteUint16 write an uint16 in io.Writer
func (p *BOStream) WriteUint16(n uint16) *BOStream {
	return p.catchError(func() {
		buf := make([]byte, 2)
		p.endian.PutUint16(buf, n)
		_, p.err = p.writer.Write(buf)
	})
}

// WriteShort write a short in io.Writer
func (p *BOStream) WriteShort(n int16) *BOStream {
	return p.WriteInt16(n)
}

// WriteInt16 write an int16 in io.Writer
func (p *BOStream) WriteInt16(n int16) *BOStream {
	return p.WriteUint16(uint16(n))
}

// WriteUint32 write an uint32 in io.Writer
func (p *BOStream) WriteUint32(n uint32) *BOStream {
	return p.catchError(func() {
		buf := make([]byte, 4)
		p.endian.PutUint32(buf, n)
		_, p.err = p.writer.Write(buf)
	})
}

// WriteInt32 write an int32 in io.Writer
func (p *BOStream) WriteInt32(n int32) *BOStream {
	return p.WriteUint32(uint32(n))
}

// WriteUint64 write an uint64 in io.Writer
func (p *BOStream) WriteUint64(n uint64) *BOStream {
	return p.catchError(func() {
		buf := make([]byte, 8)
		p.endian.PutUint64(buf, n)
		_, p.err = p.writer.Write(buf)
	})
}

// WriteInt64 write an int64 in io.Writer
func (p *BOStream) WriteInt64(n int64) *BOStream {
	return p.WriteUint64(uint64(n))
}

// WriteFloat32 write 4 byte in io.Writer
func (p *BOStream) WriteFloat32(n float32) *BOStream {
	return p.WriteUint32(math.Float32bits(n))
}

// WriteFloat64 write 8 byte in io.Writer
func (p *BOStream) WriteFloat64(n float64) *BOStream {
	return p.WriteUint64(math.Float64bits(n))
}

// WriteString write n bytes as string length, then write n bytes
func (p *BOStream) WriteString(str string) *BOStream {
	return p.WriteBytesWithLengthPrefix([]byte(str))
}

// WriteBytesWithLengthPrefix write n bytes as bytes length, then write n bytes
func (p *BOStream) WriteBytesWithLengthPrefix(bytes []byte) *BOStream {
	return p.catchError(func() {
		bLen := len(bytes)
		if 0xff > bLen {
			p.WriteByte(byte(bLen))
		} else if 0xfffe > bLen {
			p.WriteByte(0xff).WriteUint16(uint16(bLen))
		} else if 0xfffffffe > bLen {
			p.WriteByte(0xff).WriteUint16(0xffff).WriteUint32(uint32(bLen))
		} else {
			p.WriteByte(0xff).WriteUint16(0xffff).WriteUint32(0xffffffff).WriteUint64(uint64((bLen)))
		}
		p.WriteBytes(bytes)
	})
}
