package models

import (
	"github.com/pkg/errors"
	"strings"
)

// Level is an enum that describes the sectioning types that can be used as a top-level section. This value is passed
// to Pandoc while converting Markdown documents.
// The valid values are:
//		- Chapter
//		- Section
type Level int

const (
	Chapter Level = iota + 1
	Section
)

// String returns the Level value as the string that Pandoc will be expecting.
func (l Level) String() string {
	switch l {
	case Chapter:
		return "chapter"
	case Section:
		return "section"
	default:
		return ""
	}
}

// Parse converts a string to the proper Level value. Strings are not case sensitive.
func (l *Level) Parse(s string) error {
	v := strings.ToLower(s)

	switch v {
	case "chapter":
		*l = Chapter
	case "section":
		*l = Section
	default:
		return errors.Errorf("%s is not a vaild Level value", s)
	}

	return nil
}

// MarshalYAML implements the Marshaler interface from go-yaml. This ensures that Level values are written in YAML
// documents correctly.
func (l Level) MarshalYAML() (string, error) {
	return l.String(), nil
}

// UnmarshalYAML implements the Unmarshaler interface from go-yaml. This ensures that Level values are read
// correctly from YAML documents.
func (l *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s); if err != nil {
		return err
	}

	return l.Parse(s)
}

