package strategy

import (
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"container/list"
	"github.com/nordsoyv/colorDrawer/config"
	"image/color"
)

func NearestNeighbor(c config.Config) ColorStrategy {
	imageSize := 1 << uint(((c.ColorCubeBitSize + c.ColorCubeBitSize + c.ColorCubeBitSize) / 2))
	surface := workSurface.New(imageSize)


	return nearestNeighborStrategy{list.New(), color.RGBA{255, 255, 255, 255  }, surface}
}

type pixel struct {
	x, y int
}

type nearestNeighborStrategy struct {
	pixelBuffer *list.List
	startColor color.RGBA
	surface    workSurface.Surface
}

func (n nearestNeighborStrategy) GenerateImage(cube *colorCube.ColorCube) workSurface.Surface {
	n.addPixelToDraw(pixel{0, 0})
	for n.pixelBuffer.Len() > 0 {
		nextPixel := n.getNextPixel()
		_, unUsedPixels := n.findNeighborPixels(nextPixel)
		n.addPixelsToDraw(unUsedPixels)

		//get average color for used neighbor pixels
		//find index for this color in colorcube


	}


	return n.surface;

}

func (n nearestNeighborStrategy) addPixelToDraw(p pixel) {
	n.pixelBuffer.PushBack(p)
}

func (n nearestNeighborStrategy) addPixelsToDraw(l *list.List) {
	n.pixelBuffer.PushBackList(l)
}


func (n nearestNeighborStrategy) getNextPixel() (pixel) {
	elem := n.pixelBuffer.Front()
	n.pixelBuffer.Remove(elem)
	p, ok := elem.Value.(pixel)
	if !ok {
		panic("Not a pixel in list!")
	}
	return p
}

func (n nearestNeighborStrategy) findNeighborPixels(p pixel) (used, unUsed  *list.List) {
	unUsed = list.New()
	used = list.New()
	if p.x > 0 {
		n.filterPixel(leftPixel(p), used, unUsed)
	}
	if p.x < n.surface.Size-1 {
		n.filterPixel(rightPixel(p), used, unUsed)
	}
	if p.y < n.surface.Size-1 {
		n.filterPixel(upPixel(p), used, unUsed)
	}
	if p.y > 0 {
		n.filterPixel(downPixel(p), used, unUsed)
	}
	if p.y < n.surface.Size-1 && p.x > 0 {
		n.filterPixel(upLeftPixel(p), used, unUsed)
	}
	if p.y < n.surface.Size-1 && p.x < n.surface.Size-1 {
		n.filterPixel(upRightPixel(p), used, unUsed)
	}
	if p.y > 0 && p.x > 0 {
		n.filterPixel(downLeftPixel(p), used, unUsed)
	}
	if p.y > 0 && p.x < n.surface.Size-1 {
		n.filterPixel(downRightPixel(p), used, unUsed)
	}
	return used, unUsed
}

func (n nearestNeighborStrategy) filterPixel(p pixel, used, unUsed *list.List) {
	if n.surface.IsUsed(p.x, p.y) {
		used.PushBack(p)
	}else {
		unUsed.PushBack(p)
	}
}

/*
++++
/\
|
|
|
|
y
  x  -----------> +++

*/
func leftPixel(p pixel) pixel {
	return pixel{p.x - 1, p.y}
}

func rightPixel(p pixel) pixel {
	return pixel{p.x + 1, p.y}
}

func upPixel(p pixel) pixel {
	return pixel{p.x, p.y + 1}
}

func downPixel(p pixel) pixel {
	return pixel{p.x, p.y - 1}
}

func downLeftPixel(p pixel) pixel {
	return pixel{p.x - 1, p.y - 1}
}

func downRightPixel(p pixel) pixel {
	return pixel{p.x + 1, p.y - 1}
}

func upLeftPixel(p pixel) pixel {
	return pixel{p.x - 1, p.y + 1}
}

func upRightPixel(p pixel) pixel {
	return pixel{p.x + 1, p.y + 1}
}
