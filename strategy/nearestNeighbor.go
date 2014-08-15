package strategy

import (
	"container/list"
	// "fmt"
	"github.com/nordsoyv/colorDrawer/colorCube"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/workSurface"
	"image/color"
	"math/rand"
)

func NearestNeighbor(c config.Config) ColorStrategy {
	imageSize := 1 << uint(((c.ColorCubeBitSize + c.ColorCubeBitSize + c.ColorCubeBitSize) / 2))
	surface := workSurface.New(imageSize)
	fileName := c.OutputFilename
	cube := colorCube.New(uint8(c.ColorCubeBitSize))
	return nearestNeighborStrategy{list.New(), color.RGBA{255, 255, 255, 255}, surface, fileName, cube}
}

type pixel struct {
	x, y int
}

type nearestNeighborStrategy struct {
	pixelBuffer *list.List
	startColor  color.RGBA
	surface     workSurface.Surface
	fileName    string
	cube        *colorCube.ColorCube
}

func (n nearestNeighborStrategy) GenerateImage(doneChan chan bool, imageUpdateChan chan ImageUpdate) {
	n.addPixelToDraw(workSurface.Coord2D{10, 10})
	//n.addPixelToDraw(workSurface.Coord2D{250, 250})
	//n.addPixelToDraw(workSurface.Coord2D{450, 450})

	n.surface.SetColor(10, 10, color.RGBA{uint8(255), uint8(255), uint8(255), 255})
	//n.surface.SetColor(250, 250, color.RGBA{uint8(0), uint8(255), uint8(0), 255})
	//n.surface.SetColor(450, 450, color.RGBA{uint8(0), uint8(0), uint8(255), 255})

	// totalNumberOfPixels := n.surface.Size * n.surface.Size
	currentPixel := 1

	for n.pixelBuffer.Len() > 0 {
		nextPixel := n.getNextPixel()
		usedPixels, unUsedPixels := n.surface.FindNeighborPixels(nextPixel)
		n.addPixelsToDraw(unUsedPixels)
		if currentPixel%1000 == 0 {
			// fmt.Printf("Current pixel (%4v , %4v) %6v / %6v , queueLenght %4v\n", nextPixel.X, nextPixel.Y, currentPixel, totalNumberOfPixels, n.pixelBuffer.Len())
		}
		currentPixel++

		//get average color for used neighbor pixels
		var avgColor color.RGBA
		if usedPixels.Len() > 0 {
			avgColor = n.getAverageColor(usedPixels)
		} else {
			avgColor = n.surface.GetColor(nextPixel.X, nextPixel.Y)
		}
		//find index for this color in colorcube
		x, y, z := n.cube.GetIndexForColor(avgColor)
		//if color at that index is not used
		if !n.cube.IsUsed(x, y, z) {
			n.cube.SetUsed(x, y, z)
			n.surface.SetColor(nextPixel.X, nextPixel.Y, n.cube.GetColor(x, y, z))
		} else {
			//  find nearest free color in cube
			found, foundX, foundY, foundZ := n.cube.FindUnusedColorInCube(x, y, z)
			if found {
				//	set as used, and color surface with it. continue loop
				n.cube.SetUsed(foundX, foundY, foundZ)
				n.surface.SetColor(nextPixel.X, nextPixel.Y, n.cube.GetColor(foundX, foundY, foundZ))
				imageUpdateChan <- ImageUpdate{nextPixel.X, nextPixel.Y, byte(foundX), byte(foundY), byte(foundZ)}
			} else {
				panic("Coudnt fint color!")
			}
		}
	}
	n.surface.ToPng(n.fileName)
	doneChan <- true
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
	var randVal int
	if n.pixelBuffer.Len() < 2 {
		randVal = rand.Intn(n.pixelBuffer.Len())
	} else {
		randVal = rand.Intn(2)
	}

	elem := n.pixelBuffer.Front()
	for i := 0; i < randVal; i++ {
		elem = elem.Next()
	}
	//	elem := n.pixelBuffer.Front()
	p, ok := elem.Value.(workSurface.Coord2D)
	if !ok {
		panic("Not a pixel in list!")
	}
	n.pixelBuffer.Remove(elem)
	return p
}
