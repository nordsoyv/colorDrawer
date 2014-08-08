package colorCube

import (
	"image/color"
	"testing"
)

func TestNewSetsSize(t *testing.T) {
	colorCube := New(uint8(4))
	if colorCube.SideSize != 16 {
		t.Fail()
	}
}

func TestNewCreatesCubeAllFalse(t *testing.T) {
	colorCube := New(uint8(5))
	for r := 0; r < colorCube.SideSize; r++ {
		for g := 0; g < colorCube.SideSize; g++ {
			for b := 0; b < colorCube.SideSize; b++ {
				if colorCube.Cube[r][g][b] == true {
					t.Fail()
				}
			}
		}
	}
}

func TestCanSetValueToUsed(t *testing.T) {
	colorCube := New(uint8(5))
	colorCube.SetUsed(1, 1, 1)
	if !colorCube.IsUsed(1, 1, 1) {
		t.Fail()
	}
}

func TestCubeShouldNotBeUsedOnCreation(t *testing.T) {
	colorCube := New(uint8(8))
	for r := 0; r < colorCube.SideSize; r++ {
		for g := 0; g < colorCube.SideSize; g++ {
			for b := 0; b < colorCube.SideSize; b++ {
				if colorCube.IsUsed(r, g, b) {
					t.Fail()
				}
			}
		}
	}
}

func TestCanGetColorForCoord(t *testing.T) {
	colorCube := New(uint8(4))
	col := colorCube.GetColor(0, 0, 0) //black
	if col.R != 0 && col.B != 0 && col.G != 0 {
		t.Error("Not black color : ", col)
	}
	col = colorCube.GetColor(15, 15, 15) //white
	if col.R != 255 && col.B != 255 && col.G != 255 {
		t.Error("Not white color : ", col)
	}
	col = colorCube.GetColor(15, 0, 0) //just red
	if col.R != 255 && col.B == 0 && col.G == 0 {
		t.Error("Not red color : ", col)
	}
	col = colorCube.GetColor(7, 7, 7) //grey
	if col.R != 119 && col.B != 119 && col.G != 119 {
		t.Error("Not grey color : ", col)
	}
}

func TestGetIndexForColor(t *testing.T) {
	colorCube := New(uint8(4))

	x, y, z := colorCube.GetIndexForColor(color.RGBA{255, 255, 255, 255})
	if x != colorCube.SideSize-1 {
		t.Error("X is not 16 but ", x)
	}
	if y != colorCube.SideSize-1 {
		t.Error("Y is not 16 but ", y)
	}
	if z != colorCube.SideSize-1 {
		t.Error("X is not 16 but ", z)
	}
}
