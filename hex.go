package figma

import (
	"fmt"
	"regexp"

	"github.com/Hatch1fy/errors"
)

const (
	// ErrEmptyHexValue will return when provided API key is empty
	ErrEmptyHexValue = errors.Error("provided hex value for HexBrandColor is empty")
)

func newHex(r, g, b uint8) (h Hex) {
	h.Hex = fmt.Sprintf("%02x%02x%02x", r, g, b)
	return
}

// Hex is the Hex entry
type Hex struct {
	// Hex represents the hex representation of a brand color
	Hex string
}

// Validate will validate a HexBrandColor entry
func (h *Hex) Validate() (err error) {
	if h.Hex == "" {
		err = ErrEmptyHexValue
		return
	}

	// Hex values should be six characters in length
	if len(h.Hex) != 6 {
		err = fmt.Errorf("invalid length for hex value: expected 6, got %d", len(h.Hex))
		return
	}

	// Hex values should only be alphanumeric
	var isAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	if !isAlphaNumeric(h.Hex) {
		err = fmt.Errorf("invalid charater: expected only alphanumeric characters, received %s", h.Hex)
		return
	}

	return
}
