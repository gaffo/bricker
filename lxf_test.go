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

func Test_ConvertSimple(t *testing.T) {
	parser := bricker.LFXParser{}
	lxfml, err := parser.Parse("simple.lxf")
	assert.NoError(t, err)

	colorMap, err := bricker.ParseColorMap("ldd_bl_colors.csv")
	assert.NoError(t, err)
	partsMap, err := bricker.ParsePartsMap("ldd_bl_parts.csv")
	assert.NoError(t, err)

	blParts := lxfml.ConvertWithSources(colorMap, partsMap)

	assert.Equal(t, 3, len(blParts))
	assert.Equal(t,
		bricker.BLPart{
			ItemID:   "99475",
			Color:    "110",
			Quantity: 1,
		},
		blParts[0])
	assert.Equal(t,
		bricker.BLPart{
			ItemID:   "99476",
			Color:    "5",
			Quantity: 1,
		},
		blParts[1])
	assert.Equal(t,
		bricker.BLPart{
			ItemID:   "99495",
			Color:    "11",
			Quantity: 1,
		},
		blParts[2])
}
