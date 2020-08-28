package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTLPParse_UpperCase(t *testing.T) {
	var red TLP
	assert.NoError(t, red.Parse("RED"), "Error while parsing 'RED'")
	assert.Equal(t, Red, red)

	var amber TLP
	assert.NoError(t, amber.Parse("AMBER"), "Error while parsing 'AMBER'")
	assert.Equal(t, Amber, amber)

	var green TLP
	assert.NoError(t, green.Parse("GREEN"), "Error while parsing 'GREEN'")
	assert.Equal(t, Green, green)

	var white TLP
	assert.NoError(t, white.Parse("WHITE"), "Error while parsing 'WHITE'")
	assert.Equal(t, White, white)
}

func TestTLPParse_LowerCase(t *testing.T) {
	var red TLP
	assert.NoError(t, red.Parse("red"), "Error while parsing 'red'")
	assert.Equal(t, Red, red)

	var amber TLP
	assert.NoError(t, amber.Parse("amber"), "Error while parsing 'amber'")
	assert.Equal(t, Amber, amber)

	var green TLP
	assert.NoError(t, green.Parse("green"), "Error while parsing 'green'")
	assert.Equal(t, Green, green)

	var white TLP
	assert.NoError(t, white.Parse("white"), "Error while parsing 'white'")
	assert.Equal(t, White, white)

}

func TestTLPParse_WithColon(t *testing.T) {
	var red TLP
	assert.NoError(t, red.Parse("TLP:RED"), "Error while parsing 'TLP:RED'")
	assert.Equal(t, Red, red)

	var amber TLP
	assert.NoError(t, amber.Parse("TLP:AMBER"), "Error while parsing 'TLP:AMBER'")
	assert.Equal(t, Amber, amber)

	var green TLP
	assert.NoError(t, green.Parse("TLP:GREEN"), "Error while parsing 'TLP:GREEN'")
	assert.Equal(t, Green, green)

	var white TLP
	assert.NoError(t, white.Parse("TLP:WHITE"), "Error while parsing 'TLP:WHITE'")
	assert.Equal(t, White, white)
}

func TestTLP_Parse_InvalidTLP(t *testing.T) {
	var invalid TLP
	assert.Error(t, invalid.Parse("invalid"), "Error not thrown while parsing invalid TLP value")
}
