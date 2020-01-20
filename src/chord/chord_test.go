package chord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChord(t *testing.T) {
	c, err := Parse("bviis4")
	assert.Nil(t, err)
	assert.Equal(t, c.Name, "Bbms4")
	c, err = Parse("i")
	assert.Nil(t, err)
	assert.Equal(t, c.Name, "Cm")
	c, err = Parse("I6")
	assert.Nil(t, err)
	assert.Equal(t, c.Name, "C6")
	c, err = Parse("bVII")
	assert.Nil(t, err)
	assert.Equal(t, c.Name, "Bb")
}
