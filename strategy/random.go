package strategy

import (
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"math/rand"
)

func Random(c config.Config) ColorStrategy {
	var s randomImageStrategy
	s.fileName = c.OutputFilename
	return s
}

type randomImageStrategy struct {
	fileName string
}

func (s randomImageStrategy) GenerateImage(cube *colorCube.ColorCube) {
	imageSize := 1 << uint(((cube.BitSize + cube.BitSize + cube.BitSize) / 2))
	surface := workSurface.New(imageSize)
	for x := 0; x < surface.Size; x++ {
		for y := 0; y < surface.Size; y++ {
			R := uint8(rand.Intn(256))
			G := uint8(rand.Intn(256))
			B := uint8(rand.Intn(256))
			surface.SetColorRGB(x, y, R, G, B)
			surface.SetUsed(x, y)
		}
	}
	surface.ToPng(s.fileName)
}
