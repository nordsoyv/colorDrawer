package main

import (
	//	"fmt"
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/strategy"
)

func main() {
	configuration := config.Read("config.json")
	cube := colorCube.New(uint8(configuration.ColorCubeBitSize))
	strat := strategy.NearestNeighbor(configuration)
	surface := strat.GenerateImage(cube)
	surface.ToPng(configuration.OutputFilename)
}
