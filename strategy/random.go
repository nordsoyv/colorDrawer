package strategy

import (
	"math/rand"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"github.com/nordsoyv/colorDrawer/colorCube"
)


func Random() ColorStrategy {
	var s randomImageStrategy
	return s
}

type randomImageStrategy struct {
}

func (s randomImageStrategy) GenerateImage(cube *colorCube.ColorCube) workSurface.Surface {
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
	return surface
}


