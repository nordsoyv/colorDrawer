package colorCube

import (
	"image/color"
	"math"
)

const MAX_SIZE = 256

type ColorCube struct {
	Size int
	Cube [][][]bool
}

func New(size int) (*ColorCube ) {
	if size > MAX_SIZE {
		panic("Size to big")
	}
	topLevel := make([][][]bool, size)
	for i := range topLevel {
		middleLevel := make([][]bool, size)
		for j := range middleLevel {
			middleLevel[j] = make([]bool, size)
		}
		topLevel[i] = middleLevel
	}
	c := ColorCube{size, topLevel}
	return &c
}

func (c ColorCube) IsUsed(r, g, b int) (bool) {
	return c.Cube[r][g][b]
}

func (c ColorCube) SetUsed(r, g, b int) {
	c.Cube[r][g][b] = true
}

func (c ColorCube) GetColor(x, y, z int) (color.RGBA ) {
	if x > MAX_SIZE || y >= MAX_SIZE || z >= MAX_SIZE {
		panic("index out of range")
	}
	if x > c.Size || y >= c.Size || z >= c.Size {
		panic("index out of range")
	}

	ratio := float64(MAX_SIZE-1) / float64(c.Size-1)
	xIndex := uint8(math.Trunc(float64(x) * ratio))
	yIndex := uint8(math.Trunc(float64(y) * ratio))
	zIndex := uint8(math.Trunc(float64(z) * ratio))
	return color.RGBA{xIndex, yIndex, zIndex, 255}
}
