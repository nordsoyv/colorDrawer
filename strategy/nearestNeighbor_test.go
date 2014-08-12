package strategy

import (
	"testing"
	"github.com/nordsoyv/colorDrawer/config"
	_ "container/list"
)

func TestNew(t *testing.T) {
	c := config.Read("config_test.json")
	n := NearestNeighbor(c)
	if n == nil {
		t.Error("New failed")
	}
}

func setup() nearestNeighborStrategy {
	c := config.Read("config_test.json")
	colorStrat := NearestNeighbor(c)
	n, _ := colorStrat.(nearestNeighborStrategy)
	return n
}


