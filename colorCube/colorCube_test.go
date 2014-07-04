package colorCube

import (
	"testing"
	"fmt"
	"math"
)

func TestNewSetsSize(t *testing.T) {
	colorCube := New(5)
	if (colorCube.Size != 5 ) {
		t.Fail()
	}
}

func TestNewCreatesCubeAllFalse(t *testing.T) {
	size := 5
	colorCube := New(5)
	for r := 0; r < size; r++ {
		for g := 0; g < size; g++ {
			for b := 0; b < size; b++ {
				if colorCube.Cube[r][g][b] == true {
					t.Fail()
				}
			}
		}
	}
}

func TestCanSetValueToUsed(t *testing.T) {
	colorCube := New(5)
	colorCube.SetUsed(1,1,1)
	if ! colorCube.IsUsed(1,1,1){
		t.Fail()
	}
}

func TestCubeShouldNotBeUsedOnCreation(t *testing.T) {
	size := 256
	colorCube := New(size)
	for r := 0; r < size; r++ {
		for g := 0; g < size; g++ {
			for b := 0; b < size; b++ {
				if colorCube.IsUsed(r, g, b) {
					t.Fail()
				}
			}
		}
	}
}

func TestCanGetColorForCoord(t *testing.T) {
	colorCube := New(5)
	color := colorCube.GetColor(0, 0, 0) //black
	if color.R != 0 && color.B != 0 && color.G != 0 {
		t.Fail()
	}
	color = colorCube.GetColor(4, 4, 4) //white
	if color.R != 255 && color.B != 255 && color.G != 255 {
		t.Fail()
	}

	color = colorCube.GetColor(4, 0, 0) //just red
	if color.R != 255 && color.B != 0 && color.G != 0 {
		t.Fail()
	}
	color = colorCube.GetColor(2, 2, 2) //grey
	if color.R != 127 && color.B != 127 && color.G != 127 {
		fmt.Printf("color : %d \n" , color.R )
		t.Fail()
	}
}

func TestCanGetColorFor(t *testing.T) {
	a:=  math.Pow(2,4)
	fmt.Println("a : " ,a)
}
