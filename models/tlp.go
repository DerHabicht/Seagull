package models

import (
	"github.com/pkg/errors"
	"strings"
)

// TLP is an enum that describes the levels of the Traffic Light Protocol.
// The valid values are:
//		- Red
//		- Amber
//		- Green
//		- White
type TLP int

const (
	Red TLP = iota + 1
	Amber
	Green
	White
)

// String returns the correct symbol for the TLP level.
func (t TLP) String() string {
	switch t {
	case Red:
		return "TLP:RED"
	case Amber:
		return "TLP:AMBER"
	case Green:
		return "TLP:GREEN"
	case White:
		return "TLP:WHITE"
	default:
		return ""
	}
}

// AsKey returns the TLP level as the string key that would be needed in using the tlp LaTeX package.
func (t TLP) AsKey() string {
	switch t {
	case Red:
		return "red"
	case Amber:
		return "amber"
	case Green:
		return "green"
	case White:
		return "white"
	default:
		return ""
	}
}

// Parse turns a string into its appropriate TLP value. Strings are not case sensitive and they may optionally
// include the TLP: prefix.
func (t *TLP) Parse(s string) error {
	var v string
	if len(s) > 4 && strings.ToLower(s)[:4] == "tlp:" {
		v = strings.ToLower(s)[4:]
	} else {
		v = strings.ToLower(s)
	}

	switch v {
	case "red":
		*t = Red
	case "amber":
		*t = Amber
	case "green":
		*t = Green
	case "white":
		*t = White
	default:
		return errors.Errorf("%s is not a valid TLP level", s)
	}

	return nil
}

// MarshalYAML implements the Marshaler interface from go-yaml. This ensures that TLP values are written into YAML
// documents correctly. Since YAML documents used by Seagull will want the key version of the TLP, AsKey() is invoked
// instead of String().
func (t TLP) MarshalYAML() (string, error) {
	return t.AsKey(), nil
}

// UnmarshalYAML implements the Unmarshaler interface from go-yaml. This ensures that TLP values are interpreted
// correctly from YAML documents.
func (t *TLP) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(s); if err != nil {
		return err
	}

	return t.Parse(s)
}
