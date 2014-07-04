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

	imageSize := (configuration.ColorCubeSide + configuration.ColorCubeSide + configuration.ColorCubeSide) / 2
	fmt.Println("imageSize : ", imageSize)
	img := imageutil.CreateRandomImage(int(imageSize))
	imageutil.WriteImageToDisk(configuration.OutputFilename, img)

}

