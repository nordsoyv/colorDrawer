package colorCube

import (
	"testing"
)

func TestFindUnusedColorsInTopAllUsed(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	foundIt, _, _, _ := cube.findUnusedColorsInTop(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
}

func TestFindUnusedColorsInTopOneUnused(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	cube.SetUnUsed(5, 6, 5)
	foundIt, _, _, _ := cube.findUnusedColorsInTop(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5, 6, 5)
	cube.SetUnUsed(5, 7, 5)
	foundIt, _, _, _ = cube.findUnusedColorsInTop(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
	foundIt, _, _, _ = cube.findUnusedColorsInTop(5, 5, 5, 2)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5, 7, 5)
	cube.SetUnUsed(4, 6, 4)
	foundIt, _, _, _ = cube.findUnusedColorsInTop(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4, 6, 4)
	cube.SetUnUsed(1, 6, 1)
	foundIt, _, _, _ = cube.findUnusedColorsInTop(0, 5, 0, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
}

func TestFindUnusedColorsInBottomAllUsed(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	foundIt, _, _, _ := cube.findUnusedColorsInBottom(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
}

func TestFindUnusedColorsInBottomOneUnused(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	cube.SetUnUsed(5, 4, 5)
	foundIt, _, _, _ := cube.findUnusedColorsInBottom(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5, 4, 5)
	cube.SetUnUsed(5, 3, 5)
	foundIt, _, _, _ = cube.findUnusedColorsInBottom(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
	foundIt, _, _, _ = cube.findUnusedColorsInBottom(5, 5, 5, 2)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(5, 3, 5)
	cube.SetUnUsed(4, 4, 4)
	foundIt, _, _, _ = cube.findUnusedColorsInBottom(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4, 4, 4)
	cube.SetUnUsed(1, 4, 1)
	foundIt, _, _, _ = cube.findUnusedColorsInBottom(0, 5, 0, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
}

func TestFindUnusedColorsInLeftAllUsed(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	foundIt, _, _, _ := cube.findUnusedColorsInLeft(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
}

func TestFindUnusedColorsInLeftOneUnused(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	cube.SetUnUsed(4, 5, 5)
	foundIt, _, _, _ := cube.findUnusedColorsInLeft(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4, 5, 5)
	cube.SetUnUsed(3, 5, 5)
	foundIt, _, _, _ = cube.findUnusedColorsInLeft(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
	foundIt, _, _, _ = cube.findUnusedColorsInLeft(5, 5, 5, 2)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(3, 5, 5)
	cube.SetUnUsed(4, 4, 4)
	foundIt, _, _, _ = cube.findUnusedColorsInLeft(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4, 4, 4)
	cube.SetUnUsed(4, 1, 1)
	foundIt, _, _, _ = cube.findUnusedColorsInLeft(5, 0, 0, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
}

func TestFindUnusedColorsInRightAllUsed(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	foundIt, _, _, _ := cube.findUnusedColorsInRight(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
}

func TestFindUnusedColorsInRightOneUnused(t *testing.T) {
	cube := New(uint8(5))
	setCubeAsUsed(cube)
	cube.SetUnUsed(4, 5, 5)
	foundIt, _, _, _ := cube.findUnusedColorsInRight(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4, 5, 5)
	cube.SetUnUsed(3, 5, 5)
	foundIt, _, _, _ = cube.findUnusedColorsInRight(5, 5, 5, 1)
	if foundIt {
		t.Error("Should not find an unused color")
	}
	foundIt, _, _, _ = cube.findUnusedColorsInRight(5, 5, 5, 2)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(3, 5, 5)
	cube.SetUnUsed(4, 4, 4)
	foundIt, _, _, _ = cube.findUnusedColorsInRight(5, 5, 5, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
	cube.SetUsed(4, 4, 4)
	cube.SetUnUsed(4, 1, 1)
	foundIt, _, _, _ = cube.findUnusedColorsInRight(5, 0, 0, 1)
	if !foundIt {
		t.Error("Should find an unused color")
	}
}

func setCubeAsUsed(cube *ColorCube) {
	for x := 0; x < cube.SideSize; x++ {
		for y := 0; y < cube.SideSize; y++ {
			for z := 0; z < cube.SideSize; z++ {
				cube.Cube[x][y][z] = true
			}
		}
	}
}
