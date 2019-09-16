package figma

import (
	"encoding/json"
)

// EntryColor represents a color entry
type EntryColor struct {
	// R represents the r value of an RGBA color scheme
	R float64
	// G represents the g value of an RGBA color scheme
	G float64
	// B represents the b value of an RGBA color scheme
	B float64
	// A represents the a value of an RGBA color scheme
	A float64
	// Hex represents the hex value of an RGBA color scheme
	//  NOTE: this is not returned by Figma API, rather calculated by lib and added to entry.
	Hex string
}

// UnmarshalJSON is a helper method for the json.Decoding process
func (e *EntryColor) UnmarshalJSON(bs []byte) (err error) {
	var m map[string]interface{}

	if err = json.Unmarshal(bs, &m); err != nil {
		return
	}

	var ec EntryColor
	ec.R = m["r"].(float64)
	ec.G = m["g"].(float64)
	ec.B = m["b"].(float64)
	ec.A = m["a"].(float64)

	if ec.Hex, err = ec.getHex(); err != nil {
		return
	}

	*e = ec
	return
}

func (e *EntryColor) getHex() (hex string, err error) {
	r := uint8(e.R * 255)
	g := uint8(e.G * 255)
	b := uint8(e.B * 255)
	h := newHex(r, g, b)

	if err = h.Validate(); err != nil {
		return
	}

	hex = h.Hex
	return
}
