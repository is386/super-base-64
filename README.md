# super-base-64

A Base64 encoding and decoding library for Go, implementing RFC 4648.

## Features

- Standard base64 encoding/decoding (`+/` alphabet)
- URL-safe base64 encoding/decoding (`-_` alphabet)

## Installation

```bash
go get github.com/is386/super-base-64
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/is386/super-base-64/superbase64"
)

func main() {
	// Encode
	encoded := superbase64.NewStdEncoding().Encode([]byte("Hello, World!"))
	fmt.Println(encoded) // SGVsbG8sIFdvcmxkIQ==

	// Decode
	decoded, err := superbase64.NewStdEncoding().Decode(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decoded)) // Hello, World!

	// URL-safe encoding
	urlEncoded := superbase64.NewURLEncoding().Encode([]byte("Hello, World!"))
	fmt.Println(urlEncoded) // SGVsbG8sIFdvcmxkIQ==
}
```
