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

	n.surface.SetColor(0, 0, color.RGBA{uint8(0), uint8(0), uint8(0), 255})
	for n.pixelBuffer.Len() > 0 {
		nextPixel := n.getNextPixel()
		usedPixels, unUsedPixels := n.findNeighborPixels(nextPixel)
		n.addPixelsToDraw(unUsedPixels)

		//get average color for used neighbor pixels
		avgColor := n.getAverageColor(usedPixels)
		//find index for this color in colorcube
		x, y, z := cube.GetIndexForColor(avgColor)
		//if color at that index is not used
		if !cube.IsUsed(x, y, z) {
			cube.SetUsed(x, y, z)
			n.surface.SetColor(nextPixel.x, nextPixel.y, cube.GetColor(x, y, z))
			continue
		}else {
			//  find nearest free color in cube
			//	set as used, and color surface with it. continue loop
		}



	}


	return n.surface;

}

func findUnusedColorsInTop(startX, startY, startZ, distFromCenter int, cube *colorCube.ColorCube) (foundIt bool, foundX, foundY , foundZ int) {
	minX := startX - distFromCenter
	minZ := startZ - distFromCenter
	maxX := startX + distFromCenter
	maxZ := startZ + distFromCenter
	yPos := startY + distFromCenter
	for x := minX ; x <= maxX; x++ {
		for z := minZ; z <= maxZ; z++ {
			if !cube.IsUsed(x, yPos, z) {
				return true, x, yPos, z
			}
		}
	}
	return false, 0, 0, 0
}




func (n nearestNeighborStrategy) getAverageColor(l *list.List) color.RGBA {
	var totR, totG, totB int
	totR = 0
	totG = 0
	totB = 0
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(pixel)
		col := n.surface.GetColor(p.x, p.y)
		totR += int(col.R)
		totG += int(col.G)
		totB += int(col.B)
	}
	numElem := l.Len()
	return color.RGBA { uint8(totR / numElem), uint8(totG / numElem), uint8(totB / numElem), 255  }

}

func (n nearestNeighborStrategy) addPixelToDraw(p pixel) {
	n.pixelBuffer.PushBack(p)
}

func (n nearestNeighborStrategy) addPixelsToDraw(l *list.List) {
	n.pixelBuffer.PushBackList(l)
}


func (n nearestNeighborStrategy) getNextPixel() (pixel) {
	return getFrontPixelInList(n.pixelBuffer)
}

func getFrontPixelInList(l *list.List) pixel {
	elem := l.Front()
	l.Remove(elem)
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
