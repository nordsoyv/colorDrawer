package workSurface

import (
	"image/color"
	"os"
	"bufio"
	"image/png"
	"image"
	"fmt"
)

type pixel struct {
	Color color.RGBA
	Used  bool
}

type Surface struct{
	pixels [][]pixel
	Size   int
}

func New(sideSize int) Surface {
	topLevel := make([][] pixel, sideSize)
	for i := range topLevel {
		topLevel[i] = make([]pixel, sideSize)
	}
	return Surface{topLevel, sideSize}
}

func (s *Surface) GetColor(x, y int) color.RGBA {
	if x >= s.Size || y >= s.Size {
		panic("GetColor :: index out of range")
	}
	return s.pixels[x][y].Color
}

func (s *Surface) SetColorRGB(x, y int, r, g, b uint8) {
	if x >= s.Size || y >= s.Size {
		panic(fmt.Sprintf("SetColor :: index out of range, was %v or %v, should be max %v", x, y, s.Size))
	}
	s.pixels[x][y].Color = color.RGBA{r,g,b,255}
}

func (s *Surface) SetColor(x, y int, c color.RGBA) {
	if x >= s.Size || y >= s.Size {
		panic(fmt.Sprintf("SetColor :: index out of range, was %v or %v, should be max %v", x, y, s.Size))
	}
	s.pixels[x][y].Color = color.RGBA{c.R, c.G, c.B, 255}
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

func (s *Surface) ToPng(fileName string) {
	fmt.Println("Writing to file : ", fileName)
	outfile , err := os.Create(fileName)
	check(err)
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)
	defer writer.Flush()
	err = png.Encode(writer, s.toImage())
	check(err)
}

func (s *Surface) toImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, s.Size, s.Size))
	for x := 0; x < s.Size; x++ {
		for y := 0 ; y < s.Size; y++ {
			img.Set(x, y, s.GetColor(x, y))
		}
	}
	return img
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
