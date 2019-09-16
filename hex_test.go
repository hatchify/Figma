package figma

import (
	"errors"
	"testing"
)

func TestHex_new(t *testing.T) {
	r, g, b := uint8(1), uint8(1), uint8(1)
	hex := newHex(r, g, b)

	if hex.Hex == "" {
		err := errors.New("hex was not set")
		t.Fatal(err)
	}
}

func TestHex_Validate(t *testing.T) {
	r, g, b := uint8(1), uint8(1), uint8(1)
	hex := newHex(r, g, b)

	if hex.Hex == "" {
		err := errors.New("hex was not set")
		t.Fatal(err)
	}

	if err := hex.Validate(); err != nil {
		t.Fatal(err)
	}
}
