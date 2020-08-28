package models

import "time"

// DocumentVersion defines data that is entered into the document's version history table.
type DocumentVersion struct {
	// Version is the user-designated version string to be included in the version history table.
	Version string `yaml:"version"`

	// Date is the date of this version, as defined by the user.
	Date time.Time `yaml:"date"`

	// Author is typically the initials of the author writing or editing this version of the document. At the author's
	// option, however, full names may be included here.
	Author string `yaml:"author"`

	// Remarks is the space provided to describe which changes were made in this version.
	Remarks string `yaml:"author"`
}
