package strategy

import (
	"container/list"
	"fmt"
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
	n.surface.SetColor(10, 10, color.RGBA{uint8(255), uint8(255), uint8(255), 255})

	pixelInputQueue := make(chan workSurface.Coord2D, 10)
	pixelOutputQueue := make(chan workSurface.Coord2D, 1)

	go pixelQueue(pixelInputQueue, pixelOutputQueue)
	pixelInputQueue <- workSurface.Coord2D{11, 11}
	pixelInputQueue <- workSurface.Coord2D{11, 10}
	pixelInputQueue <- workSurface.Coord2D{11, 9}

	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("Start loop")

			for !n.surface.IsFilled() {
				// fmt.Println("Getting pixel")
				nextPixel := <-pixelOutputQueue
				// fmt.Println("GOT pixel")
				usedPixels, unUsedPixels := n.surface.FindNeighborPixels(nextPixel)
				addPixelsToDraw(pixelInputQueue, unUsedPixels)

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
		}()

	}

	// n.surface.ToPng(n.fileName)
	// doneChan <- true
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

func pixelQueue(in chan workSurface.Coord2D, out chan workSurface.Coord2D) {
	queue := list.New()

	var newPixel workSurface.Coord2D
	var nextPixel workSurface.Coord2D
	counter := 0
	for {
		counter++
		for len(in) > 0 {
			// fmt.Printf("Queue is 0\n")
			newPixel = <-in
			addPixelToQueue(newPixel, queue)
		}
		nextPixel = getNextPixel(queue)
		for queue.Len() > 0 {
			// fmt.Printf("Looping : %v , length : %v\n", counter, queue.Len())

			select {
			case newPixel = <-in:
				// fmt.Println("Adding pixel select")
				addPixelToQueue(newPixel, queue)
			case out <- nextPixel:
				// fmt.Println("Removing pixel select")
				nextPixel = getNextPixel(queue)
			}
		}

	}

}

func addPixelToQueue(pixelToAdd workSurface.Coord2D, queue *list.List) {
	// if n.surface.IsUsed(pixelToAdd.X, pixelToAdd.Y) {
	// 	return false
	// }
	// fmt.Println("Adding pixel")
	for e := queue.Front(); e != nil; e = e.Next() {
		p := e.Value.(workSurface.Coord2D)
		if p.X == pixelToAdd.X && p.Y == pixelToAdd.Y {
			//all ready in queue
			return
		}
	}
	queue.PushBack(pixelToAdd)
}

func addPixelsToDraw(in chan workSurface.Coord2D, l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		p := e.Value.(workSurface.Coord2D)
		in <- p
	}
}

func getNextPixel(queue *list.List) workSurface.Coord2D {
	// fmt.Println("Removing pixel inner")
	var randVal int
	if queue.Len() == 0 {
		panic("Queue is 0")
	}
	if queue.Len() < 5 {
		randVal = rand.Intn(queue.Len())
	} else {
		randVal = rand.Intn(5)
	}

	elem := queue.Back()
	for i := 0; i < randVal; i++ {
		elem = elem.Prev()
	}
	p, ok := elem.Value.(workSurface.Coord2D)
	if !ok {
		panic("Not a pixel in list!")
	}
	queue.Remove(elem)
	return p
}
