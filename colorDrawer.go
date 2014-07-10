package main

import (
	//	"fmt"
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"image/color"
	"math/rand"
)

type ColorStrategy interface {
	GenerateImage(*colorCube.ColorCube) workSurface.Surface
}

type IteratorStrategy struct {
}

func (s IteratorStrategy) GenerateImage(cube *colorCube.ColorCube) workSurface.Surface {
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
	return surface
}

type RandomImageStrategy struct {
}

func (s RandomImageStrategy) GenerateImage(cube *colorCube.ColorCube) workSurface.Surface {
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

func randomColor() color.RGBA {
	R := uint8(rand.Intn(256))
	G := uint8(rand.Intn(256))
	B := uint8(rand.Intn(256))
	return color.RGBA{R, G, B, 255}
}

func main() {

	configuration := config.Read("config.json")

	cube := colorCube.New(uint8(configuration.ColorCubeBitSize))

	var strategy IteratorStrategy
	surface := strategy.GenerateImage(cube)
	surface.ToPng(configuration.OutputFilename)

}
