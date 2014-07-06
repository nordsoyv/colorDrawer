package main

import (
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/imageutil"
	//	"github.com/nordsoyv/colorDrawer/colorCube"
	"fmt"
)

func main() {

	configuration := config.Read("config.json")

	//	cube := colorCube.New(configuration.ColorCubeSide)

	imageSize := (configuration.ColorCubeBitSize + configuration.ColorCubeBitSize + configuration.ColorCubeBitSize) / 2
	fmt.Println("imageSize : ", imageSize)
	img := imageutil.CreateRandomImage(imageSize)
	imageutil.WriteImageToDisk(configuration.OutputFilename, img)

}

