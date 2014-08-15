package workSurface

import (
	"bufio"
	"container/list"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
)

type pixel struct {
	Color color.RGBA
	Used  bool
}

type Coord2D struct {
	X, Y int
}

type Surface struct {
	pixels [][]pixel
	Size   int
	lock   *sync.RWMutex
}

func New(sideSize int) Surface {
	topLevel := make([][]pixel, sideSize)
	for i := range topLevel {
		topLevel[i] = make([]pixel, sideSize)
	}
	var lock sync.RWMutex
	return Surface{topLevel, sideSize, &lock}
}

func (s *Surface) GetColor(x, y int) color.RGBA {
	if x >= s.Size || y >= s.Size {
		panic("GetColor :: index out of range")
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.pixels[x][y].Color
}

func (s *Surface) SetColorRGB(x, y int, r, g, b uint8) {
	if x >= s.Size || y >= s.Size {
		panic(fmt.Sprintf("SetColor :: index out of range, was %v or %v, should be max %v", x, y, s.Size))
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.pixels[x][y].Color = color.RGBA{r, g, b, 255}
}

func (s *Surface) SetColor(x, y int, c color.RGBA) {
	if x >= s.Size || y >= s.Size {
		panic(fmt.Sprintf("SetColor :: index out of range, was %v or %v, should be max %v", x, y, s.Size))
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.pixels[x][y].Color = color.RGBA{c.R, c.G, c.B, 255}
	s.pixels[x][y].Used = true
}

func (s *Surface) SetUsed(x, y int) {
	if x >= s.Size || y >= s.Size {
		panic("SetUsed :: index out of range")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.pixels[x][y].Used = true
}

func (s *Surface) SetNotUsed(x, y int) {
	if x >= s.Size || y >= s.Size {
		panic("SetNotUsed :: index out of range")
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.pixels[x][y].Used = false
}

func (s *Surface) IsUsed(x, y int) bool {
	if x >= s.Size || y >= s.Size {
		panic(fmt.Sprintf("IsUsed :: index out of range (%v,%v)", x, y))
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.pixels[x][y].Used
}

func (s *Surface) ToPng(fileName string) {
	fmt.Println("Writing to file : ", fileName)
	outfile, err := os.Create(fileName)
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
		for y := 0; y < s.Size; y++ {
			img.Set(x, y, s.GetColor(x, y))
		}
	}
	return img
}

func (s *Surface) FindNeighborPixels(p Coord2D) (used, unUsed *list.List) {
	unUsed = list.New()
	used = list.New()
	if p.X > 0 {
		s.filterPixel(leftPixel(p), used, unUsed)
	}
	if p.X < s.Size-1 {
		s.filterPixel(rightPixel(p), used, unUsed)
	}
	if p.Y < s.Size-1 {
		s.filterPixel(upPixel(p), used, unUsed)
	}
	if p.Y > 0 {
		s.filterPixel(downPixel(p), used, unUsed)
	}
	if p.Y < s.Size-1 && p.X > 0 {
		s.filterPixel(upLeftPixel(p), used, unUsed)
	}
	if p.Y < s.Size-1 && p.X < s.Size-1 {
		s.filterPixel(upRightPixel(p), used, unUsed)
	}
	if p.Y > 0 && p.X > 0 {
		s.filterPixel(downLeftPixel(p), used, unUsed)
	}
	if p.Y > 0 && p.X < s.Size-1 {
		s.filterPixel(downRightPixel(p), used, unUsed)
	}
	return used, unUsed
}

func (s *Surface) filterPixel(p Coord2D, used, unUsed *list.List) {
	if s.IsUsed(p.X, p.Y) {
		used.PushBack(p)
	} else {
		unUsed.PushBack(p)
	}
}

/*
++++
/\
|
|
|
|
y
  x  -----------> +++

*/
func leftPixel(p Coord2D) Coord2D {
	return Coord2D{p.X - 1, p.Y}
}

func rightPixel(p Coord2D) Coord2D {
	return Coord2D{p.X + 1, p.Y}
}

func upPixel(p Coord2D) Coord2D {
	return Coord2D{p.X, p.Y + 1}
}

func downPixel(p Coord2D) Coord2D {
	return Coord2D{p.X, p.Y - 1}
}

func downLeftPixel(p Coord2D) Coord2D {
	return Coord2D{p.X - 1, p.Y - 1}
}

func downRightPixel(p Coord2D) Coord2D {
	return Coord2D{p.X + 1, p.Y - 1}
}

func upLeftPixel(p Coord2D) Coord2D {
	return Coord2D{p.X - 1, p.Y + 1}
}

func upRightPixel(p Coord2D) Coord2D {
	return Coord2D{p.X + 1, p.Y + 1}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
