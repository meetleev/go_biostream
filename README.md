# go_bitstream [![GoDoc][1]][2] [![Go Report Card][3]][4] [![MIT licensed][5]][6]

---
[1]: https://godoc.org/github.com/meetleev/go_bitstream?status.svg
[2]: https://godoc.org/github.com/meetleev/go_bitstream
[3]: https://goreportcard.com/badge/github.com/meetleev/go_bitstream
[4]: https://goreportcard.com/report/github.com/meetleev/go_bitstream
[5]: https://img.shields.io/badge/license-Apache-blue.svg
[6]: LICENSE
[7]: https://github.com/meetleev/go_bitstream/actions/workflows/tests.yaml/badge.svg
[8]: https://github.com/meetleev/go_bitstream/actions/workflows/tests.yaml

A binary write and read.

# Install

```bash
go get github.com/meetleev/go_bitstream
```

# Usage

* Write

``` go
import bitstream "github.com/meetleev/go_bitstream"

buffer := new(bytes.Buffer)
writer := bitstream.NewBOStream(binary.BigEndian, buffer)
writer.WriteByte(1).WriteString("golang").WriteUint16(math.MaxUint16)
err := writer.Error()
```

* Read

``` go
import bitstream "github.com/meetleev/go_bitstream"

buffer := new(bytes.Buffer)
writer := bitstream.NewBOStream(binary.BigEndian, buffer)
writer.WriteByte(1).WriteString("golang").WriteUint16(math.MaxUint16)
err := writer.Error()

val1, err := unpacker.ReadByte()
val2, err := unpacker.ReadString()
val3, err := unpacker.ReadUint16()
```