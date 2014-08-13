package strategy

import (
	"container/list"
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"image/color"
	"fmt"
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
	n.addPixelToDraw(workSurface.Coord2D{0, 0})

	n.surface.SetColor(0, 0, color.RGBA{uint8(255), uint8(255), uint8(255), 255})

	totalNumberOfPixels := n.surface.Size * n.surface.Size
	currentPixel := 1

	for n.pixelBuffer.Len() > 0 {
		nextPixel := n.getNextPixel()
		usedPixels, unUsedPixels := n.surface.FindNeighborPixels(nextPixel)

		numAdded := n.addPixelsToDraw(unUsedPixels)

		fmt.Printf("Current pixel (%v , %v) %v / %v . Adding %v new pixels, used : %v, unUsed : %v , queueLenght %v\n", nextPixel.X, nextPixel.Y, currentPixel, totalNumberOfPixels, numAdded, usedPixels.Len(), unUsedPixels.Len(), n.pixelBuffer.Len())
		currentPixel++

		//get average color for used neighbor pixels
		var avgColor color.RGBA
		if usedPixels.Len() > 0 {
			avgColor = n.getAverageColor(usedPixels)
		} else {
			avgColor = n.surface.GetColor(nextPixel.X, nextPixel.Y)
		}
		//find index for this color in colorcube
		x, y, z := cube.GetIndexForColor(avgColor)
		//if color at that index is not used
		if !cube.IsUsed(x, y, z) {
			cube.SetUsed(x, y, z)
			n.surface.SetColor(nextPixel.X, nextPixel.Y, cube.GetColor(x, y, z))
		} else {
			//  find nearest free color in cube
			found, foundX, foundY, foundZ := cube.FindUnusedColorInCube(x, y, z)
			if found {
				//	set as used, and color surface with it. continue loop
				cube.SetUsed(foundX, foundY, foundZ)
				n.surface.SetColor(nextPixel.X, nextPixel.Y, cube.GetColor(foundX, foundY, foundZ))
			}else {
				panic("Coudnt fint color!")
			}
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
		p := e.Value.(workSurface.Coord2D)
		col := n.surface.GetColor(p.X, p.Y)
		totR += int(col.R)
		totG += int(col.G)
		totB += int(col.B)
	}
	numElem := l.Len()
	return color.RGBA{uint8(totR / numElem), uint8(totG / numElem), uint8(totB / numElem), 255}

}

func (n nearestNeighborStrategy) addPixelToDraw(pixelToAdd workSurface.Coord2D) bool {
	if n.surface.IsUsed(pixelToAdd.X, pixelToAdd.Y) {
		return false
	}
	for e := n.pixelBuffer.Front(); e != nil; e = e.Next() {
		p := e.Value.(workSurface.Coord2D)
		if p.X == pixelToAdd.X && p.Y == pixelToAdd.Y {
			//all ready in queue
			return false
		}
	}
	n.pixelBuffer.PushBack(pixelToAdd)
	return true
}

func (n nearestNeighborStrategy) addPixelsToDraw(l *list.List) int {
	numAdded := 0
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(workSurface.Coord2D)
		if n.addPixelToDraw(p) {
			numAdded++
		}
	}
	return numAdded
}

func (n nearestNeighborStrategy) getNextPixel() workSurface.Coord2D {
	elem := n.pixelBuffer.Front()
	p, ok := elem.Value.(workSurface.Coord2D)
	if !ok {
		panic("Not a pixel in list!")
	}
	n.pixelBuffer.Remove(elem)
	return p
}

