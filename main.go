package main

import (
	"fmt"

	"github.com/is386/super-base-64/superbase64"
)

func main() {
	input := "Heck"
	fmt.Printf("Text:    %s\n", input)

	bytesData := []byte(input)
	fmt.Printf("Binary:  %v\n", bytesData)

	encoded := superbase64.NewStdEncoding().Encode(bytesData)
	fmt.Println("Encoded:", encoded)

	decoded, err := superbase64.NewStdEncoding().Decode(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Decoded: %v\n", decoded)
}
