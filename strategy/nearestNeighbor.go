package strategy

import (
	"container/list"
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"image/color"
)

func NearestNeighbor(c config.Config) ColorStrategy {
	imageSize := 1 << uint(((c.ColorCubeBitSize + c.ColorCubeBitSize + c.ColorCubeBitSize) / 2))
	surface := workSurface.New(imageSize)

	return nearestNeighborStrategy{list.New(), color.RGBA{255, 255, 255, 255}, surface}
}

type pixel struct {
	x, y int
}

type nearestNeighborStrategy struct {
	pixelBuffer *list.List
	startColor  color.RGBA
	surface     workSurface.Surface
}

func (n nearestNeighborStrategy) GenerateImage(cube *colorCube.ColorCube) workSurface.Surface {
	n.addPixelToDraw(workSurface.Coord{0, 0})

	n.surface.SetColor(0, 0, color.RGBA{uint8(0), uint8(0), uint8(0), 255})
	for n.pixelBuffer.Len() > 0 {
		nextPixel := n.getNextPixel()
		usedPixels, unUsedPixels := n.surface.FindNeighborPixels(nextPixel)
		n.addPixelsToDraw(unUsedPixels)

		//get average color for used neighbor pixels
		avgColor := n.getAverageColor(usedPixels)
		//find index for this color in colorcube
		x, y, z := cube.GetIndexForColor(avgColor)
		//if color at that index is not used
		if !cube.IsUsed(x, y, z) {
			cube.SetUsed(x, y, z)
			n.surface.SetColor(nextPixel.X, nextPixel.Y, cube.GetColor(x, y, z))
			continue
		} else {
			//  find nearest free color in cube
			//	set as used, and color surface with it. continue loop
		}

	}

	return n.surface

}

func (n nearestNeighborStrategy) getAverageColor(l *list.List) color.RGBA {
	var totR, totG, totB int
	totR = 0
	totG = 0
	totB = 0
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(workSurface.Coord)
		col := n.surface.GetColor(p.X, p.Y)
		totR += int(col.R)
		totG += int(col.G)
		totB += int(col.B)
	}
	numElem := l.Len()
	return color.RGBA{uint8(totR / numElem), uint8(totG / numElem), uint8(totB / numElem), 255}

}

func (n nearestNeighborStrategy) addPixelToDraw(p workSurface.Coord) {
	n.pixelBuffer.PushBack(p)
}

func (n nearestNeighborStrategy) addPixelsToDraw(l *list.List) {
	n.pixelBuffer.PushBackList(l)
}

func (n nearestNeighborStrategy) getNextPixel() workSurface.Coord {
	return getFrontPixelInList(n.pixelBuffer)
}

func getFrontPixelInList(l *list.List) workSurface.Coord {
	elem := l.Front()
	l.Remove(elem)
	p, ok := elem.Value.(workSurface.Coord)
	if !ok {
		panic("Not a pixel in list!")
	}
	return p
}
