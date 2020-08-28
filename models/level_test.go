package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevel_Parse_Uppercase(t *testing.T) {
	var chapter Level
	assert.NoError(t, chapter.Parse("CHAPTER"), "Error while parsing 'CHAPTER'")
	assert.Equal(t, Chapter, chapter)

	var section Level
	assert.NoError(t, section.Parse("SECTION"), "Error while parsing 'SECTION'")
	assert.Equal(t, Section, section)
}

func TestLevel_Parse_Lowercase(t *testing.T) {
	var chapter Level
	assert.NoError(t, chapter.Parse("chapter"), "Error while parsing 'chapter'")
	assert.Equal(t, Chapter, chapter)

	var section Level
	assert.NoError(t, section.Parse("section"), "Error while parsing 'section'")
	assert.Equal(t, Section, section)
}

func TestLevel_Parse_InvalidLevel(t *testing.T) {
	var invalid Level
	assert.Error(t, invalid.Parse("invalid"), "Error not thrown when parsing invalid Level value")
}
