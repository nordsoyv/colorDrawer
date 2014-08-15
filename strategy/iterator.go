package strategy

import (
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"image/color"
)

type iteratorStrategy struct {
	fileName string
	cube     *colorCube.ColorCube
}

func Iterator(c config.Config) ColorStrategy {
	var s iteratorStrategy
	s.fileName = c.OutputFilename
	s.cube = colorCube.New(uint8(c.ColorCubeBitSize))
	return s
}

func (s iteratorStrategy) GenerateImage(doneChan chan bool, imageUpdateChan chan ImageUpdate) {
	imageSize := 1 << uint(((s.cube.BitSize + s.cube.BitSize + s.cube.BitSize) / 2))
	surface := workSurface.New(imageSize)
	colorStorage := make([]color.RGBA, imageSize*imageSize)
	nextColorSpace := 0
	for x := 0; x < s.cube.SideSize; x++ {
		for y := 0; y < s.cube.SideSize; y++ {
			for z := 0; z < s.cube.SideSize; z++ {
				colorStorage[nextColorSpace] = s.cube.GetColor(x, y, z)
				nextColorSpace++
			}
		}
	}
	nextColorSpace = 0
	for x := 0; x < surface.Size; x++ {
		for y := 0; y < surface.Size; y++ {
			surface.SetColor(x, y, colorStorage[nextColorSpace])
			surface.SetUsed(x, y)
			nextColorSpace++
		}
	}
	surface.ToPng(s.fileName)
	doneChan <- true
}
