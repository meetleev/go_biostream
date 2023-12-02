package bitstream

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
	"testing"
)

func TestBIStream_ReadBool(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      bool
		wantError error
	}{
		{name: "TestBIStream_ReadBool", endian: binary.LittleEndian, want: true, args: []byte{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			b, err := p.ReadBool()
			if err != tt.wantError {
				t.Errorf("ReadBool() has error")
			}
			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("ReadBool() = %v, want %v", b, tt.want)
			}
		})
	}
}

func TestBIStream_ReadByte(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      byte
		wantError error
	}{
		{name: "TestBIStream_ReadByte", endian: binary.LittleEndian, args: []byte{1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			b, err := p.ReadByte()
			if err != tt.wantError {
				t.Errorf("ReadByte() has error")
			}
			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("ReadByte() = %v, want %v", b, tt.want)
			}
		})
	}
}

func TestBIStream_ReadBytes(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      []byte
		wantError error
	}{
		{name: "TestBIStream_ReadBytes", endian: binary.LittleEndian, args: []byte{1, 3}, want: []byte{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadBytes(2)
			if err != tt.wantError {
				t.Errorf("ReadBytes() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadBytes() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadBytesWithLengthPrefix(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      []byte
		wantError error
	}{
		{name: "TestBIStream_ReadBytesWithLengthPrefix", endian: binary.LittleEndian, want: []byte{1, 3}, args: []byte{2, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadBytesWithLengthPrefix()
			if err != tt.wantError {
				t.Errorf("ReadBytesWithLengthPrefix() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadBytesWithLengthPrefix() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadFloat32(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      float32
		wantError error
	}{
		{name: "TestBIStream_ReadFloat32_BigEndian", endian: binary.BigEndian, want: math.SmallestNonzeroFloat32, args: []byte{0, 0, 0, 1}},
		{name: "TestBIStream_ReadFloat32_LittleEndian", endian: binary.LittleEndian, want: math.SmallestNonzeroFloat32, args: []byte{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadFloat32()
			if err != tt.wantError {
				t.Errorf("ReadFloat32() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadFloat32() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadFloat64(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      float64
		wantError error
	}{
		{name: "TestBIStream_ReadFloat64_BigEndian", endian: binary.BigEndian, want: math.SmallestNonzeroFloat64, args: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "TestBIStream_ReadFloat64_LittleEndian", endian: binary.LittleEndian, want: math.SmallestNonzeroFloat64, args: []byte{1, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadFloat64()
			if err != tt.wantError {
				t.Errorf("ReadFloat64() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadFloat64() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadInt16(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      int16
		wantError error
	}{
		{name: "TestBIStream_ReadInt16_BigEndian_1", endian: binary.BigEndian, want: 1, args: []byte{0, 1}},
		{name: "TestBIStream_ReadInt16_LittleEndian_1", endian: binary.LittleEndian, want: 1, args: []byte{1, 0}},
		{name: "TestBIStream_ReadInt16_BigEndian_10-", endian: binary.BigEndian, want: -10, args: []byte{255, 246}},
		{name: "TestBIStream_ReadInt16_LittleEndian_10-", endian: binary.LittleEndian, want: -10, args: []byte{246, 255}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadInt16()
			if err != tt.wantError {
				t.Errorf("ReadInt16() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadInt16() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadInt32(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      int32
		wantError error
	}{
		{name: "TestBIStream_ReadInt32_BigEndian_1", endian: binary.BigEndian, want: 1, args: []byte{0, 0, 0, 1}},
		{name: "TestBIStream_ReadInt32_LittleEndian_1", endian: binary.LittleEndian, want: 1, args: []byte{1, 0, 0, 0}},
		{name: "TestBIStream_ReadInt32_BigEndian_10-", endian: binary.BigEndian, want: -10, args: []byte{255, 255, 255, 246}},
		{name: "TestBIStream_ReadInt32_LittleEndian_10-", endian: binary.LittleEndian, want: -10, args: []byte{246, 255, 255, 255}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadInt32()
			if err != tt.wantError {
				t.Errorf("ReadInt32() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadInt32() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadInt64(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      int64
		wantError error
	}{
		{name: "TestBIStream_ReadInt64_BigEndian_1", endian: binary.BigEndian, want: 1, args: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "TestBIStream_ReadInt64_LittleEndian_1", endian: binary.LittleEndian, want: 1, args: []byte{1, 0, 0, 0, 0, 0, 0, 0}},
		{name: "TestBIStream_ReadInt64_BigEndian_10-", endian: binary.BigEndian, want: -10, args: []byte{255, 255, 255, 255, 255, 255, 255, 246}},
		{name: "TestBIStream_ReadInt64_LittleEndian_10-", endian: binary.LittleEndian, want: -10, args: []byte{246, 255, 255, 255, 255, 255, 255, 255}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadInt64()
			if err != tt.wantError {
				t.Errorf("ReadInt64() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadInt64() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadInt8(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      int8
		wantError error
	}{
		{name: "TestBIStream_ReadInt8", endian: binary.LittleEndian, want: 1, args: []byte{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadInt8()
			if err != tt.wantError {
				t.Errorf("ReadInt8() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadInt8() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadShort(t *testing.T) {
	TestBIStream_ReadInt16(t)
}

func TestBIStream_ReadString(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      string
		wantError error
	}{
		{name: "TestBIStream_ReadString_golang", endian: binary.BigEndian, want: "golang", args: []byte{6, 'g', 'o', 'l', 'a', 'n', 'g'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadString()
			if err != tt.wantError {
				t.Errorf("ReadString() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadString() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadUShort(t *testing.T) {
	TestBIStream_ReadUint16(t)
}

func TestBIStream_ReadUint16(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      uint16
		wantError error
	}{
		{name: "TestBIStream_ReadUint16_BigEndian_1", endian: binary.BigEndian, want: 1, args: []byte{0, 1}},
		{name: "TestBIStream_ReadUint16_LittleEndian_1", endian: binary.LittleEndian, want: 1, args: []byte{1, 0}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadUint16()
			if err != tt.wantError {
				t.Errorf("ReadUint16() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadUint16() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadUint32(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      uint32
		wantError error
	}{
		{name: "TestBIStream_ReadUint32_BigEndian_1", endian: binary.BigEndian, want: 1, args: []byte{0, 0, 0, 1}},
		{name: "TestBIStream_ReadUint32_LittleEndian_1", endian: binary.LittleEndian, want: 1, args: []byte{1, 0, 0, 0}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadUint32()
			if err != tt.wantError {
				t.Errorf("ReadUint32() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadUint32() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadUint64(t *testing.T) {
	tests := []struct {
		name      string
		endian    binary.ByteOrder
		args      []byte
		want      uint64
		wantError error
	}{
		{name: "TestBIStream_ReadUint64_BigEndian_1", endian: binary.BigEndian, want: 1, args: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{name: "TestBIStream_ReadUint64_LittleEndian_1", endian: binary.LittleEndian, want: 1, args: []byte{1, 0, 0, 0, 0, 0, 0, 0}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			data, err := p.ReadUint64()
			if err != tt.wantError {
				t.Errorf("ReadUint64() has error")
			}
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ReadUint64() = %v, want %v", data, tt.want)
			}
		})
	}
}

func TestBIStream_ReadUint8(t *testing.T) {
	TestBIStream_ReadByte(t)
}

func TestBIStream_Read_Combined(t *testing.T) {
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
		args      []byte
		want      *args
		wantError error
	}{
		{name: "TestBIStream_Read_Combined_BigEndian_1", endian: binary.BigEndian, want: &args{B: 1, U16: 10, I32: 5, F64: math.SmallestNonzeroFloat64, S: "golang"}, args: []byte{1, 0, 10, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 1, 6, 103, 111, 108, 97, 110, 103}},
		{name: "TestBIStream_Read_Combined_LittleEndian_1", endian: binary.LittleEndian, want: &args{B: 1, U16: 10, I32: 5, F64: math.SmallestNonzeroFloat64, S: "golang"}, args: []byte{1, 10, 0, 5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 6, 103, 111, 108, 97, 110, 103}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			buf.Write(tt.args)
			p := NewBIStream(tt.endian, buf)
			result := &args{}
			var err error = nil
			for nil == err {
				result.B, err = p.ReadByte()
				result.U16, err = p.ReadUint16()
				result.I32, err = p.ReadInt32()
				result.F64, err = p.ReadFloat64()
				result.S, err = p.ReadString()
				break
			}
			if err != nil {
				t.Errorf("Read_Combined() has error")
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Read_Combined() = %v, want %v", result, tt.want)
			}
		})
	}
}
