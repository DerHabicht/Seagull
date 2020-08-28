package models

// Project defines the attributes of a specific writing project.
type Project struct {
	// Name is the name of the project.
	Name string `yaml:"name"`

	// LatexPath is the directory where the LaTeX template and other files needed to build the LaTeX document are
	// located.
	LatexPath string `yaml:"latex_path"`

	// ContentPath is the directory where the Markdown files are located.
	ContentPath string `yaml:"content_path"`

	// Compendium is a flag that indicates whether this project should be included in the workspace's compendium
	// project.
	Compendium bool `yaml:"compendium"`
}
