package colorCube

import (
	"image/color"
	"math"
)

const MAX_BIT_SIZE = 8
const MAX_SIDE_SIZE = 1 << MAX_BIT_SIZE


type ColorCube struct {
	SideSize int
	BitSize uint8
	Cube [][][]bool
}

func New(bitSize uint8) (*ColorCube ) {
	if bitSize > MAX_BIT_SIZE {
		panic("Bit size to big")
	}
	sideSize := 1 << bitSize
	topLevel := make([][][]bool, sideSize)
	for i := range topLevel {
		middleLevel := make([][]bool, sideSize)
		for j := range middleLevel {
			middleLevel[j] = make([]bool, sideSize)
		}
		topLevel[i] = middleLevel
	}
	c := ColorCube{sideSize, bitSize, topLevel}
	return &c
}

func (c ColorCube) IsUsed(r, g, b int) (bool) {
	return c.Cube[r][g][b]
}

func (c ColorCube) SetUsed(r, g, b int) {
	c.Cube[r][g][b] = true
}

func (c ColorCube) GetColor(x, y, z int) (color.RGBA ) {
	if x > c.SideSize || y >= c.SideSize || z >= c.SideSize {
		panic("index out of range")
	}

	ratio := float64(MAX_SIDE_SIZE-1) / float64(c.SideSize-1)
	xIndex := uint8(math.Trunc(float64(x) * ratio))
	yIndex := uint8(math.Trunc(float64(y) * ratio))
	zIndex := uint8(math.Trunc(float64(z) * ratio))
	return color.RGBA{xIndex, yIndex, zIndex, 255}
}
