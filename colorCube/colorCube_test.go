package colorCube

import "testing"

func TestNewSetsSize(t *testing.T) {
	colorCube := New(uint8(4))
	if (colorCube.SideSize != 16 ) {
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
	color := colorCube.GetColor(0, 0, 0) //black
	if color.R != 0 && color.B != 0 && color.G != 0 {
		t.Error("Not black color : ", color)
	}
	color = colorCube.GetColor(15, 15, 15) //white
	if color.R != 255 && color.B != 255 && color.G != 255 {
		t.Error("Not white color : ", color)
	}
	color = colorCube.GetColor(15, 0, 0) //just red
	if color.R != 255 && color.B == 0 && color.G == 0 {
		t.Error("Not red color : ", color)
	}
	color = colorCube.GetColor(7, 7, 7) //grey
	if color.R != 119 && color.B != 119 && color.G != 119 {
		t.Error("Not grey color : ", color)
	}
}
