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
		t.Error("FirstColor is used")
	}
	if w.pixels[7][7].Used == true{
		t.Error("LastColor is used")
	}
}

func TestSetGetColor(t *testing.T){
	w := New(8)
	w.SetColorRGB(0,0,uint8(100),uint8(100),uint8(100))
	color := w.GetColor(0,0)
	if color.R != 100 && color.G != 100 &&  color.B != 100 {
		t.Error("Wrong color : ", color)
	}
}



func TestSetUsed(t *testing.T){
	w := New(8)
	w.SetUsed(0,0)
	if !w.IsUsed(0,0) {
		t.Error("Not Used when should be ")
	}
}


func TestFindNeighborPixelsInCenter(t *testing.T) {
	p := Coord{3, 3}
	surface := New(64)
	used, unUsed := surface.FindNeighborPixels(p)
	if used == nil {
		t.Error("findNeighborPixels used failed")
	}
	if unUsed == nil {
		t.Error("findNeighborPixels unUsed failed")
	}
	if unUsed.Len() != 8 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
	if used.Len() != 0 {
		t.Error("findNeighborPixels pixels in used list")
	}
}

func TestFindNeighborPixelsInCorners(t *testing.T) {
	surface := New(64)
	p := Coord{0, 0}
	_, unUsed := surface.FindNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
	p = Coord{0, 63}
	_, unUsed = surface.FindNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
	p = Coord{63, 0}
	_, unUsed = surface.FindNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
	p = Coord{63, 63}
	_, unUsed = surface.FindNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
}

func TestFindNeighborPixelsOnEdge(t *testing.T) {
	surface := New(64)
	p := Coord{0, 5}
	_, unUsed := surface.FindNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 0 5")
	}
	p = Coord{5, 0}
	_, unUsed = surface.FindNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 5 0")
	}
	p = Coord{63, 5}
	_, unUsed = surface.FindNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 15 5")
	}
	p = Coord{5, 63}
	_, unUsed = surface.FindNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 5 15", unUsed.Len())
	}
}
