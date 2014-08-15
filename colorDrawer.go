package main

import (
	// "fmt"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/nordsoyv/colorDrawer/config"
	"github.com/nordsoyv/colorDrawer/strategy"
)

const FPS int = 60

func main() {
	var (
		display    *allegro.Display
		eventQueue *allegro.EventQueue
		running    bool = true
		err        error
	)

	configuration := config.Read("config.json")
	strat := strategy.NearestNeighbor(configuration)

	if err := allegro.Install(); err != nil {
		panic(err)
	}
	defer allegro.Uninstall()

	// Create a 640x480 window and give it a title.
	allegro.SetNewDisplayFlags(allegro.WINDOWED)
	if display, err = allegro.CreateDisplay(512, 512); err == nil {
		defer display.Destroy()
		display.SetWindowTitle("colorDrawer")
	} else {
		panic(err)
	}

	// Create an event queue. All of the event sources we care about should
	// register themselves to this queue.
	if eventQueue, err = allegro.CreateEventQueue(); err == nil {
		defer eventQueue.Destroy()
	} else {
		panic(err)
	}

	// Calculate the timeout value based on the desired FPS.
	// timeout := float64(1) / float64(FPS)
	// timeout := float64(1)

	// Register event sources.
	eventQueue.Register(display)

	// Set the screen to black.
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
	allegro.FlipDisplay()

	doneChan := make(chan bool)
	imageUpdateChan := make(chan strategy.ImageUpdate, 100)
	go strat.GenerateImage(doneChan, imageUpdateChan)

	running = true
	// Main loop.
	var event allegro.Event
	for running {
		e, _ := eventQueue.GetNextEvent(&event)
		if e != nil {
			switch e.(type) {
			case allegro.DisplayCloseEvent:
				running = false
				break

				// Handle other events here.
			}
		}

		if numItems := len(imageUpdateChan); numItems > 0 {
			buffer := display.Backbuffer()
			allegro.SetTargetBitmap(buffer)
			for i := 0; i < numItems; i++ {
				imgUp := <-imageUpdateChan
				allegro.PutPixel(imgUp.X, imgUp.Y, allegro.MapRGB(imgUp.R, imgUp.G, imgUp.B))

			}
			allegro.FlipDisplay()
		}

		select {
		case <-doneChan:
			running = false
		default:
			continue
		}

	}

}
