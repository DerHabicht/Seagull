package models

// BuildConfig defines the data needed to assemble the document.
type BuildConfig struct {
	// TopLevel defines how top level sections should be interpreted while processing.
	TopLevel Level `yaml:"top_level"`

	// TLP is the Traffic Light Protocol level that defines how to release this document.
	TLP TLP `yaml:"tlp"`
	Inputs []string `yaml:"inputs"`
	Annexes []string `yaml:"annexes"`
	VersionHistory []DocumentVersion `yaml:"version_history"`
}
