// Package superbase64 implements base64 encoding and decoding
// as specified in RFC 4648.
package superbase64

import (
	"errors"
	"strings"
)

type Encoding struct {
	table        string
	reverseTable [256]int
}

func NewStdEncoding() Encoding {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	return Encoding{table: table, reverseTable: buildReverseLookupTable(table)}
}

func NewURLEncoding() Encoding {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	return Encoding{table: table, reverseTable: buildReverseLookupTable(table)}
}

func buildReverseLookupTable(table string) [256]int {
	var reverseTable [256]int
	for i := range reverseTable {
		reverseTable[i] = -1
	}
	for i, c := range table {
		reverseTable[c] = i
	}
	return reverseTable
}

func (encoding Encoding) Encode(data []byte) string {
	var encoded strings.Builder
	padding := (3 - len(data)%3) % 3

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
			bits <<= (8 * (padding))
		}

		for j := 3; j >= 3-(len(byteSlice)); j-- {
			sixBitGroup := (bits & (0b111111 << (j * 6))) >> (j * 6)
			encoded.WriteByte(encoding.table[sixBitGroup])
		}
	}

	for i := 0; i != padding; i++ {
		encoded.WriteString("=")
	}

	return encoded.String()
}

func (encoding Encoding) Decode(str string) ([]byte, error) {
	if len(str)%4 != 0 {
		return nil, errors.New("invalid base64 string")
	}

	trimmedStr := strings.TrimRight(str, "=")
	padding := len(str) - len(trimmedStr)
	decoded := make([]byte, (len(str)/4*3)-padding)
	decodedIndex := 0

	for i := 0; i < len(trimmedStr); i += 4 {
		var substr string
		if i+4 > len(trimmedStr) {
			substr = trimmedStr[i:]
		} else {
			substr = trimmedStr[i : i+4]
		}

		var bits uint32
		for _, c := range substr {
			tableIndex := encoding.reverseTable[c]
			if tableIndex == -1 {
				return nil, errors.New("invalid base64 string")
			}

			bits = uint32(tableIndex) | bits<<6
		}
		if len(substr) != 4 {
			bits <<= (6 * (padding))
		}

		for j := 2; j >= 4-(len(substr)); j-- {
			eightBitGroup := (bits & (0b11111111 << (j * 8))) >> (j * 8)
			decoded[decodedIndex] = byte(eightBitGroup)
			decodedIndex++
		}
	}

	return decoded, nil
}
