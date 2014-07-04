package imageutil

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"bufio"
)

func CreateRandomImage(size int) (image.Image) {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for y := 0 ; y < size; y++ {
		for x := 0 ; x < size; x++ {
			img.Set(x, y, randomColor())
		}
	}
	return img
}

func WriteImageToDisk(path string, img image.Image) {
	outfile , err := os.Create(path)
	check(err)

	defer outfile.Close()

	writer := bufio.NewWriter(outfile)

	png.Encode(writer, img)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func randomColor() (color.RGBA){
	R := uint8(rand.Intn(256))
	G := uint8(rand.Intn(256))
	B := uint8(rand.Intn(256))
	return color.RGBA{R, G, B, 255}
}
