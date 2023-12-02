package bitstream

import (
	"encoding/binary"
	"io"
	"math"
)

type BIStream struct {
	endian binary.ByteOrder
	reader io.Reader
	err    error
}

func NewBIStream(endian binary.ByteOrder, reader io.Reader) *BIStream {
	return &BIStream{
		endian: endian,
		reader: reader,
	}
}

func (b *BIStream) catchError(f func()) *BIStream {
	if nil == b.err {
		f()
	}
	return b
}

func (b *BIStream) Error() error {
	return b.err
}

// ReadBool read 1 byte in io.Reader. Returns a bool and an error if exists
func (b *BIStream) ReadBool() (bool, error) {
	buf, err := b.ReadBytes(1)
	return buf[0] == 1, err
}

// FetchBool fetch 1 byte in io.Reader.
func (b *BIStream) FetchBool(value *bool) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadBool()
	})
}

// ReadByte read 1 byte in io.Reader. Returns a byte and an error if exists
func (b *BIStream) ReadByte() (byte, error) {
	buf, err := b.ReadBytes(1)
	return buf[0], err
}

// FetchByte fetch 1 byte in io.Reader.
func (b *BIStream) FetchByte(value *byte) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadByte()
	})
}

// ReadBytes read n bytes in io.Reader. Returns a byte array and an error if exists
func (b *BIStream) ReadBytes(n uint64) ([]byte, error) {
	buf := make([]byte, n)
	_, err := io.ReadFull(b.reader, buf)
	return buf, err
}

// FetchBytes fetch n byte in io.Reader.
func (b *BIStream) FetchBytes(value *[]byte, n uint64) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadBytes(n)
	})
}

// ReadUint8 read 1 byte in io.Reader and covert it to uint8 and an error if exists
func (b *BIStream) ReadUint8() (uint8, error) {
	return b.ReadByte()
}

// FetchUint8 fetch 1 byte in io.Reader.
func (b *BIStream) FetchUint8(value *byte) *BIStream {
	return b.FetchByte(value)
}

// ReadInt8 read 1 byte in io.Reader and covert it to int8 and an error if exists
func (b *BIStream) ReadInt8() (int8, error) {
	u, err := b.ReadUint8()
	return int8(u), err
}

// FetchInt8 fetch 1 byte in io.Reader.
func (b *BIStream) FetchInt8(value *int8) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadInt8()
	})
}

// ReadUShort read 2 byte in io.Reader and covert it to uint16 and an error if exists
func (b *BIStream) ReadUShort() (uint16, error) {
	return b.ReadUint16()
}

// FetchUShort fetch 2 byte in io.Reader.
func (b *BIStream) FetchUShort(value *uint16) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadUShort()
	})
}

// ReadUint16 read 2 byte in io.Reader and covert it to uint16 and an error if exists
func (b *BIStream) ReadUint16() (uint16, error) {
	buf, err := b.ReadBytes(2)
	if err != nil {
		return 0, err
	}
	return b.endian.Uint16(buf), nil
}

// FetchUint16 fetch 2 byte in io.Reader.
func (b *BIStream) FetchUint16(value *uint16) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadUint16()
	})
}

// ReadShort read 2 byte in io.Reader and covert it to int16 and an error if exists
func (b *BIStream) ReadShort() (int16, error) {
	return b.ReadInt16()
}

// FetchShort fetch 2 byte in io.Reader
func (b *BIStream) FetchShort(value *int16) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadShort()
	})
}

// ReadInt16 read 2 byte in io.Reader and covert it to int16 and an error if exists
func (b *BIStream) ReadInt16() (int16, error) {
	i, err := b.ReadUint16()
	return int16(i), err
}

// FetchInt16 fetch 2 byte in io.Reader
func (b *BIStream) FetchInt16(value *int16) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadInt16()
	})
}

// ReadUint32 read 4 byte in io.Reader and covert it to uint32 and an error if exists
func (b *BIStream) ReadUint32() (uint32, error) {
	buf, err := b.ReadBytes(4)
	if err != nil {
		return 0, err
	}
	return b.endian.Uint32(buf), nil
}

// FetchUint32 fetch 4 byte in io.Reader
func (b *BIStream) FetchUint32(value *uint32) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadUint32()
	})
}

// ReadInt32 read 4 byte in io.Reader and covert it to int32 and an error if exists
func (b *BIStream) ReadInt32() (int32, error) {
	i, err := b.ReadUint32()
	return int32(i), err
}

// FetchInt32 fetch 4 byte in io.Reader
func (b *BIStream) FetchInt32(value *int32) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadInt32()
	})
}

// ReadUint64 read 8 byte in io.Reader and covert it to uint64 and an error if exists
func (b *BIStream) ReadUint64() (uint64, error) {
	buf, err := b.ReadBytes(8)
	if err != nil {
		return 0, err
	}
	return b.endian.Uint64(buf), nil
}

// FetchUint64 fetch 8 byte in io.Reader
func (b *BIStream) FetchUint64(value *uint64) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadUint64()
	})
}

// ReadInt64 read 8 byte in io.Reader and covert it to int64 and an error if exists
func (b *BIStream) ReadInt64() (int64, error) {
	i, err := b.ReadUint64()
	return int64(i), err
}

// FetchInt64 fetch 8 byte in io.Reader
func (b *BIStream) FetchInt64(value *int64) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadInt64()
	})
}

// ReadFloat32 read 4 byte in io.Reader and covert it to float32 and an error if exists
func (b *BIStream) ReadFloat32() (float32, error) {
	buf, err := b.ReadUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(buf), nil
}

// FetchFloat32 fetch 4 byte in io.Reader
func (b *BIStream) FetchFloat32(value *float32) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadFloat32()
	})
}

// ReadFloat64 read 8 byte in io.Reader and covert it to float64 and an error if exists
func (b *BIStream) ReadFloat64() (float64, error) {
	buf, err := b.ReadUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(buf), nil
}

// FetchFloat64 fetch 8 byte in io.Reader
func (b *BIStream) FetchFloat64(value *float64) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadFloat64()
	})
}

// ReadString read n bytes as string length, then read n bytes, and covert it to string and an error if exists
func (b *BIStream) ReadString() (string, error) {
	buf, err := b.ReadBytesWithLengthPrefix()
	return string(buf), err
}

// FetchString fetch n bytes as string length in io.Reader
func (b *BIStream) FetchString(value *string) *BIStream {
	return b.catchError(func() {
		*value, b.err = b.ReadString()
	})
}

// ReadBytesWithLengthPrefix read n bytes as bytes length, then read n bytes, and an error if exists
func (b *BIStream) ReadBytesWithLengthPrefix() ([]byte, error) {
	bLen, err := b.ReadByte()
	if err != nil {
		return []byte{}, err
	}
	if 0xff > bLen {
		return b.ReadBytes(uint64(bLen))
	}
	wLen, err := b.ReadUint16()
	if 0xfffe > wLen {
		return b.ReadBytes(uint64(wLen))
	}

	len32, err := b.ReadUint32()
	if 0xfffffffe > len32 {
		return b.ReadBytes(uint64(len32))
	}

	len64, err := b.ReadUint64()
	return b.ReadBytes(len64)
}
