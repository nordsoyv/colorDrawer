package strategy

import (
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"image/color"
)

type iteratorStrategy struct {
	fileName string
}

func Iterator(c config.Config) ColorStrategy {
	var s iteratorStrategy
	s.fileName = c.OutputFilename
	return s
}

func (s iteratorStrategy) GenerateImage(cube *colorCube.ColorCube) {
	imageSize := 1 << uint(((cube.BitSize + cube.BitSize + cube.BitSize) / 2))
	surface := workSurface.New(imageSize)
	colorStorage := make([]color.RGBA, imageSize*imageSize)
	nextColorSpace := 0
	for x := 0; x < cube.SideSize; x++ {
		for y := 0; y < cube.SideSize; y++ {
			for z := 0; z < cube.SideSize; z++ {
				colorStorage[nextColorSpace] = cube.GetColor(x, y, z)
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
}
