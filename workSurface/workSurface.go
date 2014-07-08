package workSurface

import "image/color"

type pixel struct {
	Color color.RGBA
	Used  bool
}

type Surface struct{
	pixels [][]pixel
	Size   int
}

func New(size int) Surface {
	topLevel := make([][] pixel, size)
	for i := range topLevel {
		topLevel[i] = make([]pixel, size)
	}
	return Surface{topLevel, size}
}

func (s *Surface) GetColor(x, y int) color.RGBA {
	if x >= s.Size || y >= s.Size {
		panic("GetColor :: index out of range")
	}
	return s.pixels[x][y].Color
}

func (s *Surface) SetColor(x, y int, r, g, b uint8) {
	if x >= s.Size || y >= s.Size {
		panic("SetColor :: index out of range")
	}
	s.pixels[x][y].Color = color.RGBA{r,g,b,255}

}

func (s *Surface) SetUsed(x, y int) {
	if x >= s.Size || y >= s.Size {
		panic("SetUsed :: index out of range")
	}
	s.pixels[x][y].Used = true
}

func (s *Surface) IsUsed(x, y int) bool {

	if x >= s.Size || y >= s.Size {
		panic("IsUsed :: index out of range")
	}
	return s.pixels[x][y].Used
}
