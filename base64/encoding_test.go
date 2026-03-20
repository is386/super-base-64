package base64

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"1 byte", "M"},
		{"2 bytes", "Ma"},
		{"3 bytes", "Man"},
		{"multi block", "Many"},
		{"all zero bytes", "\x00\x00\x00"},
		{"all bits set", "\xff\xff\xff"},
		{"full alphabet range", "\xfb\xef\xbe"},
	}

	t.Run("Encode: Standard", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("when input is %s", tt.name), func(t *testing.T) {
				t.Run("it then encodes the given byte array", func(t *testing.T) {
					input := []byte(tt.input)
					expected := base64.StdEncoding.EncodeToString(input)
					result := NewStdEncoding().Encode(input)
					if result != expected {
						t.Errorf(`expected: %s, actual: %s`, expected, result)
					}
				})
			})
		}
	})
}
