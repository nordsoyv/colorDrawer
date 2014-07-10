package strategy

import (
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/workSurface")

type ColorStrategy interface {
	GenerateImage(*colorCube.ColorCube) workSurface.Surface
}
