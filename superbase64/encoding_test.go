package superbase64 

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
		{"5 bytes", "Hello"},
		{"6 bytes", "Helloo"},
		{"long string", "Many hands make light work."},
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
	
	t.Run("Encode: URL", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("when input is %s", tt.name), func(t *testing.T) {
				t.Run("it then encodes the given byte array", func(t *testing.T) {
					input := []byte(tt.input)
					expected := base64.URLEncoding.EncodeToString(input)
					result := NewURLEncoding().Encode(input)
					if result != expected {
						t.Errorf(`expected: %s, actual: %s`, expected, result)
					}
				})
			})
		}
	})

	t.Run("Decode: Standard", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("when input is %s", tt.name), func(t *testing.T) {
				t.Run("it then decodes the given base64 string", func(t *testing.T) {
					input := base64.StdEncoding.EncodeToString([]byte(tt.input))
					expected := tt.input
					result, err := NewStdEncoding().Decode(input)
					if err != nil {
						t.Fatalf(`unexpected error: %v`, err)
					}
					if string(result) != expected {
						t.Errorf(`expected: %s, actual: %s`, expected, string(result))
					}
				})
			})
		}
	})

	t.Run("Decode: URL", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("when input is %s", tt.name), func(t *testing.T) {
				t.Run("it then decodes the given base64 string", func(t *testing.T) {
					input := base64.URLEncoding.EncodeToString([]byte(tt.input))
					expected := tt.input
					result, err := NewURLEncoding().Decode(input)
					if err != nil {
						t.Fatalf(`unexpected error: %v`, err)
					}
					if string(result) != expected {
						t.Errorf(`expected: %s, actual: %s`, expected, string(result))
					}
				})
			})
		}
	})

	t.Run("Decode: invalid input", func(t *testing.T) {
		invalidTests := []struct {
			name  string
			input string
		}{
			{"invalid characters", "!!!!"},
			{"invalid length 1", "A"},
			{"invalid length 2", "AB"},
			{"invalid length 3", "ABC"},
			{"invalid length 5", "ABCDA"},
		}

		for _, tt := range invalidTests {
			t.Run(fmt.Sprintf("when input is %s", tt.name), func(t *testing.T) {
				_, err := NewStdEncoding().Decode(tt.input)
				if err == nil {
					t.Errorf("expected error for invalid base64 string: %s", tt.input)
				}
			})
		}
	})

}
