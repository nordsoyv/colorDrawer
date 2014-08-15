package strategy

import (
	"github.com/nordsoyv/colorDrawer/colorCube"
)

type ColorStrategy interface {
	GenerateImage(*colorCube.ColorCube)
}

type ImageUpdate struct {
	X, Y    int
	R, G, B int
}
