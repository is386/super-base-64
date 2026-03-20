package main

import (
	"fmt"
	"strings"
)

const table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	input := "And"
	bytesData := []byte(input)

	fmt.Println("Original:", input)

	var encoded strings.Builder
	var missing int
	for i := 0; i < len(bytesData); i += 3 {
		missing = 0

		var byteSlice []byte
		if i+3 > len(bytesData) {
			missing = i + 3 - len(bytesData)
			byteSlice = bytesData[i:]
		} else {
			byteSlice = bytesData[i : i+3]
		}

		var bits uint32
		for _, b := range byteSlice {
			bits = uint32(b) | bits<<8
		}
		if missing != 0 {
			bits <<= (8 * (missing))
		}

		// TODO: simplify the missing logic by maybe taking the length of the group and if its less than 6 then shift it over to the left X spaces
		for j := 3; j >= missing; j-- {
			sixBitGroup := (bits & (0b111111 << (j * 6))) >> (j * 6)
			fmt.Fprintf(&encoded, "%c", table[sixBitGroup])
		}
	}

	for i := 0; i != missing; i++ {
		fmt.Fprintf(&encoded, "%s", "=")
	}

	fmt.Println("Encoded: ", encoded.String())
}
