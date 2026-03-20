package base64

import (
	"strings"
)

type Encoding struct {
	table string
}

func NewStdEncoding() Encoding {
	return Encoding{table: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"}
}

func (encoding Encoding) Encode(data []byte) string {
	var encoded strings.Builder
	numMissingBytes := (3 - len(data)%3) % 3

	for i := 0; i < len(data); i += 3 {
		var byteSlice []byte
		if i+3 > len(data) {
			byteSlice = data[i:]
		} else {
			byteSlice = data[i : i+3]
		}

		var bits uint32
		for _, b := range byteSlice {
			bits = uint32(b) | bits<<8
		}
		if len(byteSlice) != 3 {
			bits <<= (8 * (numMissingBytes))
		}

		for j := 3; j >= 3-(len(byteSlice)); j-- {
			sixBitGroup := (bits & (0b111111 << (j * 6))) >> (j * 6)
			encoded.WriteByte(encoding.table[sixBitGroup])
		}
	}

	for i := 0; i != numMissingBytes; i++ {
		encoded.WriteString("=")
	}

	return encoded.String()
}
