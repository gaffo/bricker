package bricker_test

import (
	"github.com/gaffo/bricker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ParseLXF(t *testing.T) {
	parser := bricker.LFXParser{}
	lxfml, err := parser.Parse("simple.lxf")
	assert.NoError(t, err)
	assert.NotNil(t, lxfml)
	assert.Equal(t, 3, len(lxfml.Bricks))

	assert.Equal(t, 1, len(lxfml.Bricks[0].Parts))
	assert.Equal(t, "90617", lxfml.Bricks[0].Parts[0].DesignID)
	assert.Equal(t, "191", lxfml.Bricks[0].Parts[0].Materials)

	assert.Equal(t, 1, len(lxfml.Bricks[1].Parts))
	assert.Equal(t, "90608", lxfml.Bricks[1].Parts[0].DesignID)
	assert.Equal(t, "21", lxfml.Bricks[1].Parts[0].Materials)

	assert.Equal(t, 1, len(lxfml.Bricks[2].Parts))
	assert.Equal(t, "90609", lxfml.Bricks[2].Parts[0].DesignID)
	assert.Equal(t, "26", lxfml.Bricks[2].Parts[0].Materials)
}
