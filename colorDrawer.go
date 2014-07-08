package main

import (
	"github.com/nordsoyv/colorDrawer/config"
	//	"github.com/nordsoyv/colorDrawer/colorCube"
	"fmt"
	"github.com/nordsoyv/colorDrawer/workSurface"
)

func main() {

	configuration := config.Read("config.json")

	//	cube := colorCube.New(uint8(configuration.ColorCubeBitSize))

	imageSize := 1 << uint(( (configuration.ColorCubeBitSize + configuration.ColorCubeBitSize + configuration.ColorCubeBitSize) / 2))
	fmt.Println("imageSize : ", imageSize)
	//	img := imageutil.CreateRandomImage(imageSize)
	//	imageutil.WriteImageToDisk(configuration.OutputFilename, img)

	work := workSurface.New(imageSize)

	work.FillWithRandomColors()
	work.ToPng(configuration.OutputFilename)



}

