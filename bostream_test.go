package bitstream

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
	"testing"
)

func TestBOStream_WriteBool(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      bool
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteBool", endian: binary.LittleEndian, args: true, want: []byte{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteBool(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteBool() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteBool() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteByte(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      byte
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteByte", endian: binary.LittleEndian, args: 0x01, want: []byte{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteByte(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteByte() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteByte() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteBytes(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteBytes", endian: binary.LittleEndian, args: []byte{1, 3}, want: []byte{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteBytes(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteBytes() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteBytes() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteBytesWithLengthPrefix(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteBytesWithLengthPrefix", endian: binary.LittleEndian, args: []byte{1, 3}, want: []byte{2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteBytesWithLengthPrefix(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteBytesWithLengthPrefix() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteBytesWithLengthPrefix() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteFloat32(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      float32
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteFloat32_BigEndian", endian: binary.BigEndian, args: math.SmallestNonzeroFloat32, want: []byte{0, 0, 0, 1}},
		{name: "TestBOStream_WriteFloat32_LittleEndian", endian: binary.LittleEndian, args: math.SmallestNonzeroFloat32, want: []byte{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteFloat32(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteFloat32() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteFloat32() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteFloat64(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      float64
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteFloat64_BigEndian", endian: binary.BigEndian, args: math.SmallestNonzeroFloat64, want: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "TestBOStream_WriteFloat64_LittleEndian", endian: binary.LittleEndian, args: math.SmallestNonzeroFloat64, want: []byte{1, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteFloat64(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteFloat64() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteFloat64() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteInt16(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      int16
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteInt16_BigEndian_1", endian: binary.BigEndian, args: 1, want: []byte{0, 1}},
		{name: "TestBOStream_WriteInt16_LittleEndian_1", endian: binary.LittleEndian, args: 1, want: []byte{1, 0}},
		{name: "TestBOStream_WriteInt16_BigEndian_10-", endian: binary.BigEndian, args: -10, want: []byte{255, 246}},
		{name: "TestBOStream_WriteInt16_LittleEndian_10-", endian: binary.LittleEndian, args: -10, want: []byte{246, 255}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteInt16(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteInt16() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteInt16() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteInt32(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      int32
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteInt32_BigEndian_1", endian: binary.BigEndian, args: 1, want: []byte{0, 0, 0, 1}},
		{name: "TestBOStream_WriteInt32_LittleEndian_1", endian: binary.LittleEndian, args: 1, want: []byte{1, 0, 0, 0}},
		{name: "TestBOStream_WriteInt32_BigEndian_10-", endian: binary.BigEndian, args: -10, want: []byte{255, 255, 255, 246}},
		{name: "TestBOStream_WriteInt32_LittleEndian_10-", endian: binary.LittleEndian, args: -10, want: []byte{246, 255, 255, 255}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteInt32(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteInt32() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteInt32() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteInt64(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      int64
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteInt64_BigEndian_1", endian: binary.BigEndian, args: 1, want: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "TestBOStream_WriteInt64_LittleEndian_1", endian: binary.LittleEndian, args: 1, want: []byte{1, 0, 0, 0, 0, 0, 0, 0}},
		{name: "TestBOStream_WriteInt64_BigEndian_10-", endian: binary.BigEndian, args: -10, want: []byte{255, 255, 255, 255, 255, 255, 255, 246}},
		{name: "TestBOStream_WriteInt64_LittleEndian_10-", endian: binary.LittleEndian, args: -10, want: []byte{246, 255, 255, 255, 255, 255, 255, 255}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteInt64(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteInt64() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteInt64() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteInt8(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      int8
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteInt8", endian: binary.LittleEndian, args: 0x01, want: []byte{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteInt8(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteInt8() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteInt8() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteShort(t *testing.T) {
	TestBOStream_WriteInt16(t)
}

func TestBOStream_WriteString(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      string
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteString_golang", endian: binary.BigEndian, args: "golang", want: []byte{6, 'g', 'o', 'l', 'a', 'n', 'g'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteString(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteString() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteString() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteUShort(t *testing.T) {
	TestBOStream_WriteUint16(t)
}

func TestBOStream_WriteUint16(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      uint16
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteUShort_BigEndian_1", endian: binary.BigEndian, args: 1, want: []byte{0, 1}},
		{name: "TestBOStream_WriteUShort_LittleEndian_1", endian: binary.LittleEndian, args: 1, want: []byte{1, 0}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteUShort(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteUShort() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteUShort() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteUint32(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      uint32
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteUint32_BigEndian_1", endian: binary.BigEndian, args: 1, want: []byte{0, 0, 0, 1}},
		{name: "TestBOStream_WriteUint32_LittleEndian_1", endian: binary.LittleEndian, args: 1, want: []byte{1, 0, 0, 0}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteUint32(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteUint32() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteUint32() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteUint64(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      uint64
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_WriteUint64_BigEndian_1", endian: binary.BigEndian, args: 1, want: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "TestBOStream_WriteUint64_LittleEndian_1", endian: binary.LittleEndian, args: 1, want: []byte{1, 0, 0, 0, 0, 0, 0, 0}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteUint64(tt.args)
			if p.Error() != tt.wantError {
				t.Errorf("WriteUint64() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("WriteUint64() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}

func TestBOStream_WriteUint8(t *testing.T) {
	TestBOStream_WriteByte(t)
}

func TestBOStream_Write_Combined(t *testing.T) {
	type args struct {
		B   byte
		U16 uint16
		I32 int32
		F64 float64
		S   string
	}
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      *args
		want      []byte
		wantError error
	}{
		{name: "TestBOStream_Write_Combined_BigEndian_1", endian: binary.BigEndian, args: &args{B: 1, U16: 10, I32: 5, F64: math.SmallestNonzeroFloat64, S: "golang"}, want: []byte{1, 0, 10, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 1, 6, 103, 111, 108, 97, 110, 103}},
		{name: "TestBOStream_Write_Combined_LittleEndian_1", endian: binary.LittleEndian, args: &args{B: 1, U16: 10, I32: 5, F64: math.SmallestNonzeroFloat64, S: "golang"}, want: []byte{1, 10, 0, 5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 6, 103, 111, 108, 97, 110, 103}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			p := NewBOStream(tt.endian, buf)
			p.WriteByte(tt.args.B).WriteUint16(tt.args.U16).WriteInt32(tt.args.I32).WriteFloat64(tt.args.F64).WriteString(tt.args.S)

			if p.Error() != tt.wantError {
				t.Errorf("Write_Combined() has error")
			}
			if !reflect.DeepEqual(buf.Bytes(), tt.want) {
				t.Errorf("Write_Combined() = %v, want %v", buf.Bytes(), tt.want)
			}
		})
	}
}
