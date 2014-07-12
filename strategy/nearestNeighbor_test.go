package strategy

import (
	"testing"
	"github.com/nordsoyv/colorDrawer/config"
	 _ "container/list"
)


func TestNew(t *testing.T){
	c := config.Read("config_test.json")
	n := NearestNeighbor(c)
	if n == nil {
		t.Error("New failed")
	}
}

func TestFindNeighborPixelsInCenter(t *testing.T ){
	p := pixel{3,3}
	c := config.Read("config_test.json")
	colorStrat := NearestNeighbor(c)
	n,_ := colorStrat.(nearestNeighborStrategy)
	used,unUsed := n.findNeighborPixels(p)
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

func TestFindNeighborPixelsInCorners(t *testing.T ){
	t.SkipNow()
	c := config.Read("config_test.json")
	colorStrat := NearestNeighbor(c)
	n,_ := colorStrat.(nearestNeighborStrategy)
	p := pixel{0,0}
	_,unUsed := n.findNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
	p = pixel{0,15}
	_,unUsed = n.findNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
	p = pixel{15,0}
	_,unUsed = n.findNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
	p = pixel{15,15}
	_,unUsed = n.findNeighborPixels(p)
	if unUsed.Len() != 3 {
		t.Error("findNeighborPixels not all pixels in unused list")
	}
}

func TestFindNeighborPixeOnEdge(t *testing.T ){
	c := config.Read("config_test.json")
	colorStrat := NearestNeighbor(c)
	n,_ := colorStrat.(nearestNeighborStrategy)
	p := pixel{0,5}
	_,unUsed := n.findNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 0 5")
	}
	p = pixel{5,0}
	_,unUsed = n.findNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 5 0")
	}
	p = pixel{63,5}
	_,unUsed = n.findNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 15 5")
	}
	p = pixel{5,63}
	_,unUsed = n.findNeighborPixels(p)
	if unUsed.Len() != 5 {
		t.Error("findNeighborPixels 5 15", unUsed.Len())
	}
}
