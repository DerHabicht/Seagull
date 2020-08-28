package models

// LaTeXDocument is a struct that governs the creation of LaTeX Documents from build settings.
type LaTeXDocument struct {
	Slug string
	Project Project
	BuildConfig BuildConfig
}