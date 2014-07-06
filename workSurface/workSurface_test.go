package workSurface

import "testing"

func TestPixelIsZeroOnCreate(t *testing.T) {
	pixel := new(pixel)
	if !( pixel.Color.R == 0 && pixel.Color.G == 0 && pixel.Color.B == 0) {
		t.Error("color not 0")
	}
	if pixel.Used {
		t.Error("Is used on init")
	}
}

func TestNew(t *testing.T){
	w := New(8)
	if w.Size != 8 {
		t.Error("Wrong size")
	}
	if w.pixels[0][0].Used == true{
		t.Error("FirstColor is used size")
	}
	if w.pixels[7][7].Used == true{
		t.Error("LastColor is used size")
	}
}

