package main

import (
	"fmt"

	"github.com/is386/super-base-64/base64"
)

func main() {
	input := "Heck"
	bytesData := []byte(input)
	encoded := base64.NewStdEncoding().Encode(bytesData)
	fmt.Println(encoded)
}
