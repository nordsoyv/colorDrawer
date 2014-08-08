package strategy

import (
	"testing"
	"github.com/nordsoyv/colorDrawer/colorCube"
)

func TestFindUnusedColorsInTopAllUsed(t *testing.T) {
	cube := colorCube.New(uint8(5))
	setCubeAsUsed(cube)
	foundIt, _, _, _ := findUnusedColorsInTop(5, 5, 5, 1, cube)
	if foundIt {
		t.Error("Should not find an unused color")
	}
}

func TestFindUnusedColorsInTopOneUnused(t *testing.T) {
	cube := colorCube.New(uint8(5))
	setCubeAsUsed(cube)
	cube.SetUnUsed(5, 6, 5)
	foundIt, _, _, _ := findUnusedColorsInTop(5, 5, 5, 1, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5,6,5)
	cube.SetUnUsed(5, 7, 5)
	foundIt, _, _, _ = findUnusedColorsInTop(5, 5, 5, 1, cube)
	if foundIt {
		t.Error("Should not find an unused color")
	}
	foundIt, _, _, _ = findUnusedColorsInTop(5, 5, 5, 2, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5,7,5)
	cube.SetUnUsed(4,6,4)
	foundIt, _, _, _ = findUnusedColorsInTop(5, 5, 5, 1, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4,6,4)
	cube.SetUnUsed(1,6,1)
	foundIt, _, _, _ = findUnusedColorsInTop(0, 5, 0, 1, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
}

func TestFindUnusedColorsInBottomAllUsed(t *testing.T) {
	cube := colorCube.New(uint8(5))
	setCubeAsUsed(cube)
	foundIt, _, _, _ := findUnusedColorsInBottom(5, 5, 5, 1, cube)
	if foundIt {
		t.Error("Should not find an unused color")
	}
}

func TestFindUnusedColorsInBottomOneUnused(t *testing.T) {
	cube := colorCube.New(uint8(5))
	setCubeAsUsed(cube)
	cube.SetUnUsed(5, 4, 5)
	foundIt, _, _, _ := findUnusedColorsInBottom(5, 5, 5, 1, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5,4,5)
	cube.SetUnUsed(5, 3, 5)
	foundIt, _, _, _ = findUnusedColorsInBottom(5, 5, 5, 1, cube)
	if foundIt {
		t.Error("Should not find an unused color")
	}
	foundIt, _, _, _ = findUnusedColorsInBottom(5, 5, 5, 2, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5,3,5)
	cube.SetUnUsed(4,4,4)
	foundIt, _, _, _ = findUnusedColorsInBottom(5, 5, 5, 1, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4,4,4)
	cube.SetUnUsed(1,4,1)
	foundIt, _, _, _ = findUnusedColorsInBottom(0, 5, 0, 1, cube)
	if !foundIt {
		t.Error("Should find an unused color")
	}
}


func setCubeAsUsed(cube *colorCube.ColorCube) {
	for x := 0; x < cube.SideSize; x++ {
		for y := 0; y < cube.SideSize; y++ {
			for z := 0; z < cube.SideSize; z++ {
				cube.Cube[x][y][z] = true;
			}
		}
	}
}

